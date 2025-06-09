package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/app"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/modules/auth/v1/handler"
	"starter-go-gin/modules/auth/v1/repository"
	"starter-go-gin/modules/auth/v1/service"
	"starter-go-gin/sdk/gcs"
)

var AuthModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			repository.NewAuthRepository,
			fx.As(new(repository.AuthRepositoryUseCase)),
		),
		fx.Annotate(
			gcs.NewGoogleCloudStorage,
			fx.As(new(interfaces.CloudStorageUseCase)),
		),
		fx.Annotate(
			service.NewAuthService,
			fx.As(new(service.AuthUseCase)),
		),
		handler.NewAuthHandler,
	),
	fx.Invoke(app.AuthHTTPHandler),
)
