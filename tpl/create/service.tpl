package service

import (
	"{{ .ProjectName }}/internal/model"
	"{{ .ProjectName }}/internal/data"
)

type {{ .FileName }}Service interface {
	Get{{ .FileName }}ById(id int64) (*model.{{ .FileName }}, error)
}

func New{{ .FileName }}Service(service *Service, {{ .FileNameTitleLower }}Repo data.{{ .FileName }}Repo) {{ .FileName }}Service {
	return &{{ .FileNameTitleLower }}Service{
		Service:        service,
		{{ .FileNameTitleLower }}Repo: {{ .FileNameTitleLower }}Repo,
	}
}

type {{ .FileNameTitleLower }}Service struct {
	*Service
	{{ .FileNameTitleLower }}Repo data.{{ .FileName }}Repo
}

func (s *{{ .FileNameTitleLower }}Service) Get{{ .FileName }}ById(id int64) (*model.{{ .FileName }}, error) {
	return s.{{ .FileNameTitleLower }}Repo.FirstById(id)
}
