package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/app"
	"starter-go-gin/modules/master/v1/handler"
	"starter-go-gin/modules/master/v1/repository"
	"starter-go-gin/modules/master/v1/service"
)

var MasterModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			repository.NewShiftRepository,
			fx.As(new(repository.ShiftRepositoryUseCase)),
		),
		fx.Annotate(
			repository.NewShiftDetailRepository,
			fx.As(new(repository.ShiftDetailRepositoryUseCase)),
		),
		fx.Annotate(
			service.NewMasterCreator,
			fx.As(new(service.MasterCreatorUseCase)),
		),
		fx.Annotate(
			service.NewMasterFinder,
			fx.As(new(service.MasterFinderUseCase)),
		),
		fx.Annotate(
			service.NewMasterUpdater,
			fx.As(new(service.MasterUpdaterUseCase)),
		),
		fx.Annotate(
			service.NewMasterDeleter,
			fx.As(new(service.MasterDeleterUseCase)),
		),
		handler.NewMasterFinderHandler,
		handler.NewMasterCreatorHandler,
		handler.NewMasterDeleterHandler,
		handler.NewMasterUpdaterHandler,
	),
	fx.Invoke(app.MasterFinderHTTPHandler),
)
