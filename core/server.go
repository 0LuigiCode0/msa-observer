package core

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	core_helper "x-msa-core/helper"
	"x-msa-observer/core/database"
	"x-msa-observer/helper"
	"x-msa-observer/hub"

	"github.com/0LuigiCode0/logger"
)

type Server interface {
	Start() error
	Close()
}

type server struct {
	srv http.Server
	hub hub.Hub
	db  database.DB
}

func InitServer(conf *helper.Config) (S Server, err error) {
	s := &server{}
	S = s
	s.db, err = database.InitDB(conf)
	if err != nil {
		s.srv.Close()
		err = fmt.Errorf("db not initialized: %v", err)
		return
	}
	s.hub, err = hub.InitHub(s.db, conf)
	if err != nil {
		err = fmt.Errorf("hub not initialized: %v", err)
		return
	}
	s.srv.Handler = s.hub.GetHandler()
	s.srv.Addr = fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	logger.Log.Service("server initialized")
	return
}

func (s *server) Start() error {
	signal.Notify(core_helper.C, os.Interrupt)

	core_helper.Wg.Add(1)
	go s.loop()

	core_helper.Wg.Add(1)
	go func() {
		defer core_helper.Wg.Done()
		if err := s.srv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Log.Service("serve stoped")
				core_helper.C <- os.Interrupt
				return
			}
			logger.Log.Errorf("serve error: %v", err)
			core_helper.C <- os.Interrupt
			return
		}
	}()

	logger.Log.Service("server started at address:", s.srv.Addr)
	<-core_helper.C
	return nil
}

func (s *server) loop() {
	defer core_helper.Wg.Done()
	for {
		select {
		case <-core_helper.Ctx.Done():
			return
		default:
			if err := recover(); err != nil {
				logger.Log.Errorf("critical damage: %v", err)
			}
		}
	}
}

func (s *server) Close() {
	s.srv.Shutdown(core_helper.Ctx)
	if s.hub != nil {
		s.hub.Close()
	}
	if s.db != nil {
		s.db.Close()
	}
	core_helper.CloseCtx()
}
