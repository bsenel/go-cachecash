package provider

import (
	"context"
	"net"

	"github.com/kelleyk/go-cachecash/ccmsg"
	"github.com/kelleyk/go-cachecash/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// An Application is the top-level content provider.  It takes a configuration struct.  Its children are the several
// protocol servers (that deal with clients, caches, and so forth).
type Application interface {
	common.StarterShutdowner
}

// XXX: Right now, this is shared between the client- and cache-facing servers.
type Config struct {
	ClientProtocolAddr string
	CacheProtocolAddr  string
}

func (c *Config) FillDefaults() {
	if c.ClientProtocolAddr == "" {
		c.ClientProtocolAddr = ":8080"
	}
	if c.CacheProtocolAddr == "" {
		c.CacheProtocolAddr = ":8082"
	}
}

type application struct {
	l *logrus.Logger

	clientProtocolServer *clientProtocolServer
	cacheProtocolServer  *cacheProtocolServer
	// TODO: ...
}

var _ Application = (*application)(nil)

// XXX: Should this take p as an argument, or be responsible for setting it up?
func NewApplication(l *logrus.Logger, p *ContentProvider, conf *Config) (Application, error) {
	conf.FillDefaults()

	clientProtocolServer, err := newClientProtocolServer(l, p, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client protocol server")
	}

	cacheProtocolServer, err := newCacheProtocolServer(l, p, conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cache protocol server")
	}

	return &application{
		l:                    l,
		clientProtocolServer: clientProtocolServer,
		cacheProtocolServer:  cacheProtocolServer,
	}, nil
}

func (a *application) Start() error {
	if err := a.clientProtocolServer.Start(); err != nil {
		return errors.Wrap(err, "failed to start client protocol server")
	}
	if err := a.cacheProtocolServer.Start(); err != nil {
		return errors.Wrap(err, "failed to start cache protocol server")
	}
	return nil
}

func (a *application) Shutdown(ctx context.Context) error {
	return a.clientProtocolServer.Shutdown(ctx)
}

type clientProtocolServer struct {
	l          *logrus.Logger
	conf       *Config
	provider   *ContentProvider
	grpcServer *grpc.Server
}

var _ common.StarterShutdowner = (*clientProtocolServer)(nil)

func newClientProtocolServer(l *logrus.Logger, p *ContentProvider, conf *Config) (*clientProtocolServer, error) {
	grpcServer := grpc.NewServer()
	ccmsg.RegisterClientProviderServer(grpcServer, &grpcClientProviderServer{provider: p})

	return &clientProtocolServer{
		l:          l,
		conf:       conf,
		provider:   p,
		grpcServer: grpcServer,
	}, nil
}

func (s *clientProtocolServer) Start() error {
	s.l.Info("clientProtocolServer - Start - enter")

	lis, err := net.Listen("tcp", s.conf.ClientProtocolAddr)
	if err != nil {
		return errors.Wrap(err, "failed to bind listener")
	}

	go func() {
		// This will block until we call `Stop`.
		if err := s.grpcServer.Serve(lis); err != nil {
			s.l.WithError(err).Error("failed to serve clientProtocolServer")
		}
	}()

	s.l.Info("clientProtocolServer - Start - exit")
	return nil
}

func (s *clientProtocolServer) Shutdown(ctx context.Context) error {
	// TODO: Should use `GracefulStop` until context expires, and then fall back on `Stop`.
	s.grpcServer.Stop()

	return nil
}

type cacheProtocolServer struct {
	l          *logrus.Logger
	conf       *Config
	provider   *ContentProvider
	grpcServer *grpc.Server
}

var _ common.StarterShutdowner = (*cacheProtocolServer)(nil)

func newCacheProtocolServer(l *logrus.Logger, p *ContentProvider, conf *Config) (*cacheProtocolServer, error) {
	grpcServer := grpc.NewServer()
	ccmsg.RegisterCacheProviderServer(grpcServer, &grpcCacheProviderServer{provider: p})

	return &cacheProtocolServer{
		l:          l,
		conf:       conf,
		provider:   p,
		grpcServer: grpcServer,
	}, nil
}

func (s *cacheProtocolServer) Start() error {
	s.l.Info("cacheProtocolServer - Start - enter")

	lis, err := net.Listen("tcp", s.conf.CacheProtocolAddr)
	if err != nil {
		return errors.Wrap(err, "failed to bind listener")
	}

	go func() {
		// This will block until we call `Stop`.
		if err := s.grpcServer.Serve(lis); err != nil {
			s.l.WithError(err).Error("failed to serve cacheProtocolServer")
		}
	}()

	s.l.Info("cacheProtocolServer - Start - exit")
	return nil
}

func (s *cacheProtocolServer) Shutdown(ctx context.Context) error {
	// TODO: Should use `GracefulStop` until context expires, and then fall back on `Stop`.
	s.grpcServer.Stop()

	return nil
}
