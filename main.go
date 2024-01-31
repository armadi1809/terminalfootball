/*
Copyright © 2024 Aziz Rmadi azizrmadi@gmail.com
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	terminalfootballcmd "github.com/armadi1809/terminalfootball/cmd"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
	"github.com/joho/godotenv"
)

const (
	host = "localhost"
	port = 23236
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd := terminalfootballcmd.GetRoot()
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			func(h ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {

					cmd.SetArgs(s.Command())
					cmd.SetIn(s)
					cmd.SetOut(s)
					cmd.SetErr(s.Stderr())
					cmd.CompletionOptions.DisableDefaultCmd = true
					if err := cmd.Execute(); err != nil {
						_ = s.Exit(1)
						return
					}

					h(s)
				}
			},
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("could not stop server", "error", err)
	}
}
