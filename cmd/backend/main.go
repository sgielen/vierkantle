package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sgielen/vierkantle/pkg/database"
	"github.com/sgielen/vierkantle/pkg/email"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Strings []string

func (s *Strings) String() string {
	return strings.Join(*s, ";")
}

func (s *Strings) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var (
	port       = flag.Uint("port", 9001, "Listen on this TCP port")
	boards     = flag.String("boards", "./boards", "Path to the boards")
	wordLists  Strings
	bonusLists Strings
)

type vierkantleService struct {
	log      *log.Logger
	teams    VierkantleTeams
	boardDir string

	wordLists  []string
	bonusLists []string
}

func main() {
	flag.Var(&wordLists, "words", "Read words from this word list")
	flag.Var(&bonusLists, "bonus", "Read bonus words from this word list")
	flag.Parse()

	database.Init()

	if b, err := os.Stat(*boards); err != nil || !b.IsDir() {
		log.Printf("Warning: failed to stat boards directory %s, won't offer boards to clients", *boards)
	}

	if email.GetSendgridApiKey() == "" {
		log.Printf("Warning: no Sendgrid API key found, will be unable to send e-mail")
	}

	// Set up a gRPC server with interceptors for e.g. authentication and panic prevention
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterVierkantleServiceServer(grpcServer,
		&vierkantleService{
			log:        log.New(os.Stderr, "vierkantle: ", log.LstdFlags),
			teams:      *NewVierkantleTeams(),
			boardDir:   *boards,
			wordLists:  wordLists,
			bonusLists: bonusLists,
		},
	)
	grpclog.SetLogger(log.New(os.Stderr, "grpc: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
	)
	handler := func(resp http.ResponseWriter, req *http.Request) {
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
