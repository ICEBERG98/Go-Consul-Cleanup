package client

import (
	"strconv"
	"strings"

	"github.com/ICEBERG98/go-consul-cleanup/pkg/config"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	log "github.com/sirupsen/logrus"
)

var DefaultQueryOptions = api.QueryOptions{
	Namespace:         "",
	Datacenter:        "",
	AllowStale:        false,
	RequireConsistent: false,
	UseCache:          false,
	MaxAge:            0,
	StaleIfError:      0,
	WaitIndex:         0,
	WaitHash:          "",
	WaitTime:          0,
	Token:             "",
	Near:              "",
	NodeMeta:          nil,
	RelayFactor:       0,
	LocalOnly:         false,
	Connect:           false,
	Filter:            "",
}

func CreateClientForAllNodes() map[*api.Node]*api.Client {
	log.Info("Creating Client For all nodes.")
	client := createClientDefault()
	Nodes, _, _ := client.Catalog().Nodes(&DefaultQueryOptions)
	Clients := make(map[*api.Node]*api.Client)
	Datacenter := config.Config.Bootstrap.Datacenter
	for _, Node := range Nodes {
		if Node.Datacenter != Datacenter {
			log.Warningf("Node at Address- %s doesnt have the DC set to %s", Node.Address, Datacenter)
		} else {
			log.Infof("Attempting To Create New Client for Node at Address- %s", Node.Address)
			newClient := CreateClientForNode(Node.Address)
			if newClient == nil {
				log.Warningf("Client wasn't Created for Node- %s, Wont be Added to ClientObject.", Node.Address)
			} else {
				Clients[Node] = newClient
			}
		}
	}
	return Clients
}

func CreateClientForNode(nodeIp string) *api.Client {
	if nodeIp == "" {
		return createClientDefault()
	}
	log.Infof("Attempting NewClient Creation for Node with Ip- %s", nodeIp)
	client, err := api.NewClient(setupNewConfig(nodeIp))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Infof("Client Successfully Created for Node- %s", nodeIp)
	return client
}
func createClientDefault() *api.Client {
	BootStrapNode := config.Config.Bootstrap.Node
	log.Infof("Creating Default Client config- %s:%s", BootStrapNode.Address, strconv.Itoa(BootStrapNode.Port))
	consulNode := BootStrapNode.Address + ":" + strconv.Itoa(BootStrapNode.Port)
	client, err := api.NewClient(setupNewConfig(consulNode))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Infof("Successfully Created Default Client")
	return client
}

// This method is used as wrapper/Default mapper over the inbuilt api.Config.
// Eliminates the need to add Environment variables.
// Defaults- The following Params for Our Usage here-
// consul Node Address, Scheme to HTTPS and SSL_verify False.
func setupNewConfig(consulNode string) *api.Config {
	log.Infof("Setting Up New Configuration For Node- %s", consulNode)
	if strings.ContainsAny(consulNode, ":") == false {
		log.Infof("Port Not Specified in consulNode, Appending Default Port- %s",
			strconv.Itoa(config.Config.Defaults.Port))
		consulNode = consulNode + ":" + strconv.Itoa(config.Config.Defaults.Port)
	}
	configuration := &api.Config{
		Address:   consulNode,
		Scheme:    "https",
		Transport: cleanhttp.DefaultPooledTransport(),
	}
	configuration.TLSConfig.InsecureSkipVerify = true
	log.Debugf("Config For Node- %s is %#v", consulNode, configuration)
	return configuration
}
