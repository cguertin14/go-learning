package cmd

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/cguertin14/logger"
	"github.com/spf13/cobra"
)

var (
	ctxCmd = &cobra.Command{
		Use:   "context",
		Short: "Learn the context library",
		RunE:  contextHandler,
	}
)

func init() {
	rootCmd.AddCommand(ctxCmd)
}

func contextHandler(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		cancelExample(ctx)
	}()

	go func() {
		defer wg.Done()
		timeoutExample(ctx)
	}()

	go func() {
		defer wg.Done()
		deadlineExample(ctx)
	}()

	wg.Wait()

	return nil
}

func cancelExample(ctx context.Context) {
	logger := logger.NewFromContextOrDefault(ctx)
	ctx, cancel := context.WithCancel(ctx)

	firstOperation := func(ctx context.Context) error {
		time.Sleep(200 * time.Millisecond)
		return errors.New("first operation failing")
	}

	secondOperation := func(ctx context.Context) {
		// We use a similar pattern to the HTTP server
		// that we saw in the earlier example
		select {
		case <-time.After(500 * time.Millisecond):
			logger.Info("operation 2 done")
		case <-ctx.Done():
			logger.Warn("halted operation2")
		}
	}

	go func() {
		if err := firstOperation(ctx); err != nil {
			cancel()
		}
	}()

	secondOperation(ctx)
}

func timeoutExample(ctx context.Context) {
	logger := logger.NewFromContextOrDefault(ctx)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Millisecond)
	defer cancel()

	// Trigger a deadline exceeded error on purpose while
	// performing an HTTP Get request on google.com.
	req, _ := http.NewRequest(http.MethodGet, "https://google.com", http.NoBody)
	req = req.WithContext(ctx)

	client := http.Client{}
	_, err := client.Do(req)
	if err != nil {
		// Timeout error.
		logger.Warnf("error when querying google.com: %s", err.Error())
	}
}

func deadlineExample(ctx context.Context) {
	logger := logger.NewFromContextOrDefault(ctx)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

outer:
	for {
		select {
		case <-ctx.Done():
			logger.Infof("gracefully shutting down deadline: %s", ctx.Err().Error())
			break outer
		}
	}
}
