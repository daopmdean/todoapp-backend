package access_token

import (
	"errors"
	"fmt"
	"time"
	"todoapp/internal/common/config"
	"todoapp/internal/entity"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user *entity.User, timeInit time.Time) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["username"] = user.Username
	atClaims["exp"] = timeInit.Add(time.Hour * 24 * 7).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(config.GetAccessSecret()))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetAccessSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenClaims(token *jwt.Token) (*TokenData, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("cannot get token claims")
	}

	if !token.Valid {
		return nil, errors.New("token not valid")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("cannot get username")
	}

	return &TokenData{
		Username: username,
	}, nil
}

type TokenData struct {
	Username string
}
