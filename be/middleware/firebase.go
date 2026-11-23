package middleware

import (
	"context"
	"net/http"
	"strings"
	"zydukas-rankings/models"
	"zydukas-rankings/utils"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func InitFirebase() (*auth.Client, error) {
	firebaseOptPath, frErr := utils.EnvString("FIREBASE_CONFIG_PATH")
	if frErr != nil {
		utils.Fatal("Error reading FIREBASE_CONFIG_PATH")
		return nil, frErr
	}

	opt := option.WithAuthCredentialsFile(option.ServiceAccount, firebaseOptPath)
	app, fErr := firebase.NewApp(context.Background(), nil, opt)

	if fErr != nil {
		utils.Fatal("error initializing app: %v")
		return nil, fErr
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		utils.Fatal("Failed to initialize firebase auth client.")
		return nil, err
	}

	return authClient, nil
}

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
		tokenString = strings.TrimSpace(tokenString)

		//zap.L().Info("Firebase token:" + tokenString)

		token, err := authClient.VerifyIDToken(g.Request.Context(), tokenString)
		if err != nil {
			msg := err.Error()
			utils.Error(msg)
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrInvalidJwt)
			return
		}

		//utils.Info(header)

		firebaseId, ok := token.Claims["user_id"].(string)

		//utils.Info(firebaseId)

		if !ok {
			g.AbortWithStatusJSON(http.StatusUnauthorized, &utils.ErrInvalidJwt)
			return
		}

		authUser := models.UserContext{
			FireBaseId: firebaseId,
		}

		g.Set(utils.UserContextTitle, authUser)

		utils.Info("Not firebase right?")

		g.Next()
	}
}
