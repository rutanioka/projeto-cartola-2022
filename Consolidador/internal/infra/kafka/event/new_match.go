package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/usecase"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)


type ProcessNewMatch struct {}

func (p ProcessNewMatch) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface)error {
	var input usecase.MatchInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil{
		return err
	}
	addNewMatchUseCase := usecase.NewMatchUseCase(uow)
	err = addNewMatchUseCase.Execute(ctx, input)
	if err != nil{
		return err
	}
	return nil
}