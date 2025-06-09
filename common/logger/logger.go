package logger

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"starter-go-gin/common/constant"
	"starter-go-gin/config"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
	Severity   string
	UserAgent  string
}

func ZeroGinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		// full absolute url
		scheme := "http://"
		if c.Request.TLS != nil {
			scheme = "https://"
		}
		path := scheme + c.Request.Host + c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// put trace id to context
		ctx, _ := trace.StartSpan(c.Request.Context(), "gin.request")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = c.Request.Method + " - " + path
		}
		cData := &ginHands{
			SerName:    "GIN",
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			MsgStr:     msg,
			UserAgent:  c.Request.UserAgent(),
		}

		logGinSwitch(c.Request.Context(), cData)
	}
}

func Initialize(cfg *config.Config) {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())

	zerolog.LevelFieldName = "severity"
	zerolog.CallerSkipFrameCount = 8
}

func logGinSwitch(ctx context.Context, data *ginHands) {
	switch {
	case data.StatusCode >= constant.FourHundred && data.StatusCode < constant.FiveHundred:
		{
			getHTTPRequestEventMessage(ctx, data, zerolog.WarnLevel).
				Msg(data.MsgStr)
		}
	case data.StatusCode >= constant.FiveHundred:
		{
			getHTTPRequestEventMessage(ctx, data, zerolog.ErrorLevel).
				Msg(data.MsgStr)
		}
	default:
		getHTTPRequestEventMessage(ctx, data, zerolog.InfoLevel).
			Msg(data.MsgStr)
	}
}

// GetTraceID returns the trace ID from the context.
func GetTraceID(ctx context.Context) string {
	return trace.FromContext(ctx).SpanContext().TraceID.String()
}

// GetSpanID returns the span ID from the context.
func GetSpanID(ctx context.Context) string {
	return trace.FromContext(ctx).SpanContext().SpanID.String()
}

// Error logs a message at level Error on the standard logger.
func Error(ctx context.Context, err error) {
	getEventMessage(ctx, zerolog.ErrorLevel).
		Err(err)
	sentry.CaptureException(err)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(ctx context.Context, msg string, data ...interface{}) {
	data = append([]interface{}{msg}, data...)
	getEventMessage(ctx, zerolog.InfoLevel).
		Msgf("%v", data...)
}

// WarnFromStr logs a message at level Warn on the standard logger.
func WarnFromStr(ctx context.Context, message string) {
	getEventMessage(ctx, zerolog.WarnLevel).
		Msg(message)
}

// ErrorWithStr logs a message at level Error on the standard logger.
func ErrorWithStr(ctx context.Context, msg string, err error) {
	getEventMessage(ctx, zerolog.ErrorLevel).
		Str("msg", msg).
		Err(err)
	sentry.CaptureException(err)
}

// ErrorFromStr logs a message at level Error on the standard logger.
func ErrorFromStr(ctx context.Context, errStr ...interface{}) {
	getEventMessage(ctx, zerolog.ErrorLevel).
		Msgf("%v", errStr...)
	sentry.CaptureException(errors.New(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(errStr...)), " "), "[]")))
}

// Info logs a message at level Info on the standard logger.
func Info(ctx context.Context, msg string, data ...interface{}) {
	data = append([]interface{}{msg}, data...)
	getEventMessage(ctx, zerolog.InfoLevel).
		Msgf("%v", data...)
}

func getHTTPRequestEventMessage(ctx context.Context, data *ginHands, level zerolog.Level) *zerolog.Event {
	// get userID string form context
	userID := ""
	usr := ctx.Value(constant.UserIDKey)
	if usr != nil {
		userID = usr.(uuid.UUID).String()
	}

	return log.WithLevel(level).
		Timestamp().
		Caller().
		Str("source", data.SerName).
		Dict("httpRequest", zerolog.Dict().
			Str("requestMethod", data.Method).
			Str("requestUrl", data.Path).
			Str("status", fmt.Sprintf("%d", data.StatusCode)).
			Str("userAgent", data.UserAgent).
			Str("remoteIp", data.ClientIP).
			Dur("latency", data.Latency)).
		Str("logging.googleapis.com/trace", GetTraceID(ctx)).
		Str("logging.googleapis.com/spanId", GetSpanID(ctx)).
		Str("userId", userID)
}

func getEventMessage(ctx context.Context, level zerolog.Level) *zerolog.Event {
	// get userID string form context
	userID := ""
	usr := ctx.Value(constant.UserIDKey)
	if usr != nil {
		userID = usr.(uuid.UUID).String()
	}

	return log.WithLevel(level).
		Timestamp().
		Caller().
		Str("logging.googleapis.com/trace", GetTraceID(ctx)).
		Str("logging.googleapis.com/spanId", GetSpanID(ctx)).
		Str("userId", userID)
}
