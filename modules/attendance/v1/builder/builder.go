package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/app"
	"starter-go-gin/modules/attendance/v1/handler"
	"starter-go-gin/modules/attendance/v1/repository"
	"starter-go-gin/modules/attendance/v1/service"
)

var AttendanceModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			repository.NewAttendanceRepository,
			fx.As(new(repository.AttendanceRepositoryUseCase)),
		),
		fx.Annotate(
			service.NewAttendanceFinder,
			fx.As(new(service.AttendanceFinderUseCase)),
		),
		fx.Annotate(
			service.NewAttendanceCreator,
			fx.As(new(service.AttendanceCreatorUseCase)),
		),
		fx.Annotate(
			service.NewAttendanceUpdater,
			fx.As(new(service.AttendanceUpdaterUseCase)),
		),
		handler.NewAttendanceFinderHandler,
		handler.NewAttendanceCreatorHandler,
		handler.NewAttendanceDeleterHandler,
		handler.NewAttendanceUpdaterHandler,
	),
	fx.Invoke(app.AttendanceCreatorHTTPHandler),
)
