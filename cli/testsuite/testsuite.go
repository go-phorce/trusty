package testsuite

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/go-phorce/dolly/ctl"
	"github.com/go-phorce/dolly/testify/servefiles"
	"github.com/go-phorce/dolly/xpki/cryptoprov"
	"github.com/go-phorce/trusty/api/v1/trustypb"
	"github.com/go-phorce/trusty/cli"
	"github.com/go-phorce/trusty/client"
	"github.com/go-phorce/trusty/tests/mockpb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

const projFolder = "../../"

// Suite defines helper test suite
type Suite struct {
	suite.Suite
	// Server is file based server
	Server *servefiles.Server
	// ServerURL  is the current end-point
	ServerURL string
	// Out is the outpub buffer
	Out bytes.Buffer
	// Cli is the current CLI
	Cli *cli.Cli

	MockStatus *mockpb.MockStatusServer

	appFlags       []string
	withGRPC       bool
	withFileServer bool
	withHSM        bool
}

// WithGRPC specifies to use gRPC server options
func (s *Suite) WithGRPC() *Suite {
	s.withGRPC = true
	return s
}

// WithFileServer specifies to create file-based mock HTTP server
func (s *Suite) WithFileServer() *Suite {
	s.withFileServer = true
	return s
}

// WithHSM specifies to use HSM provider
func (s *Suite) WithHSM() *Suite {
	cryptoprov.Register("SoftHSM", cryptoprov.Crypto11Loader)
	s.withHSM = true
	return s
}

// WithAppFlags specifies application flags, default: -V -d
func (s *Suite) WithAppFlags(appFlags []string) *Suite {
	s.appFlags = appFlags
	return s
}

// HasText is a helper method to assert that the out stream contains the supplied
// text somewhere
func (s *Suite) HasText(texts ...string) {
	outStr := s.Output()
	for _, t := range texts {
		s.True(strings.Index(outStr, t) >= 0, "Expecting to find text %q in value %q", t, outStr)
	}
}

// HasNoText is a helper method to assert that the out stream does contains the supplied
// text somewhere
func (s *Suite) HasNoText(texts ...string) {
	outStr := s.Output()
	for _, t := range texts {
		s.True(strings.Index(outStr, t) < 0, "Expecting to NOT find text %q in value %q", t, outStr)
	}
}

// HasFile is a helper method to assert that file exists
func (s *Suite) HasFile(file string) {
	stat, err := os.Stat(file)
	s.Require().NoError(err, "File must exist: %s", file)
	s.Require().False(stat.IsDir())
}

// HasTextInFile is a helper method to assert that file contains the supplied text
func (s *Suite) HasTextInFile(file string, texts ...string) {
	f, err := ioutil.ReadFile(file)
	s.Require().NoError(err, "Unable to read: %s", file)
	outStr := string(f)
	for _, t := range texts {
		s.True(strings.Index(outStr, t) >= 0, "Expecting to find text %q in file %q", t, file)
	}
}

// Run is a helper to run a CLI commnd
func (s *Suite) Run(action ctl.ControlAction, p interface{}) error {
	s.Out.Reset()
	return action(s.Cli, p)
}

// Output returns the current CLI output
func (s *Suite) Output() string {
	return s.Out.String()
}

// SetupSuite to set up the tests
func (s *Suite) SetupSuite() {
	app := ctl.NewApplication("cliapp", "test")
	app.UsageWriter(&s.Out)

	flags := []string{"cliapp", "-V", "-D"}

	var ops []cli.Option

	if s.withFileServer || s.withGRPC {
		ops = append(ops, cli.WithServer(""))
	}
	if s.withHSM {
		ops = append(ops, cli.WithHsmCfg())
	}

	if s.withFileServer {
		s.Server = servefiles.New(s.T())
		s.Server.SetBaseDirs("testdata")
		s.ServerURL = strings.Replace(s.Server.URL(), "127.0.0.1", "localhost", 1)
	}

	if s.withFileServer || s.withGRPC {
		cfg, err := filepath.Abs(filepath.Join(projFolder, "etc/dev/trusty-config.json"))
		s.Require().NoError(err)

		flags = append(flags, "-s", s.ServerURL, "--cfg", cfg)
		ops = append(ops, cli.WithServiceCfg())
	}

	if len(s.appFlags) > 0 {
		flags = append(flags, s.appFlags...)
	}

	s.Cli = cli.New(&ctl.ControlDefinition{
		App:       app,
		Output:    &s.Out,
		ErrOutput: &s.Out,
	}, ops...)

	s.Cli.Parse(flags)
	s.Cli.PopulateControl()

	if s.withFileServer {
		err := s.Cli.EnsureClient()
		s.Require().NoError(err)
	}
}

// TearDownSuite to clean up resources
func (s *Suite) TearDownSuite() {
	if s.Server != nil {
		s.Server.Close()
	}
}

var nextPort = int32(31234)

// SetupMockGRPC for testing
func (s *Suite) SetupMockGRPC() *grpc.Server {
	serv := grpc.NewServer()
	trustypb.RegisterStatusServer(serv, s.MockStatus)

	addr := fmt.Sprintf("localhost:%d", atomic.AddInt32(&nextPort, 1))
	lis, err := net.Listen("tcp", addr)
	s.Require().NoError(err)

	client, err := client.NewFromURL(lis.Addr().String())
	s.Require().NoError(err)
	s.Cli.WithClient(client)

	go serv.Serve(lis)

	return serv
}
