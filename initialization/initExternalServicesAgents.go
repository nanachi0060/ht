package initialization

import (
	"httpHandler/libs"
)

type Agents struct {
	Config libs.Configuration
}

func InitExternalServicesAgents() (Agents, error) {

	return Agents{}, nil
}
