package service

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"ums/src/models"
	"ums/src/models/users"
)

type UserManagement struct{}

func (u *UserManagement) Create(ctx echo.Context) error {
	user := new(users.User)
	userRequest := new(users.UserRequestDTO)

	if err := ctx.Bind(userRequest); err != nil {
		return err
	}

	user.RequestDtoToObject(*userRequest)
	insertOneResult, err := models.UserInfoDatabase().InsertOne(context.TODO(), user)
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

func (u *UserManagement) GetUserById(userObjID primitive.ObjectID) users.User {
	var user users.User

	findUser := models.UserInfoDatabase().FindOne(context.TODO(), bson.M{"_id": userObjID})
	userInfo := new(users.User)
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
	findOneAndDelete := models.UserInfoDatabase().FindOneAndDelete(context.TODO(), bson.M{"_id": userObjID})

	if findOneAndDelete.Err() != nil {
		return findOneAndDelete.Err()
	}

	return ctx.JSON(http.StatusOK, "Deleted successfully!")

}

func (u *UserManagement) GetAll(ctx echo.Context) error {
	findUser, err := models.UserInfoDatabase().Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	var allUsers []users.UserResponse
	for findUser.Next(context.Background()) {
		var user users.UserResponse
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

	userRequestDTO := new(users.UserRequestDTO)
	err = ctx.Bind(userRequestDTO)
	if err != nil {
		return ctx.JSON(http.StatusOK, "[ERROR]: Operation Failed 2")
	}

	user := new(users.User)
	user.RequestDtoToObject(*userRequestDTO)
	updateData := bson.M{"$set": user}
	result := models.UserInfoDatabase().FindOneAndUpdate(context.TODO(), bson.M{"_id": userID}, updateData)
	if result.Err() != nil {
		return ctx.JSON(http.StatusOK, "[ERROR]: Operation Failed 3 ")
	}

	return ctx.JSON(http.StatusOK, "Updated Successfully")
}
