package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/usecase"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)


type ProcessNewPlayer struct{}

func (p ProcessNewPlayer) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface)error {
	var input usecase.AddPlayerInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil{
		return err
	}
	return nil
}