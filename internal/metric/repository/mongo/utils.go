package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) getCommand(c context.Context, cmd bson.D) (bson.M, error) {
	var result bson.M
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	return result, err
}
