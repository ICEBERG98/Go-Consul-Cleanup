package tests

import (
	"github.com/ICEBERG98/go-consul-cleanup/pkg/client"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"net"
	"testing"
)

func TestCreateClientForNode(t *testing.T) {
	testConsulNode, err := net.LookupHost("consul.service.consul")
	if err != nil {
		err, ok := err.(*net.DNSError)
		if ok {
			t.Skipf("Consul.service.consul not Resolved,error- %s, Skipping Test", err)
		}
		t.Errorf("Error while Trying to resolve Consul.service.consul")
	}
	testConsulNodeAddress := testConsulNode[0] + ":8600"
	config := &api.Config{
		Address:   testConsulNodeAddress,
		Scheme:    "https",
		Transport: cleanhttp.DefaultPooledTransport(),
	}
	config.TLSConfig.InsecureSkipVerify = true
	_, err = api.NewClient(config)
	if err != nil {
		t.Skipf("Some error with creating new client in consul Library")
	}
	got := client.CreateClientForNode(testConsulNodeAddress)
	if got == nil {
		t.Fail()
	}
}
