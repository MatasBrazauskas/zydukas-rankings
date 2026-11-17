package middleware

import (
	"net/http"
	"strings"
	"zydukas-rankings/models"
	"zydukas-rankings/utils"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func FirebaseMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(g *gin.Context) {

		header := g.GetHeader("Authorization")
		if header == "" {
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrNoAuthHeader)
			return
		}

		tokenString, ok := strings.CutPrefix(header, "Bearer ")
		if !ok {
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrNoAuthHeader)
			return
		}

		token, err := authClient.VerifyIDToken(g.Request.Context(), tokenString)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrInvalidJwt)
			return
		}

		email, ok := token.Claims["email"].(string)

		if !ok {
			email = ""
		}

		name, ok := token.Claims["name"].(string)

		if !ok {
			name = ""
		}

		authUser := models.UserContext{
			Username: name,
			Email:    email,
		}

		g.Set(utils.UserContextTitle, authUser)
		g.Next()
	}
}
