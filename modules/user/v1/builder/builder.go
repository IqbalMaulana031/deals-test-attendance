package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/common/interfaces"
	"starter-go-gin/modules/user/v1/repository"
	"starter-go-gin/modules/user/v1/service"
	"starter-go-gin/sdk/excelize"
	"starter-go-gin/sdk/gotenberg"
)

var UserModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			gotenberg.NewGotenberg,
			fx.As(new(interfaces.GotenbergUseCase)),
		),
		fx.Annotate(
			excelize.NewExcelize,
			fx.As(new(interfaces.ExcelizeUseCase)),
		),
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(repository.UserRepositoryUseCase)),
		),
		fx.Annotate(
			repository.NewRoleRepository,
			fx.As(new(repository.RoleRepositoryUseCase)),
		),
		fx.Annotate(
			repository.NewUserRoleRepository,
			fx.As(new(repository.UserRoleRepositoryUseCase)),
		),
		fx.Annotate(
			service.NewUserCreator,
			fx.As(new(service.UserCreatorUseCase)),
		),
		fx.Annotate(
			service.NewUserFinder,
			fx.As(new(service.UserFinderUseCase)),
		),
		fx.Annotate(
			service.NewUserUpdater,
			fx.As(new(service.UserUpdaterUseCase)),
		),
		fx.Annotate(
			service.NewUserDeleter,
			fx.As(new(service.UserDeleterUseCase)),
		),
	),
)
