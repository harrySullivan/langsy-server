package controllers

import (
	"App/database/models"
	"App/http/services"
	"App/utils"
	
	"os"
	"fmt"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Controller
}

type AppClaims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func sign(claims *AppClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseOrError(tokenString string) (*primitive.ObjectID, error) {
	claims := &AppClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*AppClaims); ok && token.Valid {
		oid, oidErr := primitive.ObjectIDFromHex(claims.UserID)
		return &oid, oidErr
	} else {
		return nil, err
	}
}

func hash(password string) (string, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(byteHash), err
}

func verifyHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func GetUserID(c *gin.Context) *primitive.ObjectID {
	userOid, exists := c.Get("UserID")
	if !exists {
		fmt.Println("UserID does not exist")
		err := errors.New("UserID does not exist")
		if utils.HttpError(c, 500, err) { return nil }
	}
	userOidRef := (userOid.(*primitive.ObjectID))
	return userOidRef
}

func (AuthController) SignUp(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if utils.HttpError(c, 400, err) { return }

	passwordHash, hashErr := hash(user.Password)
	if utils.HttpError(c, 500, hashErr) { return }

	user.Password = passwordHash

	insertedUser, insertErr := services.UserService{}.Insert(&user)
	if utils.HttpError(c, 500, insertErr) { return }

	claims := &AppClaims{
		UserID: insertedUser.ID.Hex(),
	}

	jwt, jwtErr := sign(claims)
	if utils.HttpError(c, 500, jwtErr) { return }

	c.JSON(200, gin.H{
		"Token": jwt,
	})
}

func (AuthController) Login(c *gin.Context) {
	var loginCred models.LoginCred
	err := c.Bind(&loginCred)
	if utils.HttpError(c, 400, err) { return }

	foundUser, findErr := services.UserService{}.FindByUsername(loginCred.Username)
	if utils.HttpError(c, 400, findErr) { return }

	hashErr := verifyHash(loginCred.Password, foundUser.Password)
	if utils.HttpError(c, 400, hashErr) { return }

	claims := &AppClaims{
		UserID: foundUser.ID.Hex(),
	}

	jwt, jwtErr := sign(claims)
	if utils.HttpError(c, 500, jwtErr) { return }

	c.JSON(200, gin.H{
		"Token": jwt,
	})
}
