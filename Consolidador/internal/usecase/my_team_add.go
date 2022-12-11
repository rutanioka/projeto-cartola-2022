package usecase

import (
	"context"

	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/domain/entity"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/internal/domain/entity/repository"
	"github.com/rutanioka/Projeto-Cartola-2022/Consolidador/package/uow"
)

type AddMyTeamInput struct{
	ID string
	Name string
	Score int
}

type AddMyTeamUseCase struct{
	Uow uow.UowInterface
}

func (a *AddMyTeamUseCase) Execute (ctx context.Context, input AddMyTeamInput) error{
	myTeamRepository := a.getMyTeamRepository( ctx)
	myTeam := entity.NewMyTeam(input.ID,input.Name)
	err := myTeamRepository.Create(ctx, myTeam)
	if err != nil{
		return err
	}
	return a.Uow.CommitOrRollback()
}

func (a *AddMyTeamUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}