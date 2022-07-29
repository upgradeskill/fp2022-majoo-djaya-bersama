package security

import (
	"fmt"
	"mini-pos/dto"
	"mini-pos/util"
	"time"

	"github.com/golang-jwt/jwt"
)

//jwt service
type JWTService interface {
	GenerateToken(id uint, name string, role int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    util.GlobalConfig.JWT_ISSUER,
	}
}

func getSecretKey() (secret string) {
	secret = util.GlobalConfig.JWT_SECRET
	if secret == "" {
		secret = "BO0kSt0r3"
	}
	return
}

func getIssuer() (issuer string) {
	issuer = util.GlobalConfig.JWT_ISSUER
	if issuer == "" {
		issuer = "bookstore"
	}
	return
}

func (service *jwtServices) GenerateToken(id uint, name string, role int) (string, error) {
	claims := &dto.UserClaims{
		id,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	return t, err
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
