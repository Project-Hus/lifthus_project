package session

import (
	"context"
	"fmt"
	"hus-auth/ent"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CreateHusSession takes user's uuid and create new hus session and return it.
func CreateNewHusSession(ctx context.Context, client *ent.Client, uid uuid.UUID, exp bool) (
	new_sid, new_token string, err error,
) {
	// create new Hus session in database
	nhs := client.HusSession.Create().SetUID(uid)
	if exp { // if it's set to expired, give it 7 days expiration
		nhs = nhs.SetExp(time.Now().Add(time.Hour * 24 * 7))
	}
	hs, err := nhs.Save(ctx)
	if err != nil {
		return "", "", fmt.Errorf("!!creating new hus session failed:%w", err)
	}

	var rt *jwt.Token

	// using created session's UUID, create session token
	if exp {
		rt = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sid":     hs.ID,                     // refresh token's uuid
			"purpose": "hus_session",             // purpose"
			"iss":     os.Getenv("HUS_AUTH_URL"), // issuer
			"uid":     uid,                       // user's uuid
			"iat":     hs.Iat,                    // issued at
			"exp":     hs.Exp.Unix(),             // expiration : one week
		})
	} else {
		rt = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sid":     hs.ID,                                // refresh token's uuid
			"purpose": "hus_session",                        // purpose"
			"iss":     os.Getenv("HUS_AUTH_URL"),            // issuer
			"uid":     uid,                                  // user's uuid
			"iat":     hs.Iat,                               // issued at
			"exp":     time.Now().Add(time.Hour * 1).Unix(), // expiration : an hour
		})
	}

	hsk := []byte(os.Getenv("HUS_AUTH_TOKEN_KEY"))

	rts, err := rt.SignedString(hsk)
	if err != nil {
		return "", "", fmt.Errorf("!!signing hus-session token failed:%w", err)
	}
	log.Printf("hus-session created by (%s)", uid)
	// Sign and return the complete encoded token as a string
	return hs.ID.String(), rts, nil
}

// ValidateHusSessionToken takes hus-session token and validate it.
// and it also revokes validated token and return new token.
func ValidateHusSessionToken(ctx context.Context, client *ent.Client, st string) (new_token string, err error) {
	return "", nil
}
