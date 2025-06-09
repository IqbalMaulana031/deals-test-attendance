package utils

import (
	"crypto/md5" // #nosec
	"encoding/hex"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"starter-go-gin/config"

	//nolint:staticcheck
	"time"
)

const (
	daysInYear = 365
	dayInHour  = 24
)

// Subject is a struct that contains user id and role id
type Subject struct {
	ID         uuid.UUID `json:"id"`
	RoleID     uuid.UUID `json:"role_id"`
	Permission []string  `json:"permission"`
	MerchantID uuid.UUID `json:"merchant_id"`
	Name       string    `json:"name"`
}

// TokenClaims define available data in JWT Token
type TokenClaims struct {
	ExpiresAt int64   `json:"exp,omitempty"`
	ID        string  `json:"jti,omitempty"`
	IssuedAt  int64   `json:"iat,omitempty"`
	NotBefore int64   `json:"nbf,omitempty"`
	Subject   Subject `json:"sub,omitempty"`
	Issuer    string  `json:"iss,omitempty"`
	jwt.RegisteredClaims
}

// JWTDecode decodes JWT to token claims
func JWTDecode(cfg config.Config, t string) (*TokenClaims, error) {
	if cfg.JWTConfig.Public == "" {
		return nil, fmt.Errorf("please specify your public key path")
	}

	// publicKey, errKey := ioutil.ReadFile(cfg.JWTConfig.Public)
	// if errKey != nil {
	//	return nil, fmt.Errorf("error while reading public key file : %v", errKey)
	// }

	sourceKey := []byte(cfg.JWTConfig.Public)

	key, errParse := jwt.ParseRSAPublicKeyFromPEM(sourceKey)
	if errParse != nil {
		return nil, fmt.Errorf("failed to parse public key : %v", errParse)
	}

	token, err := jwt.ParseWithClaims(t, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("RS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token : %v", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// JWTEncode encode token claims to JWT
func JWTEncode(cfg config.Config, body map[string]interface{}, iss string, exp int64) (string, string, error) {
	if cfg.JWTConfig.Private == "" {
		return "", "", fmt.Errorf("please specify your public key path")
	}
	hashQuery := md5.New() // #nosec
	hashQuery.Write([]byte(fmt.Sprintf("secret123:%v", time.Now().Add(time.Hour*dayInHour*daysInYear).Unix())))

	jti := hex.EncodeToString(hashQuery.Sum(nil))

	log.Println("sub", body)

	// expire time in 24 hours
	sign := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": exp,
		//  "exp": time.Now().Add(time.Hour * constant.One).Unix(),
		"jti": jti,
		"sub": body,
		"iss": iss,
	})

	// privateKey, readErr := ioutil.ReadFile(cfg.JWTConfig.Private)
	// if readErr != nil {
	//	return "", "", fmt.Errorf("therewas and error while trying to read private key file : %v", readErr)
	// }

	sourceKey := []byte(cfg.JWTConfig.Private)

	key, parseErr := jwt.ParseRSAPrivateKeyFromPEM(sourceKey)
	if parseErr != nil {
		return "", "", fmt.Errorf("therewas and error while trying to parse private key : %v", parseErr)
	}

	token, err := sign.SignedString(key)

	if err != nil {
		return "", "", fmt.Errorf("therewas and error while trying to create token : %v", parseErr)
	}

	return token, jti, nil
}
