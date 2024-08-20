package utils

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword Function
// This function takes a plain text password as input and returns its hashed version using the bcrypt algorithm.
//
// Parameters:
// - password: A string representing the plain text password to be hashed.
//
// Returns:
// - A string representing the hashed password.
// - An error if there was an issue hashing the password.
//
// Note: The bcrypt algorithm is used with a default cost factor of 10.
func HashPassword(password string) (string, error) {
	var hashpassword string

	// Generate a hashed version of the password using bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hashpassword, err
	}

	// Convert the hashed bytes to a string
	hashpassword = string(hashedBytes)

	return hashpassword, nil
}

// ConfirmPassword Function
// This function compares a given password with its hashed version.
//
// Parameters:
// - password: A string representing the password to be confirmed.
// - hashpassword: A string representing the hashed password.
//
// Returns:
//   - A boolean value indicating whether the password matches the hashed password.
//     Returns true if the password matches, false otherwise.
//   - An error if there was an issue comparing the password with the hashed password.
//
// Note: This function uses bcrypt.CompareHashAndPassword to compare the password with the hashed password.
func ConfirmPassword(password string, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))

	return err == nil
}

// GetJwtToken generates a JWT token with the provided parameters.
//
// Parameters:
// - secretKey: A string representing the secret key used for signing the token.
// - iat: An int64 representing the time at which the token was issued (in seconds since Unix epoch).
// - seconds: An int64 representing the token's lifetime in seconds.
// - userID: An int64 representing the user's ID.
// - payload: An interface{} representing additional data to be included in the token's claims.
//
// Returns:
// - A string representing the generated JWT token.
// - An error if there was an issue generating the token.
//
// Note: The token is signed using the RSA algorithm.
func GetJwtToken(secretKey string, iat, seconds int64, userID int64, payload interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userID"] = userID
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
