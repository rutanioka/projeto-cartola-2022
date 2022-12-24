package event

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)


type ProcessEventStrategy interface{
	Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error
}