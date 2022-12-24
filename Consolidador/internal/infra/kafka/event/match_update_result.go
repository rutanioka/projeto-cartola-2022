package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/usecase"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)


type ProcessMatchUpdateResult struct {}

func (p ProcessMatchUpdateResult) Process (ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MatchUpdateResultInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil{
		return err
	}
	fmt.Println("input", input)
	updateMatchResultUseCase := usecase.NewMatchUpdateResultUseCase(uow)
	err = updateMatchResultUseCase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}