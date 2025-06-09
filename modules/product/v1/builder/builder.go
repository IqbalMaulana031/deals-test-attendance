package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/app"
	"starter-go-gin/modules/product/v1/handler"
	"starter-go-gin/modules/product/v1/repository"
	"starter-go-gin/modules/product/v1/service"
)

var ProductModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			repository.NewProductRepository,
			fx.As(new(repository.ProductRepositoryUseCase)),
		),
		fx.Annotate(
			service.NewProductCreator,
			fx.As(new(service.ProductCreatorUseCase)),
		),
		fx.Annotate(
			service.NewProductFinder,
			fx.As(new(service.ProductFinderUseCase)),
		),
		fx.Annotate(
			service.NewProdcutUpdater,
			fx.As(new(service.ProductUpdaterUseCase)),
		),
		fx.Annotate(
			service.NewProductDeleter,
			fx.As(new(service.ProductDeleterUseCase)),
		),
		handler.NewProductCreatorHandler,
		handler.NewProductFinderHandler,
		handler.NewProductDeleterHandler,
	),
	fx.Invoke(app.ProductCreatorHTTPHandler),
)
