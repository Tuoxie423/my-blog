package flag

import (
	"errors"
	"fmt"
	"os"
	"server/global"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initializes the srtucture of the MySQL database table.",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "Exports the MySQL database table structure to a .sql file.",
	}
	sqlImportFlag = &cli.StringFlag{
		Name:  "sql-import",
		Usage: "Imports SQL data from a specified file.",
	}
	esFlag = &cli.BoolFlag{
		Name:  "es",
		Usage: "Initializes the Elasticsearch index.",
	}
	esExportFlag = &cli.BoolFlag{
		Name:  "es-export",
		Usage: "Exports data from Elasticsearch to a specified file.",
	}
	esImportFlag = &cli.StringFlag{
		Name:  "es-import",
		Usage: "Imports data into Elasticsearch from a specified file.",
	}
	adminFlag = &cli.BoolFlag{
		Name:  "admin",
		Usage: "Creates an administrator using the name, email and address specified in the config.yaml file.",
	}
)

func Run(c *cli.Context) {
	// 检查是否设置了多个标志
	if c.NumFlags() > 1 {
		err := cli.NewExitError("只能传一个标志", 1)
		global.Log.Error("标志参数数量过多:", zap.Error(err))
		os.Exit(1)
	}
	switch {
	case c.Bool(sqlFlag.Name):
		// 初始化数据库表结构
		if err := SQL(); err != nil {
			global.Log.Error("数据库表结构迁移失败：", zap.Error(err))
		} else {
			global.Log.Info("数据库表结构迁移成功!")
		}
	case c.Bool(sqlExportFlag.Name):
		// 导出数据库表结构到 .sql 文件
		if err := SQLExport(); err != nil {
			global.Log.Error("导出数据库表结构失败：", zap.Error(err))
		} else {
			global.Log.Info("数据库表结构导出成功!")
		}
	case c.IsSet(sqlImportFlag.Name):
		// 导入 SQL 数据
		if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
			var combinedErrors string
			for _, err := range errs {
				combinedErrors += err.Error() + "\n"
			}
			err := errors.New(combinedErrors)
			global.Log.Error("Failed to import SQL data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully imported SQL data")
		}
	case c.Bool(esFlag.Name):
		// 初始化 Elasticsearch 索引
		if err := Elasticsearch(); err != nil {
			global.Log.Error("Elasticsearch 索引初始化失败：", zap.Error(err))
		} else {
			global.Log.Info("Elasticsearch 索引初始化成功!")
		}
	case c.Bool(esExportFlag.Name):
		if err := ElasticsearchExport(); err != nil {
			global.Log.Error("Failed to export ES data:", zap.Error(err))
		} else {
			global.Log.Info("Successfully exported ES data")
		}
	case c.IsSet(esImportFlag.Name):
		if num, err := ElasticsearchImport(c.String(esImportFlag.Name)); err != nil {
			global.Log.Error("Failed to import ES data:", zap.Error(err))
		} else {
			global.Log.Info(fmt.Sprintf("Successfully imported ES data, totaling %d records", num))
		}
	case c.Bool(adminFlag.Name):
		// 创建管理员
		if err := Admin(); err != nil {
			global.Log.Error("创建管理员失败：", zap.Error(err))
		} else {
			global.Log.Info("Successfully created an administrator")
		}
	default:
		err := cli.NewExitError("无效的参数", 1)
		global.Log.Error("无效的参数：", zap.Error(err))
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Go Blog"
	app.Flags = []cli.Flag{
		sqlFlag,
		sqlExportFlag,
		sqlImportFlag,
		esFlag,
		esExportFlag,
		esImportFlag,
		adminFlag,
	}
	app.Action = Run
	return app
}

func InitFlag() {
	if len(os.Args) > 1 {
		app := NewApp()
		err := app.Run(os.Args)
		if err != nil {
			global.Log.Error("flag 运行失败：", zap.Error(err))
			os.Exit(1)
		}
		if os.Args[1] == "-h" || os.Args[1] == "-help" {
			fmt.Println("显示帮助信息...")
		}
		os.Exit(0)
	}
}
