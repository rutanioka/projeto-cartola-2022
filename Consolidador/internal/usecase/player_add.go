package usecase

import (
	"context"

	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/domain/entity"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/domain/entity/repository"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)

type AddPlayerInput struct {
	ID string
	Name string
	InitialPrice float64
}

type AddPlayerUseCase struct {
	Uow uow.UowInterface
}

func NewAddPlayerUseCase( uow uow.UowInterface) *AddPlayerUseCase  {
	return &AddPlayerUseCase{
		Uow: uow,
	}
}

func (a *AddPlayerUseCase) Execute (ctx context.Context, input AddPlayerInput) error {
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
	err := playerRepository.Create(ctx, player)
	if err != nil{
		return err
	}
	a.Uow.CommitOrRollback()
	return nil
}

func (a *AddPlayerUseCase) getPlayerRepository (ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx , "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}