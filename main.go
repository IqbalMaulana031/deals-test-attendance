package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"starter-go-gin/app"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/resource"
	"starter-go-gin/utils"
	"time"

	authBuilder "starter-go-gin/modules/auth/v1/builder"
	authHandler "starter-go-gin/modules/auth/v1/handler"
	productBuilder "starter-go-gin/modules/product/v1/builder"
	userBuilder "starter-go-gin/modules/user/v1/builder"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/fx"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Define a simple service
type HelloService struct{}

func (h *HelloService) SayHello() {
	fmt.Println("Hello from Fx!")
}

const (
	statusOK = "OK"
)

var environment string
var appName string
var appVersion string

// Health is a base struct for health check
type Health struct {
	Status   string `json:"status"`
	Database string `json:"database"`
	Redis    string `json:"redis"`
}

var health *Health

// splash print plain text message to console
func splash() {
	fmt.Print(`
        .__                   __                 __
   ____ |__| ____     _______/  |______ ________/  |_  ___________
  / ___\|  |/    \   /  ___/\   __\__  \\_  __ \   __\/ __ \_  __ \
 / /_/  >  |   |  \  \___ \  |  |  / __ \|  | \/|  | \  ___/|  | \/
 \___  /|__|___|  / /____  > |__| (____  /__|   |__|  \___  >__|
/_____/         \/       \/            \/                 \/
`)
}

// Define an Fx provider for HelloService
func NewHelloService() *HelloService {
	return &HelloService{}
}

func NewConfig() *config.Config {
	// Load your configuration here
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	return cfg
}

func NewConfig2() config.Config {
	// Load your configuration here
	cfg, err := config.LoadConfig(".env")
	checkError(err)

	return *cfg
}

func NewLogger() *log.Logger {
	return &log.Logger{}
}

func NewGormLogger(cfg *config.Config) *logger.GormLogger {
	if cfg.Env == "production" {
		return logger.NewGormLogger(gormlogger.Error)
	}
	return logger.NewGormLogger(gormlogger.Info)
}

func NewRedisPool(cfg *config.Config) *redis.Pool {
	return buildRedisPool(cfg)
}

func NewRedisCacheClient(redisPool *redis.Pool) *utils.Client {
	return utils.NewRedisClient(redisPool)
}

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Recovery())
	// router.Use(logger.ZeroGinLogger())
	return router
}

var Module = fx.Options(
	fx.Provide(
		NewConfig,
		NewConfig2,
		NewLogger,
		NewGormLogger,
		utils.NewPostgresGormDB,
		NewRedisPool,
		fx.Annotate(
			NewRedisCacheClient,
			fx.As(new(interfaces.Cacheable)),
		),
		NewRouter,
	),
)

func main() {
	err := os.Setenv("TZ", "Asia/Jakarta")
	if err != nil {
		return
	}
	health = &Health{}

	splash()
	// Create and start an Fx application
	fx.New(
		Module,
		authBuilder.AuthModule,
		userBuilder.UserModule,
		productBuilder.ProductModule,
		fx.Invoke(app.DefaultHTTPHandler),
		fx.Invoke(startServer),
	).Run()
}

func startServer(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	db *gorm.DB,
	redisPool *redis.Pool,
	router *gin.Engine,
	authHandler *authHandler.AuthHandler,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Initialize(cfg)

				environment = cfg.Env
				appName = cfg.AppName
				appVersion = cfg.AppVersion
				health.Database = statusOK
				health.Redis = statusOK

				if cfg.Env == "production" {
					gin.SetMode(gin.ReleaseMode)
				}

				router.GET("/", Home)
				router.GET("/health-check", HealthGET)

				fmt.Println("Default timezone: ", time.Now().Location().String())
				fmt.Println("App version: ", appVersion)

				health.Status = statusOK
				go func() {
					if err := router.Run(fmt.Sprintf(":%s", cfg.Port.APP)); err != nil {
						panic(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				// Handle cleanup here
				return nil
			},
		},
	)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// CORSMiddleware is a function to add CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

// buildRedisPool is a function to build redis pool
func buildRedisPool(cfg *config.Config) *redis.Pool {
	cachePool := utils.NewPool(cfg.Redis.Address, cfg.Redis.Password)

	ctx := context.Background()
	_, err := cachePool.GetContext(ctx)

	if err != nil {
		checkError(err)
	}

	log.Println("redis successfully connected!")
	return cachePool
}

// HealthGET is a function to handle health check
func HealthGET(c *gin.Context) {
	c.JSON(http.StatusOK, health)
}

// Home is a function to handle home page
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"app_name":    appName,
		"environment": environment,
		"version":     appVersion,
		"status":      "running",
		"date":        time.Now().Format(time.RFC3339),
	})
}

// AssetLinks is a function to handle assetlinks.json
func AssetLinks(c *gin.Context) {
	var req resource.File

	err := c.ShouldBindUri(&req)
	if err != nil {
		return
	}
	c.File("./" + req.Path)
}
