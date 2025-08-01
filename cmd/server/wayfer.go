package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctx, finish := signal.NotifyContext(ctx, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGKILL)
	defer finish()

	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Running application...")
}

func run(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	}
}
