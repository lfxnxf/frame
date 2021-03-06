package cmd

var (
	_tplERRCode = `package code

import "github.com/lfxnxf/frame/BackendPlatform/golang/ecode"

var (
	InvalidParam      = ecode.New(10000)
)

func init() {
	ecode.Register(map[int]string{
		10000: "xxxx错误",
	})
}

`

	_tplHMain = `package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{.Path}}/conf"
	"{{.Path}}/server/http"
	"{{.Path}}/service"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	appS := flag.String("app", "", "App dir")
	flag.Parse()
	
	inits.Init(
		inits.ConfigPath(*configS),
	)
	
	if *appS != "" {
		inits.InitNamespace(*appS)
	}

}


func main() {

	defer inits.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	srv := service.New(cfg)

	// init and start http server
	http.Init(srv, cfg)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("{{.Name}} server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}

`

	_tplHConfig = `package conf

import (
	"github.com/lfxnxf/frame/logic/inits"
)

type Config struct {
}

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := inits.ConfigInstance().Scan(cfg)
	return cfg, err
}
`

	_tplHDao = `package dao

import (
	"context"

	"{{.Path}}/conf"
)

// Dao represents data access object
type Dao struct {
	c *conf.Config
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c: c,
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

`

	_tplHManager = `package manager

import (
	"context"

	"{{.Path}}/conf"
)

// Manager represents middleware component
// such as, kafka, http client or rpc client, etc.
type Manager struct {
	c *conf.Config
}

func New(conf *conf.Config) *Manager {
	return &Manager{
		c: conf,
	}
}

func (m *Manager) Ping(ctx context.Context) error {
	return nil
}

func (m *Manager) Close() error {
	return nil
}

`
	_tplHModel = `//Generated by the inits tool.  DO NOT EDIT!
package model

type Model struct {
}

`

	_tplHServer = `package http

import (
	"{{.Path}}/conf"
	"{{.Path}}/service"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	httpplugin "github.com/lfxnxf/frame/logic/inits/plugins/http"
)

var (
	svc *service.Service

	httpServer httpserver.Server
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = inits.HTTPServer()

	// add namespace plugin
	httpServer.Use(httpplugin.Namespace)

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()

}

func Shutdown() {
	if httpServer != nil {
		httpServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}
`
	_tplHRouter = `
// Generated by the inits tool.  DO NOT EDIT!
package http
import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)


func initRoute(s httpserver.Server) {

	s.ANY("/ping", ping)

{{.Router}}

}

`

	_tplHHandler = `package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)

func ping(c *httpserver.Context) {
	if err := svc.Ping(c.Ctx); err != nil {
		c.JSONAbort(nil, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(okMsg, nil)
}

`

	_tplHPBHandler = `//Generated by the inits tool.  DO NOT EDIT!
package http

import (
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"{{.Path}}/model"
)

func ping(c *httpserver.Context) {
	if err := svc.Ping(c.Ctx); err != nil {
		c.JSONAbort(nil, err)
		return
	}
	okMsg := map[string]string{"result": "ok"}
	c.JSON(okMsg, nil)
}

{{.Handler}}

`

	_tplHService = `package service

import (
	"context"
	"{{.Path}}/conf"
	"{{.Path}}/dao"
	"{{.Path}}/manager"
)

type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

func New(c *conf.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
	if s.mgr != nil {
		s.mgr.Close()
	}
}

`

	_tplHPBService = `package service

import (
	"context"
	"{{.Path}}/model"
	"{{.Path}}/conf"
	"{{.Path}}/dao"
	"{{.Path}}/manager"
)

var _  = model.Model{}
type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

func New(c *conf.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
	if s.mgr != nil {
		s.mgr.Close()
	}
}


{{range .Methods}}
//{{.Name}}...
//Generated by the inits tool.  IMPLEMENT YOUR OWN METHOD!
func (s *Service) {{.Name}}(ctx context.Context, request *model.{{.InputType}}) (*model.{{.OutputType}}, error) {
	return nil, nil
}
{{end}}

`
	_tplMethods = `
{{range .Methods}}
//{{.Name}}...
//Generated by the inits tool.  IMPLEMENT YOUR OWN METHOD!
func (s *Service) {{.Name}}(ctx context.Context, request *model.{{.InputType}}) (*model.{{.OutputType}}, error) {
	return nil, nil
}
{{end}}
`

	_tplGit = `
.DS_Store
.idea/*

# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# logs
*.log

logs/
data/

`

	_tplHAppToml = `
[server]
	service_name="{{.Name}}"
	port = 10000

[log]
	level="debug"
	logpath="logs"
	rotate="hour"

`
	_tplDoc = `---
title: {{.ServiceName}}
---
{{.Doc}}
<!--Generated by the inits tool. DO NOT EDIT! -->
`
)
