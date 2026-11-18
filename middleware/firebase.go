package middleware

import (
	"log"
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

		log.Print(header)

		firebaseId, ok := token.Claims["firebase_id"].(string)

		if !ok {
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrInvalidJwt)
			return
		}

		authUser := models.UserContext{
			FireBaseId: firebaseId,
		}

		g.Set(utils.UserContextTitle, authUser)
		g.Next()
	}
}
