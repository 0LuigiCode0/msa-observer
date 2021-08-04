package main

import (
	core_helper "x-msa-core/helper"
	"x-msa-observer/core"
	"x-msa-observer/helper"

	"github.com/0LuigiCode0/logger"
)

func main() {
	conf := &helper.Config{}
	if err := core_helper.ParseConfig(helper.ConfigDir+helper.ConfigFile, conf); err != nil {
		logger.Log.Errorf("config parse invalid: %v", err)
		core_helper.Wg.Wait()
		return
	}
	srv, err := core.InitServer(conf)
	if err != nil {
		logger.Log.Errorf("server not initialized: %v", err)
		srv.Close()
		core_helper.Wg.Wait()
		return
	}
	if err := srv.Start(); err != nil {
		logger.Log.Errorf("server not started: %v", err)
		srv.Close()
		core_helper.Wg.Wait()
		return
	}
	srv.Close()
	core_helper.Wg.Wait()
}
