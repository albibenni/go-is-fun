package helpers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/albibenni/go-exercises/auth/config"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

var jwtKey []byte

func SetJWTKey(key string) {
	jwtKey = []byte(key)
}

func GetJWTKey() []byte {
	return []byte(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {

	secretKey := GetJWTKey()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateTokens(email, userID, userType string) (string, string) {
	log.Printf("JWT Key %v Type: %T", jwtKey, jwtKey)

	//Token expiration times
	tokenExpiry := time.Now().Add(24 * time.Hour).Unix()
	refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix()

	claims := &Claims{
		Email:  email,
		UserID: userID,
		Role:   userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(tokenExpiry, 0)),
		},
	}

	refreshClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(refreshTokenExpiry, 0)),
		},
	}

	//Generate tokens
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAccessToken, err := accessToken.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}

	return signedAccessToken, signedRefreshToken
}

func HashPassword(password *string) *string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	hashedPwd := string(bytes)
	return &hashedPwd
}

func UpdateAllTokens(signedToken, signedRefreshToken, userID string, db *config.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	userCollection := config.GetUserCollection(db)

	updateObj := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "token", Value: signedToken},
			{Key: "refresh_token", Value: signedRefreshToken},
			{Key: "updated_at", Value: time.Now()},
		}},
	}

	// filter
	filter := bson.M{"user_id": userID}

	_, err := userCollection.UpdateOne(ctx, filter, updateObj)

	return err
}

func VerifyPassword(foundPwd, pwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(foundPwd), []byte(pwd))

	return err == nil, err
}
