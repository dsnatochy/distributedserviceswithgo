package loadbalance_test

import (
	"github.com/dsnatochy/proglog/internal/config"
	"github.com/dsnatochy/proglog/internal/loadbalance"
	"github.com/dsnatochy/proglog/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	"net"
	"testing"

	// https://www.yellowduck.be/posts/assert-vs-require-in-testify/
	// "With require, as soon as you run the tests, the first requirement which fails interrupts and fails the complete test.
	// So, if the first requirement fails, the rest of the test will be skipped."
	"github.com/stretchr/testify/require"
	api "github.com/dsnatochy/proglog/api/v1"
)

func TestResolver(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	tlsConfig, err := config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.ServerCertFile,
		KeyFile:       config.ServerKeyFile,
		CAFile:        config.CAFile,
		ServerAddress: "127.0.0.1",
		Server:        true,
	})
	require.NoError(t, err)
	serverCreds := credentials.NewTLS(tlsConfig)

	srv, err := server.NewGRPCServer(&server.Config{
		GetServerer: &getServers{},
	}, grpc.Creds(serverCreds))
	require.NoError(t, err)

	go srv.Serve(l)

	conn := &clientConn{}
	tlsConfig, err = config.SetupTLSConfig(config.TLSConfig{
		CertFile:      config.RootClientCertFile,
		KeyFile:       config.RootClientKeyFile,
		CAFile:        config.CAFile,
		ServerAddress: "127.0.0.1",
		Server:        false,
	})
	require.NoError(t, err)
	clientCreds := credentials.NewTLS(tlsConfig)
	opts := resolver.BuildOptions{
		DialCreds: clientCreds,
	}
	r := &loadbalance.Resolver{}
	_, err = r.Build(
		resolver.Target{
			Endpoint: l.Addr().String(),
		},
		conn,
		opts,
	)
	require.NoError(t, err)

	wantState := resolver.State{
		Addresses:     []resolver.Address{
			{
				Addr: "localhost:9001",
				Attributes: attributes.New("is_leader", true),
			},{
				Addr: "localhost:9002",
				Attributes: attributes.New("is_leader", false),
			},
		},
	}
	require.Equal(t, wantState, conn.state)

	conn.state.Addresses = nil
	r.ResolveNow(resolver.ResolveNowOptions{})
	require.Equal(t, wantState, conn.state)
}

type getServers struct{}

func (s *getServers) GetServers() ([]*api.Server, error) {
	return []*api.Server{
		{
			Id: "leader",
			RpcAddr: "localhost:9001",
			IsLeader: true,
		}, {
			Id: "follower",
			RpcAddr: "localhost:9002",
		},
	}, nil
}

var _ resolver.ClientConn = (*clientConn)(nil)

type clientConn struct {
	resolver.ClientConn
	state resolver.State
}

func (c *clientConn) NewAddress(addresses []resolver.Address) {}

func (c *clientConn) NewServiceConfig(serviceConfig string) {}

func (c *clientConn) ParseServiceConfig(config string) *serviceconfig.ParseResult {
	return nil
}

func (c *clientConn) UpdateState(state resolver.State)  {
	c.state = state
}
func (c *clientConn) ReportError(err error) {}
