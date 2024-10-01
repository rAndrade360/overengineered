package repositories

import (
	"context"
	"terraform-mongodb-pratical-example/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerMongoDBRepository struct {
	Collection *mongo.Collection
}

func (r *PlayerMongoDBRepository) SavePlayer(ctx context.Context, p *domain.Player) error {
	_, err := r.Collection.InsertOne(ctx, p, options.InsertOne())
	return err
}

func (r *PlayerMongoDBRepository) GetPlayers(ctx context.Context, page, perPage int) ([]domain.Player, error) {
	limit := int64(perPage)
	skip := int64((page - 1) * perPage)

	cursor, err := r.Collection.Find(ctx, bson.D{{}}, &options.FindOptions{Limit: &limit, Skip: &skip})

	players := []domain.Player{}
	err = cursor.All(ctx, &players)
	if err != nil {
		return []domain.Player{}, err
	}

	return players, nil
}
