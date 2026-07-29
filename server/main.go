package main

import (
	"server/core"
	"server/flag"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	global.DB = initialize.MySQLInit()
	global.ESClient = initialize.ConnectEs()
	global.Redis = initialize.ConnectRedis()
	initialize.OtherInit()

	defer global.Redis.Close()

	flag.InitFlag()
	initialize.InitCron()

	core.RunServer()
}
