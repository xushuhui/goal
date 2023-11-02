package repository

import (
	"{{ .ProjectName }}/internal/model"
)

type {{ .FileName }}Repo interface {
	FirstById(id int64) (*model.{{ .FileName }}, error)
}

func New{{ .FileName }}Repo(repo *Repo) {{ .FileName }}Repo {
	return &{{ .FileNameTitleLower }}Repo{
		Repo: repo,
	}
}

type {{ .FileNameTitleLower }}Repo struct {
	*Repo
}

func (r *{{ .FileNameTitleLower }}Repo) FirstById(id int64) (*model.{{ .FileName }}, error) {
	var {{ .FileNameTitleLower }} model.{{ .FileName }}
	// TODO: query db
	return &{{ .FileNameTitleLower }}, nil
}
