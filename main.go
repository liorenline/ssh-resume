package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress("localhost:2222"),
		wish.WithHostKeyPath(".ssh/host_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, "Hi, I'm Iva")
					wish.Println(sess, "This is my resume. More coming soon!")
					next(sess)
				}
			},
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started on localhost:2222")
	go s.ListenAndServe()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
