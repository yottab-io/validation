package validation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MongoIndexValidate(id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	return err
}
