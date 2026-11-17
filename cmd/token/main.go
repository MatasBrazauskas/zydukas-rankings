// Command token mints a real Firebase ID token for local development.
//
// It uses the admin service account to create a custom token, then exchanges
// that for an ID token through the Identity Toolkit REST API -- the same kind
// of token a browser client would send, so middleware.FirebaseMiddleware
// accepts it.
//
//	go run ./cmd/token                      # print a token for the default dev user
//	go run ./cmd/token -uid bob -email b@x.io
//	curl -H "Authorization: Bearer $(go run ./cmd/token -q)" localhost:8080/users/create
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"zydukas-rankings/utils"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const signInURL = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key="

func main() {
	uid := flag.String("uid", "dev-user", "Firebase UID to mint a token for")
	email := flag.String("email", "dev@zydukas.local", "email claim on the user record")
	name := flag.String("name", "Dev User", "display name -> \"name\" claim")
	quiet := flag.Bool("q", false, "print only the raw token (for use in $(...))")
	flag.Parse()

	logf := func(format string, args ...any) {
		if !*quiet {
			fmt.Fprintf(os.Stderr, format+"\n", args...)
		}
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("loading .env: %v", err)
	}

	apiKey := utils.EnvString("FIREBASE_WEB_API_KEY", "")
	if apiKey == "" {
		log.Fatal("FIREBASE_WEB_API_KEY is not set in .env")
	}

	ctx := context.Background()
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, utils.EnvString("FIREBASE_CONFIG_PATH", "."))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("initializing firebase: %v", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("initializing auth client: %v", err)
	}

	if err := ensureUser(ctx, client, *uid, *email, *name); err != nil {
		log.Fatalf("ensuring user %q: %v", *uid, err)
	}
	logf("user:  %s <%s>", *uid, *email)

	customToken, err := client.CustomToken(ctx, *uid)
	if err != nil {
		log.Fatalf("minting custom token: %v", err)
	}

	idToken, err := exchange(ctx, apiKey, customToken)
	if err != nil {
		log.Fatalf("exchanging custom token: %v", err)
	}

	// Self-check through the exact call the middleware makes.
	verified, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("token did not verify: %v", err)
	}
	logf("valid: iss=%s aud=%s", verified.Issuer, verified.Audience)
	logf("until: %s (1h)", time.Unix(verified.Expires, 0).Format(time.RFC3339))

	fmt.Println(idToken)
}

// ensureUser creates the account if missing, or refreshes its email and name.
func ensureUser(ctx context.Context, client *auth.Client, uid, email, name string) error {
	_, err := client.GetUser(ctx, uid)
	if err == nil {
		update := (&auth.UserToUpdate{}).Email(email).DisplayName(name).EmailVerified(true)
		_, err = client.UpdateUser(ctx, uid, update)
		return err
	}
	if !auth.IsUserNotFound(err) {
		return err
	}
	create := (&auth.UserToCreate{}).UID(uid).Email(email).DisplayName(name).EmailVerified(true)
	_, err = client.CreateUser(ctx, create)
	return err
}

// exchange trades a custom token for a real Firebase ID token.
func exchange(ctx context.Context, apiKey, customToken string) (string, error) {
	body, err := json.Marshal(map[string]any{
		"token":             customToken,
		"returnSecureToken": true,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, signInURL+apiKey, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("identitytoolkit returned %s: %s", res.Status, payload)
	}

	var parsed struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(payload, &parsed); err != nil {
		return "", err
	}
	if parsed.IDToken == "" {
		return "", fmt.Errorf("no idToken in response: %s", payload)
	}
	return parsed.IDToken, nil
}
