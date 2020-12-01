package service_test

import (
	"testing"

	"github.com/ekspand/trusty/backend/service/auth"
	"github.com/ekspand/trusty/backend/service/ca"
	"github.com/ekspand/trusty/backend/service/status"
	"github.com/ekspand/trusty/backend/service/workflow"
	"github.com/ekspand/trusty/backend/trustyserver"
)

var serviceFactories = map[string]trustyserver.ServiceFactory{
	auth.ServiceName:     auth.Factory,
	ca.ServiceName:       ca.Factory,
	status.ServiceName:   status.Factory,
	workflow.ServiceName: workflow.Factory,
}

func Test_invalidArgs(t *testing.T) {
	for _, f := range serviceFactories {
		testInvalidServiceArgs(t, f)
	}
}

func testInvalidServiceArgs(t *testing.T, f trustyserver.ServiceFactory) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected panic but didn't get one")
		}
	}()
	f(nil)
}
