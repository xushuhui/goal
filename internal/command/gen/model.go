package gen

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var CmdGen = &cobra.Command{
	Use:     "gen [model-name] ",
	Short:   "Gen new model",
	Example: "goal gen user",
	Args:    cobra.ExactArgs(1),
	Run:     run,
}
var dsn string
type Gen struct {
	Model string `survey:"model"`
}
func init() {
	CmdGen.Flags().StringVarP(&dsn, "dsn", "d", dsn, "dsn")
}

func run(cmd *cobra.Command, args []string) {
	g := &Gen{

	}
	if len(args) == 0 {
		err := survey.AskOne(&survey.Input{
			Message: "What is your gen model?",
			Help:    "model name.",
			Suggest: nil,
		}, &dsn, survey.WithValidator(survey.Required))
		if err != nil {
			return
		}
	} else {
		g.Model = args[0]
	}
	cfg := gen.Config{
		OutPath: "./internal/data/model",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	}
	cfg.WithDataTypeMap(map[string]func(gorm.ColumnType) (dataType string){
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			return "int32"
		},
	})
	generator := gen.NewGenerator(cfg)
	gormdb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("open db error: %s", err.Error())
	}
	generator.UseDB(gormdb)
	if g.Model == "all" {
		generator.GenerateAllTable()
	} else {
		generator.GenerateModel(g.Model )
	}

	generator.Execute()
}
