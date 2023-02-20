package core

import (
	"configer-service/internal/model"
	"configer-service/internal/repository"
)

type (
	ConfigServiceI interface {
		Create(cfg *model.Config) error
	}
	configService struct {
		cfgRepo repository.ConfigI
	}
)

func NewConfigService(c repository.ConfigI) configService {
	return configService{
		cfgRepo: c,
	}
}

func (c configService) Create(cfg *model.Config) error {
	return nil
}
