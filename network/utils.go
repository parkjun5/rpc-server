package network

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. BarrerToken  을 가져온다.
		t := getAuthToken(c)
		if t == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			if _, err := n.gRPCClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")

	// Bearar ~~

	authSide := strings.Split(authToken, " ")
	if len(authSide) > 1 {
		token = authSide[1]
	}

	return token
}
