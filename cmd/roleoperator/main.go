package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/slok/role-operator/pkg/log"
	"github.com/slok/role-operator/pkg/operator"
)

const (
	opResync = 30 * time.Second
)

// Main is the main function.
func Main() error {
	logger := log.Base()

	cfg := operator.Config{
		ResyncDuration: opResync,
	}

	// Run the operator.
	stopC := make(chan struct{})
	errC := make(chan error)
	op := operator.New(cfg, logger)
	go func() {
		errC <- op.Run(stopC)
	}()

	// Listen os signals.
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errC:
		return err
	case <-signalC:
		logger.Infof("Signal captured, exiting...")
	}
	return nil
}

func main() {
	err := Main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running operator: %s", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
