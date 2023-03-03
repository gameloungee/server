package account

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gameloungee/server/config"
	"github.com/gameloungee/server/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secrets = func() []byte {
	conf := config.New()

	if conf.JwtSecrets != config.CHANGE_ME_STATE {
		return []byte(conf.JwtSecrets)
	}

	return []byte("0000")
}

type JWT struct {
	Account
	jwt.StandardClaims
}

func GenerateJWT(a Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWT{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(720 * time.Hour).Unix(),
		},
		Account: a,
	})

	return token.SignedString(secrets)
}

func ParseJWT(rawToken string) (Account, error) {
	token, err := jwt.ParseWithClaims(rawToken, &JWT{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexcepted signing method: %s", t.Header["Alg"])
		}

		return secrets, nil
	})

	if err != nil {
		return Account{}, err
	}

	if claims, ok := token.Claims.(*JWT); ok && token.Valid {
		return claims.Account, nil
	}

	return Account{}, errors.New("invalid JWT token")
}

func ParseJWTFromHeader(c *gin.Context) Account {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return Account{}
	}

	headerParts := strings.Split(authHeader, " ")

	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return Account{}
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return Account{}
	}

	a, err := ParseJWT(headerParts[1])

	if err != nil {
		response.AbortWith(http.StatusUnauthorized, "You are not authorized", "The header of your request does not contain the authorisation key", c)
		return Account{}
	}

	return a
}
