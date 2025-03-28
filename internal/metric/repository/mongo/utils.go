package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (dr *DatabaseRepository) getServerStatus(c context.Context) (bson.M, error) {
	var result bson.M
	cmd := bson.D{
		{Key: "serverStatus", Value: 1},
	}
	err := dr.Client.Database("admin").RunCommand(c, cmd).Decode(&result)
	return result, err
}
