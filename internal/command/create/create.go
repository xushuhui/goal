package create

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/xushuhui/goal/internal/pkg/helper"
	"github.com/xushuhui/goal/tpl"
)

type Create struct {
	ProjectName        string
	CreateType         string
	FilePath           string
	FileName           string
	FileNameTitleLower string
	FileNameFirstChar  string
	IsFull             bool
}

func NewCreate() *Create {
	return &Create{}
}

var CmdCreate = &cobra.Command{
	Use:     "create [type] [handler-name]",
	Short:   "Create a new handler/service/repo",
	Example: "goal create handler user",
	Args:    cobra.ExactArgs(2),
	Run:     func(cmd *cobra.Command, args []string) {},
}

var (
	tplPath string
)

func init() {
	CmdCreateHandler.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateService.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateRepo.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")

	CmdCreateAll.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
}

var CmdCreateHandler = &cobra.Command{
	Use:     "handler",
	Short:   "Create a new handler",
	Example: "goal create handler user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateService = &cobra.Command{
	Use:     "service",
	Short:   "Create a new service",
	Example: "goal create service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateRepo = &cobra.Command{
	Use:     "repo",
	Short:   "Create a new repo",
	Example: "goal create repo user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

var CmdCreateAll = &cobra.Command{
	Use:     "all",
	Short:   "Create a new handler & service & repo ",
	Example: "goal create all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])

	switch c.CreateType {
	case "handler", "service", "repo":
		c.genFile()

	case "all":

		c.CreateType = "handler"
		c.genFile()

		c.CreateType = "service"
		c.genFile()

		c.CreateType = "repo"
		c.genFile()

	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}
}

func (c *Create) genFile() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/", c.CreateType)
	}
	if c.CreateType == "repo" {
		filePath = "internal/data/"
	}
	f := createFile(filePath, strings.ToLower(c.FileName)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.FileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	var t *template.Template
	var err error
	if tplPath == "" {
		t, err = template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", c.CreateType))
	} else {
		t, err = template.ParseFiles(path.Join(tplPath, fmt.Sprintf("%s.tpl", c.CreateType)))
	}
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.FileName)+".go")
}

func createFile(dirPath string, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
