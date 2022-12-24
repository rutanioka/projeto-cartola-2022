package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/usecase"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)


type ProcessChooseTeam struct{}

func (p ProcessChooseTeam) Process (ctx context.Context, msg *kafka.Message, uow uow.UowInterface)error {
	var input usecase.MyTeamChoosePlayersInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil{
		return err
	}
	addNewTeamUseCase := usecase.NewMyTeamChoosePlayersUseCase(uow)
	err = addNewTeamUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}