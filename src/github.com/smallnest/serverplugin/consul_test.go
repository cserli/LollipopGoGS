// +build consul

package serverplugin

import (
	"testing"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
)

func TestConsulRegistry(t *testing.T) {
	s := server.NewServer()

	r := &ConsulRegisterPlugin{
		ServiceAddress: "tcp@127.0.0.1:8972",
		ConsulServers:  []string{"127.0.0.1:8500"},
		BasePath:       "/rpcx_test",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		t.Fatal(err)
	}
	s.Plugins.Add(r)

	s.RegisterName("Arith", new(Arith), "")
	go s.Serve("tcp", "127.0.0.1:8972")
	defer s.Close()

	if len(r.Services) != 1 {
		t.Fatal("failed to register services in consul")
	}

}
