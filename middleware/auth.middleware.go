package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/response"
	"starter-go-gin/utils"
)

// Auth is a middleware for authentication
func Auth(cfg config.Config, cache interfaces.Cacheable) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < constant.Two {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponseWithoutReqID(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.SuccessAPIResponseWithoutReqID(http.StatusUnauthorized, err.Error(), nil))
			c.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			if err := cache.BulkRemove(fmt.Sprintf(commonCache.TokenUserByJTI, claims.Subject.ID)); err != nil {
				logger.Error(c.Request.Context(), err)
			}

			c.JSON(http.StatusUnauthorized, response.SuccessAPIResponseWithoutReqID(http.StatusUnauthorized, "token expired", nil))
			c.Abort()
			return
		}

		// add header from claims
		c.Request.Header.Set(constant.RequestHeaderName, claims.Subject.Name)
		c.Request.Header.Set(constant.RequestHeaderUserID, claims.Subject.ID.String())
		c.Request.Header.Set(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		c.Request.Header.Set(constant.RequestHeaderJTI, claims.ID)
		c.Header(constant.RequestHeaderName, claims.Subject.Name)
		c.Header(constant.RequestHeaderUserID, claims.Subject.ID.String())
		c.Header(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		c.Header(constant.RequestHeaderJTI, claims.ID)

		// put userID into context from ginContext
		reqNew := c.Request.WithContext(context.WithValue(c.Request.Context(), constant.UserIDKey, claims.Subject.ID))
		reqNew.Header.Set(constant.RequestHeaderName, claims.Subject.Name)
		reqNew.Header.Set(constant.RequestHeaderUserID, claims.Subject.ID.String())
		reqNew.Header.Set(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		reqNew.Header.Set(constant.RequestHeaderJTI, claims.ID)

		c.Request = reqNew
		if claims.Issuer != cfg.JWTConfig.Issuer {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponseWithoutReqID(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// Admin is a middleware for admin authentication
func Admin(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < constant.Two {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponseWithoutReqID(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponseWithoutReqID(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		// add header from claims
		c.Request.Header.Set(constant.RequestHeaderName, claims.Subject.Name)
		c.Request.Header.Set(constant.RequestHeaderUserID, claims.Subject.ID.String())
		c.Request.Header.Set(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		c.Request.Header.Set(constant.RequestHeaderJTI, claims.ID)
		c.Header(constant.RequestHeaderName, claims.Subject.Name)
		c.Header(constant.RequestHeaderUserID, claims.Subject.ID.String())
		c.Header(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		c.Header(constant.RequestHeaderJTI, claims.ID)

		// put userID into context from ginContext
		reqNew := c.Request.WithContext(context.WithValue(c.Request.Context(), constant.UserIDKey, claims.Subject.ID))
		reqNew.Header.Set(constant.RequestHeaderName, claims.Subject.Name)
		reqNew.Header.Set(constant.RequestHeaderUserID, claims.Subject.ID.String())
		reqNew.Header.Set(constant.RequestHeaderRoleID, claims.Subject.RoleID.String())
		reqNew.Header.Set(constant.RequestHeaderJTI, claims.ID)

		c.Request = reqNew

		if claims.Issuer != cfg.JWTConfig.IssuerAdmin {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponseWithoutReqID(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}
