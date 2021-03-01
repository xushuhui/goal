package project

import (
	"context"
	"fmt"
	"github.com/xushuhui/goal/internal/base"

	"os"
	"path"
)

const (
	serviceLayoutURL = "https://github.com/xushuhui/goal-layout.git"
)

// Project is a project template.
type Project struct {
	Name string
}

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", p.Name)
	}
	fmt.Printf("Creating project %s\n", p.Name)
	repo := base.NewRepo(serviceLayoutURL)

	if err := repo.CopyTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
		return err
	}
	os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Name),
	)
	return nil
}
