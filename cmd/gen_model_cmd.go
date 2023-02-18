package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/helloword/dao/model"
	"github.com/wuqinqiang/helloword/db"
	"gorm.io/gen"
)

var GenCmd = &cli.Command{
	Name:  "gen",
	Usage: "offline deals",
	Action: func(context *cli.Context) error {
		// specify the output directory (default: "./query")
		// ### if you want to query without context constrain, set mode gen.WithoutContext ###
		outPath := "dao/query"
		g := gen.NewGenerator(gen.Config{
			OutPath: outPath,
			OutFile: outPath + "/query.go",
			/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
			//if you want the nullable field generation property to be pointer type, set FieldNullable true
			/* FieldNullable: true,*/
			//If you need to generate index tags from the database, set FieldWithIndexTag true
			/* FieldWithIndexTag: true,*/
		})

		// reuse the database connection in Project or create a connection here
		// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
		//db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/boost?charset=utf8&parseTime=True&loc=Local"))
		g.UseDB(db.Get())
		g.GenerateAllTable()
		g.ApplyBasic(model.Word{}, model.WordPhrase{}, model.Phrase{}, model.WordPhraseUsage{})
		g.Execute()
		return nil
	},
}
