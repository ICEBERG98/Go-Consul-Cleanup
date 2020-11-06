package pkg

import (
	"github.com/ICEBERG98/go-consul-cleanup/pkg/agent"
	"github.com/ICEBERG98/go-consul-cleanup/pkg/client"
	log "github.com/Sirupsen/logrus"
	"github.com/hashicorp/consul/api"
	"strconv"
)

func getFailingHealthCheckFromAgent(consulAgent *api.Agent) []*api.AgentCheck {

	// Checks := agent.GetChecksForAgent(agent)
	Checks := agent.GetChecksForAgent(consulAgent)
	DeRegistrableServices := make([]*api.AgentCheck, 0)
	for serviceId, AgentCheck := range Checks {
		log.Infof("Checking Service with id- %s, Status - %s", serviceId, AgentCheck.Status)
		if AgentCheck.Status != "passing" {
			log.Warningf("Service with id- %s, and Name- %s has %s Check, Marking this service for Deletion",
				AgentCheck.ServiceID, AgentCheck.Status, AgentCheck.ServiceName)
			DeRegistrableServices = append(DeRegistrableServices, AgentCheck)
		}

	}
	return DeRegistrableServices
}

func getRemovableServicesFromAgent(currentAgent *api.Agent) []*api.AgentService {
	RemovableServices := make([]*api.AgentService, 0)
	ServiceMaps := make(map[string]string)
	Services, err := currentAgent.Services()
	if err != nil {
		log.Panic(err)
	}
	for _, Service := range Services {
		Address := Service.Address
		Port := strconv.Itoa(Service.Port)
		CheckTarget := Address + ":" + Port
		_, exists := ServiceMaps[CheckTarget]
		if exists {
			log.Infof("Redundant Service found for Address- %s, Marking %s for deletion", CheckTarget, Service.ID)
			RemovableServices = append(RemovableServices, Service)
		} else {
			ServiceMaps[CheckTarget] = Service.ID
		}
	}
	return RemovableServices
}

func DeRegisterRedundantThings() {
	clients := client.CreateClientForAllNodes()
	consulAgents := agent.GetAgentsForMultiClient(clients)
	deRegistrableItemsForAllAgents := make(map[*api.Agent][]string)
	for _, consulAgent := range consulAgents {
		FailingHealthChecks := getFailingHealthCheckFromAgent(consulAgent)
		deRegistrableItemsForThisAgent := make([]string, 0)
		for _, Check := range FailingHealthChecks {
			deRegistrableItemsForThisAgent = append(deRegistrableItemsForThisAgent, Check.ServiceID)
		}
		removableServices := getRemovableServicesFromAgent(consulAgent)
		for _, service := range removableServices {
			deRegistrableItemsForThisAgent = append(deRegistrableItemsForThisAgent, service.ID)
		}
		deRegistrableItemsForAllAgents[consulAgent] = deRegistrableItemsForThisAgent
	}
	for currAgent, services := range deRegistrableItemsForAllAgents {
		for _, service := range services {
			agent.DeregisterServiceById(service, currAgent)
		}
	}
}
