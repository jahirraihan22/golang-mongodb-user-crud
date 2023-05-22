package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ums/src/models"
)

type Collection interface {
	Get() *mongo.Collection
}

func User() Collection {
	return &models.UserModel{}
}
