package data

import (
	"{{ .ProjectName }}/internal/data/model"
)

type {{ .FileName }}Repo interface {
	FirstById(id int64) (*model.{{ .FileName }}, error)
}

func New{{ .FileName }}Repo(data *Data) {{ .FileName }}Repo {
	return &{{ .FileNameTitleLower }}Repo{
		Data: data,
	}
}

type {{ .FileNameTitleLower }}Repo struct {
	*Data
}

func (r *{{ .FileNameTitleLower }}Repo) FirstById(id int64) (*model.{{ .FileName }}, error) {
	var {{ .FileNameTitleLower }} model.{{ .FileName }}
	// TODO: query db
	return &{{ .FileNameTitleLower }}, nil
}
