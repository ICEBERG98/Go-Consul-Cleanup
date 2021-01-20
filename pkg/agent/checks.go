package agent

import (
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func GetChecksForAgent(agent *api.Agent) map[string]*api.AgentCheck {
	checks, err := agent.Checks()
	if err != nil {
		log.Errorf("Error while getting checks- %s", err)
		return nil
	}
	return checks
}
