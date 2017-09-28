package controllers

import (
  "crypto/rsa"
  "errors"
  "io/ioutil"
  "log"
  "strconv"

    "github.com/dgrijalva/jwt-go"
)

// RwtToken is an Struct to encapsulate username and expires as parameter
type RwtToken struct {
	ID       uint
	Expires  int64
}

// https://gist.github.com/cryptix/45c33ecf0ae54828e63b
// location of the files used for signing and verification
const (
	privKeyPath = "keys/rsakey.pem"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/rsakey.pem.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey    *rsa.PublicKey
	mySigningKey *rsa.PrivateKey
)

func init() {
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)

	if err != nil {
		log.Fatal(err)
	}

	mySigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
}

// GetToken is a function that exposes the method to get a simple token for jw
func (e *RwtToken) getToken() (string, error) {

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires, //time.Unix(c.ExpiresAt, 0)
		Issuer:    strconv.Itoa(int(e.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}

func (e *RwtToken) validToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New("token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if token == nil {
		log.Println(err)
		return false, errors.New("not work")
	}

	if token.Valid {
		//"You look nice today"
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return false, errors.New("Timing is everything")
		} else {
			//"Couldn't handle this token:"
			return false, err
		}
	} else {
		//"Couldn't handle this token:"
		return false, err
	}
}
