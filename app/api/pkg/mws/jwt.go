package mws

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/crazyfrankie/seekmall/app/api/config"
)

type Claims struct {
	UId int
	jwt.MapClaims
}

type AuthBuilder struct {
	paths map[string]struct{}
}

func NewAuthBuilder() *AuthBuilder {
	return &AuthBuilder{
		paths: make(map[string]struct{}),
	}
}

func (a *AuthBuilder) IgnorePath(path string) *AuthBuilder {
	a.paths[path] = struct{}{}
	return a
}

func (a *AuthBuilder) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := a.paths[c.Request.URL.Path]; ok {
			c.Next()
			return
		}

		tokenHeader := c.GetHeader("Authorization")
		token := extractToken(tokenHeader)

		claims, err := parseToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func extractToken(str string) string {
	if str == "" {
		return ""
	}
	token := strings.Split(str, " ")
	if token[0] != "Bearer" {
		return ""
	}

	return token[1]
}

func parseToken(str string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(str, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConf().JWT.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && claims != nil {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}
