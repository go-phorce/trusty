package mockpb

import (
	"context"

	"github.com/ekspand/trusty/api/v1/pb"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

// MockRAServer for testing
type MockRAServer struct {
	pb.RAServiceServer

	Reqs []proto.Message

	// If set, all calls return this error.
	Err error

	// responses to return if err == nil
	Resps []proto.Message
}

// SetResponse sets a single response without errors
func (m *MockRAServer) SetResponse(r proto.Message) {
	m.Err = nil
	m.Resps = []proto.Message{r}
}

// SetError sets an error response
func (m *MockRAServer) SetError(err error) {
	m.Err = err
}

// GetRoots returns the root CAs
func (m *MockRAServer) GetRoots(context.Context, *empty.Empty) (*pb.RootsResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.RootsResponse), nil
}

// RegisterRoot registers root CA
func (m *MockRAServer) RegisterRoot(ctx context.Context, in *pb.RegisterRootRequest) (*pb.RootsResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.RootsResponse), nil
}

// RegisterCertificate registers certificate
func (m *MockRAServer) RegisterCertificate(ctx context.Context, in *pb.RegisterCertificateRequest) (*pb.CertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificateResponse), nil
}

// GetCertificate returns certificate
func (m *MockRAServer) GetCertificate(ctx context.Context, in *pb.GetCertificateRequest) (*pb.CertificateResponse, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Resps[0].(*pb.CertificateResponse), nil
}
