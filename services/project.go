package services

import "github.com/koreset/timeslice/models"

type ProjectService struct {
}

func (ps *ProjectService) Create(project *models.Project) error {
	return GetDB().Create(project).Error
}
