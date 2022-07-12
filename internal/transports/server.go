package transports

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rviewer-Challenges/4TOWna2EWttcHFagNbUw/api/internal/service"
	"github.com/Rviewer-Challenges/4TOWna2EWttcHFagNbUw/api/internal/storage"
)

// WebServer has the logic to start the microservice
type WebServer struct {
}

// StartServer listens and servers this microservice
func (ws *WebServer) StartServer() {
	var httpAddr = flag.String("http", ":8081", "http listen address")

	flag.Parse()

	ctx := context.Background()

	storage, err := storage.New()
	if err != nil {
		log.Fatal("error creating the storage instance", err)
	}

	service := service.New(storage)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("listening on port", *httpAddr)
		handler := MakeHTTPHandler(ctx, service)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Println("Server ends ", <-errs)

}
