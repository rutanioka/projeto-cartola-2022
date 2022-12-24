package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/domain/entity"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/usecase"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)

type ProcessNewAction struct{}

func (p ProcessNewAction) Process (ctx context.Context, msg *kafka.Message, uow uow.UowInterface)error {
	var input usecase.ActionAddInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil{
		return err
	}
	actionTable := entity.ActionTable{}
	actionTable.Init()
	addNewActionUseCase := usecase.NewActionAddUseCase(uow, &actionTable)
	err = addNewActionUseCase.Execute(ctx, input)
	if err != nil{
		return err
	}
	return nil
}