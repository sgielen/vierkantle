package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port = flag.Uint("port", 9001, "Listen on this TCP port")
)

func main() {
	flag.Parse()

	grpcServer := grpc.NewServer()
	pb.RegisterVierkantleServiceServer(grpcServer, &vierkantleService{})
	grpclog.SetLogger(log.New(os.Stderr, "vierkantleServer: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
	)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("got request to %s", req.URL.String())
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.StripPrefix("/api", http.HandlerFunc(handler)),
	}

	grpclog.Infof("Starting server. http port: %d", *port)
	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
