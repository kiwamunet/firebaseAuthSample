package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/k0kubun/pp"
	"google.golang.org/api/option"
)

type Credential struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func main() {
	token := `eyJhbGciXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`

	ctx := context.Background()

	jwt := Lookup()
	jwt.PrivateKey = strings.Replace(jwt.PrivateKey, "\\n", "\n", -1)

	jwtByte, err := json.Marshal(*jwt)
	if err != nil {
		log.Println(err)
	}
	opt := option.WithCredentialsJSON(jwtByte)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println(err)
	}

	client, err := app.Auth(ctx)
	idToken, err := client.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		log.Println(err)
	}
	pp.Println(idToken)

}

func Lookup() *Credential {
	return &Credential{
		"service_account",
		"XXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXXX",
		`-----BEGIN PRIVATE KEY-----\nXXXXXXXXXXXXXXXXX\n-----END PRIVATE KEY-----\n`,
		"XXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXXX",
		"https://accounts.google.com/o/oauth2/auth",
		"https://oauth2.googleapis.com/token",
		"https://www.googleapis.com/oauth2/v1/certs",
		"XXXXXXXXXXXXXXXXX",
	}
}
