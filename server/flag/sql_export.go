package flag

import (
	"fmt"
	"os"
	"os/exec"
	"server/global"
)

func SQLExport() error {
	timer := "20260726"
	sqlCfg := global.Config.Mysql
	sqlpath := fmt.Sprintf("mysql_%s.sql", timer)
	cmd := exec.Command("docker", "exec", "mysql", "mysqldump", "-u"+sqlCfg.Username, "-p"+sqlCfg.Password, sqlCfg.DBName)

	outFile, err := os.Create(sqlpath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	return cmd.Run()
}
