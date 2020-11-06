package agent

import (
	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

func GetChecksForAgent(agent *api.Agent) map[string]*api.AgentCheck {
	checks, err := agent.Checks()
	if err != nil {
		log.Errorf("Error while getting checks- %s", err)
		return nil
	}
	return checks
}
