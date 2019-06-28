package helpers

import (
	"testing"

	"github.com/vasialek/VsLinks/models"
)

const Username string = "TestUser"
const Email string = "my@ema.il"

var authHelper *AuthHelper
var user *models.User

func init() {
	authHelper = NewAuthHelper()
	user = &models.User{
		Name:  Username,
		Email: Email,
	}
}

func TestValidateHeader_Error_WhenEmptyHeader(t *testing.T) {
	authenticationHelper := NewAuthHelper()

	_, err := authenticationHelper.ValidateHeader("")

	if err == nil {
		t.Error("Expecting error about empty Authentication header")
	}
}

func TestValidateHeader_Error_WhenNoBearer(t *testing.T) {
	_, err := NewAuthHelper().ValidateHeader("WithoutBearer")

	if err == nil {
		t.Error("Expecting error when Authorization header does not start with Bearer")
	}
}

func TestValidateHeader_Error_WhenExpired(t *testing.T) {
	_, err := NewAuthHelper().ValidateHeader("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiIiwibmFtZSI6IkFsZWtzZWogVi4iLCJlbWFpbCI6InAucm9nbGFtZXJAZ21haWwuY29tIiwiZXhwIjoxNTYxNjQ0NTgyfQ.IU0r1-1-lYSq6W3_d2wBriDzq-uzeiAzGJwBvwt4j04")

	if err == nil {
		t.Error("Expecting error about expired JWT")
	}
}

func TestValidateHeader_CheckUsername(t *testing.T) {
	token := createAuthHeader()

	u, _ := authHelper.ValidateHeader(token)

	if u == nil {
		t.Error("Got NULL user form ValidateHelper method, expected valid.")
	} else if u.Name != Username {
		t.Errorf("Expected decoded user name from JWT to be `%s`. Got `%s`", Username, u.Name)
	}
}

func TestValidateHeader_CheckEmail(t *testing.T) {
	header := createAuthHeader()

	u, err := authHelper.ValidateHeader(header)

	if err != nil {
		t.Error("Error validating JWT (Auth header). " + err.Error())
	} else if u.Email != Email {
		t.Errorf("Expecting user email to be `%s`, got `%s`\n", Email, u.Email)
	}
}

// #region Decode

func TestDecode_Error_WhenNotTwoDots(t *testing.T) {
	_, err := NewAuthHelper().Decode("eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlZzTGlua3MiLCJpYXQiOjE1MTYyMzkwMjJ9.")

	if err == nil {
		t.Error("Expecting error that JWT must has 2 dots", err)
	}
}

func TestDecode_ErrorIsNill(t *testing.T) {
	_, err := NewAuthHelper().Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlZzTGlua3MiLCJpYXQiOjE1MTYyMzkwMjJ9._FLAAeNLx6Gvvpv95dELIz52xDVAE500BWutDRjsUtQ")

	if err != nil {
		t.Errorf("Expecting NIL error, got %s\n", err)
	}
}

func TestDecode_CheckAlg(t *testing.T) {
	expected := "HS256"

	payload, _ := NewAuthHelper().Decode("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlZzTGlua3MiLCJpYXQiOjE1MTYyMzkwMjJ9._FLAAeNLx6Gvvpv95dELIz52xDVAE500BWutDRjsUtQ")

	alg, ok := payload["alg"]
	if ok != true {
		t.Error("Expecting map to contain `alg` key")
	}
	if ok && alg != expected {
		t.Errorf("Expected `alg` to be `%s`. Got `%s`", expected, alg)
	}
}

// #endregion

// #region Setup

func createAuthHeader() string {
	token, err := authHelper.GenerateJwt(user, 1)
	if err != nil {
		panic("Error generating JWT for unit test. " + err.Error())
	}

	return "Bearer " + token
}

// #endregion
