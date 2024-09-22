package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/iki-rumondor/go-tbc/internal/app/structs/response"
	"github.com/iki-rumondor/go-tbc/internal/utils"
)

func IsValidJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			utils.HandleError(c, response.UNAUTH_ERR("Bearer Token Tidak Ditemukan"))
			return
		}

		jwt := strings.Split(headerToken, " ")[1]
		if jwt == "null" {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		c.Set("map_claims", mapClaims)
		c.Next()

	}
}

func SetUserUuid() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		uuid, ok := mapClaims["uuid"].(string)
		if !ok {
			utils.HandleError(c, response.UNAUTH_ERR("Token Tidak Valid"))
			return
		}

		c.Set("uuid", uuid)
		c.Next()
	}
}

func IsRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		roleJwt := mapClaims["role"].(string)
		if roleJwt != role {
			utils.HandleError(c, response.UNAUTH_ERR("Hak Akses Dibatasi"))
			return
		}
		c.Next()
	}
}
