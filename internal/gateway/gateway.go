package gateway

import (
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"net/http"
	"simple-grpc-crud/config"
	"time"
)

type GrpcGateway struct {
	grpcGateway *grpcweb.WrappedGrpcServer
}

func NewGrpcGateway(grpcServer *grpc.Server) *GrpcGateway {
	options := []grpcweb.Option{
		grpcweb.WithOriginFunc(makeHttpOriginFunc()),
	}

	grpcGateway := grpcweb.WrapServer(grpcServer, options...)

	return &GrpcGateway{
		grpcGateway: grpcGateway,
	}
}

func (s *GrpcGateway) Start(ctx context.Context, cfg *config.Config) error {
	r := http.NewServeMux()
	r.Handle("/", s.grpcGateway)

	addr := getAddressString(cfg)

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
	}
	defer srv.Shutdown(ctx)

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func getAddressString(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s", cfg.GrpcGatewayAddr, cfg.GrpcGatewayPort)
}

func makeHttpOriginFunc() func(origin string) bool {
	return func(origin string) bool {
		return true
	}
}
