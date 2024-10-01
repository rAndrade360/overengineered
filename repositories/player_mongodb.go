package repositories

import (
	"context"
	"terraform-mongodb-pratical-example/domain"

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

func (r *PlayerMongoDBRepository) GetPlayers(ctx context.Context) ([]domain.Player, error) {
	cursor, err := r.Collection.Find(ctx, options.Find())

	players := []domain.Player{}
	err = cursor.All(ctx, &players)
	if err != nil {
		return []domain.Player{}, err
	}

	return players, nil
}
