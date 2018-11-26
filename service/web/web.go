package web

import (
	"context"
	"github.com/dreampuf/invoker/service/log"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"net/http"
)

type Context = gin.Context

type WebService struct {
	engine *gin.Engine
	Addr string
}

var (
	 box = packr.New("box", "./static")
	 indexRender = &PackrRender{box, "index.html"}
)

func NewWebService() *WebService {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", Index)
	engine.GET("/ws", WebSocket)
	engine.StaticFS("/static/", box)
	return &WebService{
		engine: engine,
		Addr: "localhost:8081",
	}
}

func (w *WebService) Serve(ctx context.Context) error {
	serv := http.Server{Addr: w.Addr, Handler: w.engine}
	go func() {
		select {
		case <- ctx.Done():
			err := serv.Shutdown(ctx)
			if err != nil {
				log.WithError(err).Error("web service shutdown error")
			}
		}
	}()
	if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func Index(c *Context) {
	c.Render(http.StatusOK, indexRender)
}
