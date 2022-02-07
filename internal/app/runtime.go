package app

import (
	"context"
	"fmt"
	"github.com/otter-im/identity/internal/config"
	"github.com/otter-im/identity/pkg/rpc"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"log"
	mathRand "math/rand"
	"net"
	"time"
)

var (
	exitHooks = make([]func() error, 0)
)

func Init() error {
	rand.Seed(uint64(time.Now().UnixNano()))
	mathRand.Seed(time.Now().UnixNano())

	if err := checkPostgres(); err != nil {
		return err
	}

	if err := checkRedis(); err != nil {
		return err
	}

	return nil
}

func Run() error {
	listener, err := net.Listen("tcp", net.JoinHostPort(config.ServiceHost(), config.ServicePort()))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	rpc.RegisterLookupServiceServer(srv, &LookupService{
		Users: &userDto{},
	})
	log.Printf("service listening on %v\n", listener.Addr())
	if err = srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func Exit() error {
	for _, hook := range exitHooks {
		err := hook()
		if err != nil {
			return err
		}
	}
	return nil
}

func AddExitHook(hook func() error) {
	exitHooks = append(exitHooks, hook)
}

func checkPostgres() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := Postgres().Ping(ctx); err != nil {
		return fmt.Errorf("postgresql connection failure: %v", err)
	}
	return nil
}

func checkRedis() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if cmd := RedisRing().Ping(ctx); cmd.Err() != nil {
		return fmt.Errorf("redis connection failure: %v", cmd.Err())
	}
	return nil
}
