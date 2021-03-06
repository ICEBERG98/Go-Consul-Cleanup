package agent

import (
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func GetAgentByClient(client *api.Client) *api.Agent {
	agent := client.Agent()
	return agent
}

func GetAgentsForMultiClient(clients map[*api.Node]*api.Client) map[*api.Node]*api.Agent {
	agents := make(map[*api.Node]*api.Agent)
	for Node, client := range clients {
		newAgent := GetAgentByClient(client)
		if newAgent == nil {
			log.Warningf("Agent is nil for client- %#v", client)
		}
		if newAgent != nil {
			agents[Node] = newAgent
		}
	}
	return agents
}
