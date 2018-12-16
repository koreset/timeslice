package services

import "github.com/koreset/timeslice/models"

type ClientService struct {
}

func (cs *ClientService) Create(client *models.Client) error {
	return GetDB().Create(client).Error
}
