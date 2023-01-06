package factory

import "github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/infra/kafka/event"


func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy{
	switch topic {
	case "chooseTeam":
		return event.ProcessChooseTeam{}
	case "newPlayer":
		return event.ProcessNewPlayer{}
	case "newMatch":
		return event.ProcessNewMatch{}
	case "newAction":
		return event.ProcessNewAction{}
	case "matchUpdateResult":
		return event.ProcessMatchUpdateResult{}
	}
	return nil
}