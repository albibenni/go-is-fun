package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/albibenni/go-exercises/auth/config"
	"github.com/albibenni/go-exercises/auth/helpers"
	"github.com/albibenni/go-exercises/auth/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func Signup(db *config.DB) gin.HandlerFunc {
	userCollection := config.GetUserCollection(db)
	return func(ginCtx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := ginCtx.BindJSON(&user); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Validate user input
		if validationErr := validate.Struct(user); validationErr != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{
			"$or": []bson.M{
				{"email": user.Email},
				{"phone": user.Phone},
			},
		})

		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone already exists"})
			return
		}

		//Generate rest of the user data
		user.Password = helpers.HashPassword(user.Password)
		user.Created_at = time.Now()
		user.Updated_at = time.Now()
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		accessToken, refreshToken := helpers.GenerateTokens(*user.Email, user.User_id, *user.Role)
		user.Token = &accessToken
		user.Refresh_token = &refreshToken

		_, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		ginCtx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

func Login(db *config.DB) gin.HandlerFunc {
	userCollection := config.GetUserCollection(db)
	return func(ginCtx *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := ginCtx.BindJSON(&user); err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
			return
		}

		passwordIsValid, msg := helpers.VerifyPassword(*foundUser.Password, *user.Password)
		if !passwordIsValid {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		token, refreshToken := helpers.GenerateTokens(*foundUser.Email, *&foundUser.User_id, *foundUser.Role)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id, db)

		ginCtx.JSON(http.StatusOK, gin.H{
			"user":          foundUser,
			"token":         token,
			"refresh_token": refreshToken,
		})
	}
}

func GetUser(db *config.DB) gin.HandlerFunc {
	userCollection := config.GetUserCollection(db)
	return func(ginCtx *gin.Context) {
		requestedUserId := ginCtx.Param("id")

		claims, exists := ginCtx.Get("claims")
		if !exists {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Type assertion to get the claims object
		tokenClaims, ok := claims.(*helpers.Claims)
		if !ok {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		tokenUserId := tokenClaims.UserID
		userType := tokenClaims.Role

		if userType != "ADMIN" && tokenUserId != requestedUserId {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": requestedUserId}).Decode(&user)
		if err != nil {
			ginCtx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ginCtx.JSON(http.StatusOK, user)
	}
}

func GetUsers(db *config.DB) gin.HandlerFunc {
	userCollection := config.GetUserCollection(db)
	return func(ginCtx *gin.Context) {
		claims, exists := ginCtx.Get("claims")
		if !exists {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Type assertion to get the claims object
		tokenClaims, ok := claims.(*helpers.Claims)
		if !ok {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		if tokenClaims.Role != "ADMIN" {
			ginCtx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		cursor, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		var users []models.User
		if err := cursor.All(ctx, &users); err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the list of users
		ginCtx.JSON(http.StatusOK, users)
	}
}
