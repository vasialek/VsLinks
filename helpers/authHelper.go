package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vasialek/VsLinks/models"
)

// AuthHelper used to access it as class
type AuthHelper struct{}

// NewAuthHelper returns instance of helper
func NewAuthHelper() *AuthHelper {
	return &AuthHelper{}
}

// Decode decodes JWT w/o validation
func (ah *AuthHelper) Decode(jwtToken string) (map[string]string, error) {
	payload := make(map[string]string)
	var list map[string]interface{}

	ar := strings.Split(jwtToken, ".")
	if len(ar) != 3 {
		return payload, errors.New("JWT must has 3 segments, separated by '.'")
	}

	// Decode header
	data, err := jwt.DecodeSegment(ar[0])
	if err != nil {
		return payload, err
	}

	if err = json.Unmarshal(data, &list); err != nil {
		return payload, err
	}

	payload["alg"] = fmt.Sprintf("%s", list["alg"])

	// Decode payload
	data, err = jwt.DecodeSegment(ar[1])
	if err != nil {
		return payload, err
	}
	fmt.Println("Decoded JWT[1]:", string(data))
	if err = json.Unmarshal(data, &list); err != nil {
		return payload, err
	}

	// Make simpla map of values
	for n, v := range list {
		switch v.(type) {
		case float64:
			payload[n] = fmt.Sprintf("%.0f", v)
		default:
			payload[n] = fmt.Sprintf("%v", v)
		}
	}

	return payload, nil
}

// ValidateHeader returns error in case HTTP header does not contain valid JWT
func (ah *AuthHelper) ValidateHeader(header string) (*models.User, error) {
	tempSecret := []byte("Test1234")
	if len(header) < 1 {
		return nil, errors.New("authentication header is empty")
	}
	if strings.HasPrefix(header, "Bearer ") == false {
		return nil, errors.New("authentication header should start with Bearer")
	}

	fmt.Printf("JWT w/o Bearer: `%s`\n", header[7:])
	claims := &models.UserClaims{}

	token, err := jwt.ParseWithClaims(header[7:], claims, func(t *jwt.Token) (interface{}, error) {
		return tempSecret, nil
	})

	if err == nil {
		fmt.Println("Decoded JWT:", token)
		fmt.Println("  claims:", claims)
		return &models.User{
			UserID: claims.UserID,
			Name:   claims.Name,
			Email:  claims.Email,
		}, nil
	}

	return nil, err
}

// GenerateJwt returns encoded JWT string
func (ah *AuthHelper) GenerateJwt(user *models.User, expirationInMin int) (string, error) {
	expiration := time.Now().Add(time.Minute * time.Duration(expirationInMin))
	claims := models.UserClaims{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("Test1234"))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
