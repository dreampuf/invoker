package main

import (
	"context"
	"github.com/dreampuf/invoker/service/log"
	"github.com/dreampuf/invoker/service/web"
	"github.com/oklog/run"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var g run.Group

	ctx, cancel := context.WithCancel(context.Background())

	g.Add(func() error {
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)
		select {
		case <- term:
			cancel()
			log.Info("shutdown server...")
		}
		signal.Stop(term)
		return nil
	}, func(error) {
		cancel()
	})

	g.Add(func() error {
		w := web.NewWebService()
		return w.Serve(ctx)
	}, func(err error) {
		if err != nil {
			log.WithError(err).Error()
		}
		log.Info("web service exited")
	})


	if err := g.Run(); err != nil {
		log.WithError(err).Error("exiting with an error")
	}

}