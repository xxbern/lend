package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/ogen-go/ogen/conv"
	"lend/gen/model"
	"log"
	"strconv"
	"time"
)

type TokenUser struct {
	*model.UserInfo
}

type MyTokenClaim struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles,omitempty"`
}

const jwtKey = "AllYourBase"

func (user *TokenUser) NewToken() (tokenStr string, err error) {
	mySigningKey := []byte(jwtKey)

	// Create the Claims
	claims := &MyTokenClaim{
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "lend",
			Subject:   *user.Name,
			ID:        strconv.Itoa(int(user.ID)),
		},
		[]string{
			"admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func Decode2User(tokenStr string) (user *model.UserInfo, roles []string, err error) {
	mySigningKey := []byte(jwtKey)
	token, err := jwt.ParseWithClaims(tokenStr, &MyTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if token == nil || err != nil {
		log.Println(token, err)
		return nil, nil, err
	}

	tokenClaim := (token.Claims).(*MyTokenClaim)
	id, _ := conv.ToInt32(tokenClaim.ID)

	return &model.UserInfo{
		ID:   id,
		Name: &tokenClaim.Subject,
	}, tokenClaim.Roles, nil
}
