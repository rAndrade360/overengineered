package services

import (
	"context"
	"terraform-mongodb-pratical-example/domain"
	"terraform-mongodb-pratical-example/repositories"
)

type PlayerService struct {
	Repository *repositories.PlayerMongoDBRepository
}

func (s *PlayerService) SavePlayer(ctx context.Context, p *domain.Player) error {
	return s.Repository.SavePlayer(ctx, p)
}

func (s *PlayerService) GetPlayers(ctx context.Context, page, perPage int) ([]domain.Player, error) {
	return s.Repository.GetPlayers(ctx, page, perPage)
}
