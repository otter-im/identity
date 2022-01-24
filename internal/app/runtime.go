package app

import (
	"context"
	"github.com/golang/glog"
	"github.com/otter-im/identity-service/internal/config"
	"github.com/otter-im/identity-service/pkg/rpc"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	mathRand "math/rand"
	"net"
	"time"
)

var (
	exitHooks = make([]func(ctx context.Context) error, 0)
)

func Init() {
	rand.Seed(uint64(time.Now().UnixNano()))
	mathRand.Seed(time.Now().UnixNano())
}

func Run() error {
	listener, err := net.Listen("tcp", net.JoinHostPort(config.ServiceHost(), config.ServicePort()))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	rpc.RegisterLookupServiceServer(srv, &LookupService{})
	glog.Infof("service listening on %v", listener.Addr())
	if err = srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func Exit(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	default:
	}

	for _, hook := range exitHooks {
		err := hook(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddExitHook(hook func(ctx context.Context) error) {
	exitHooks = append(exitHooks, hook)
}
