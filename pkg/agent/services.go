package agent

import (
	"github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

func DeregisterServiceById(ServiceID string, agent *api.Agent) {
	err := agent.ServiceDeregister(ServiceID)
	if err != nil {
		logrus.Errorf("Unable to Deregister Service with Id- %s, Error- %s", ServiceID, err)
	} else {
		logrus.Infof("Succesfully Deregistered Service with Id- %s", ServiceID)
	}
	return
}
