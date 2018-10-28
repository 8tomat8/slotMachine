package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/8tomat8/slotMachine/apiHTTP"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	cfg := Config{}
	envconfig.MustProcess(prefix, &cfg)

	var (
		l   *zap.Logger
		err error
	)
	if cfg.Environment == "dev" {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init logger"))
	}
	logger := l.Sugar()
	defer l.Sync()

	// Start API
	srv := apiHTTP.NewServer(logger.With("package", "api"), cfg.APIaddr, []byte(cfg.JWTSecret))

	handleSignals(logger, srv)
}

func handleSignals(logger *zap.SugaredLogger, toClose ...io.Closer) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Waiting for first signal
	<-sigs

	done := make(chan struct{})
	go func() {
		for _, e := range toClose {
			err := e.Close()
			if err != nil {
				logger.Warn(errors.Wrap(err, "failed to close gracefully"))
			}
		}
		close(done)
	}()

	select {
	case <-sigs:
		// If second signal was received, terminating application
		logger.Warn("Application was terminated")
	case <-done:
		logger.Info("Application was gracefully stopped")
	}
}

const prefix = "SLOTS"

type Config struct {
	JWTSecret   string `envconfig:"JWT_SECRET" default:"Foo-Bar-Baz-42" required:"true"`
	APIaddr     string `envconfig:"API_ADDR" default:"0.0.0.0:8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"prod"`
}
