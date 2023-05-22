package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	"ums/src/models"
	"ums/src/utilities/common"
)

type UserManagement struct {
	Collection models.UserModel
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (u *UserManagement) Create(ctx echo.Context) error {
	user := new(models.User)
	userRequest := new(models.UserRequestDTO)

	err := ctx.Bind(userRequest)

	if err != nil {
		print(err)
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Something wrong"})
	}

	err = common.LocalValidator.Struct(userRequest)
	if err != nil {
		print(err.Error())
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Operation failed"})
	}
	checkIfExist := u.GetByEmail(userRequest.Email)

	if checkIfExist.ID != "" {
		print("EXIST")
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Email already exist"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		print(err)
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Something wrong"})
	}

	userRequest.Password = string(hashedPassword)
	user.RequestDtoToObject(*userRequest)
	insertOneResult, err := u.Collection.Get().InsertOne(context.TODO(), user)
	if err != nil {
		return ctx.JSON(http.StatusOK, "[ERROR]: "+err.Error())
	}

	return ctx.JSON(http.StatusOK, insertOneResult.InsertedID)
}

func (u *UserManagement) Get(ctx echo.Context) error {
	userID := ctx.Param("id")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	userInfo := u.GetUserById(userObjID)
	return ctx.JSON(http.StatusOK, userInfo)
}

func (u *UserManagement) GetUserById(userObjID primitive.ObjectID) models.User {
	var user models.User

	findUser := u.Collection.Get().FindOne(context.TODO(), bson.M{"_id": userObjID})
	userInfo := new(models.User)
	err := findUser.Decode(userInfo)
	if err != nil {
		println("[ERROR]", err)
		return user
	}
	user = *userInfo
	return user
}

func (u *UserManagement) Delete(ctx echo.Context) error {
	userID := ctx.Param("id")

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	findOneAndDelete := u.Collection.Get().FindOneAndDelete(context.TODO(), bson.M{"_id": userObjID})

	if findOneAndDelete.Err() != nil {
		return findOneAndDelete.Err()
	}

	return ctx.JSON(http.StatusOK, "Deleted successfully!")

}

func (u *UserManagement) GetAll(ctx echo.Context) error {
	findUser, err := u.Collection.Get().Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	var allUsers []models.UserResponse
	for findUser.Next(context.Background()) {
		var user models.UserResponse
		if err := findUser.Decode(&user); err != nil {
			return err
		}
		allUsers = append(allUsers, user)
	}
	return ctx.JSON(http.StatusOK, allUsers)
}

func (u *UserManagement) Update(ctx echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusOK, "[ERROR]: Operation Failed 1")
	}

	userRequestDTO := new(models.UserRequestDTO)
	err = ctx.Bind(userRequestDTO)

	if err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Operation Failed 2"})
	}

	err = common.LocalValidator.Struct(userRequestDTO)
	if err != nil {
		print(err.Error())
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Operation failed"})
	}

	user := new(models.User)
	user.RequestDtoToObject(*userRequestDTO)
	updateData := bson.M{"$set": user}
	result := u.Collection.Get().FindOneAndUpdate(context.TODO(), bson.M{"_id": userID}, updateData)
	if result.Err() != nil {
		return ctx.JSON(http.StatusOK, "[ERROR]: Operation Failed 3 ")
	}

	return ctx.JSON(http.StatusOK, "Updated Successfully")
}

func (u *UserManagement) GetByEmail(email string) models.UserResponse {
	var res models.UserResponse
	findOne := u.Collection.Get().FindOne(context.TODO(), bson.M{"email": email})
	userInfo := new(models.UserResponse)
	err := findOne.Decode(userInfo)
	if err != nil {
		log.Println("[ERROR]", err)
		return res
	}
	res = *userInfo
	return res
}

func Login(ctx echo.Context) error {
	//var user models.User
	userLoginRequest := new(models.UserAuthDTO)

	if err := ctx.Bind(userLoginRequest); err != nil {
		return err
	}

	u := new(UserManagement)
	getUser := u.GetByEmail(userLoginRequest.Email)
	if getUser.Email == "" {
		return ctx.JSON(http.StatusOK, "User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(userLoginRequest.Password))

	if err != nil {
		print(err)
		return ctx.JSON(http.StatusOK, "Password does not match")
	}

	// Get payload in jwt
	claims := &jwtCustomClaims{
		getUser.Name,
		getUser.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func MiddlewareControl() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	return config
}
