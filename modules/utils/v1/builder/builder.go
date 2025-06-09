package builder

import (
	"go.uber.org/fx"

	"starter-go-gin/app"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/modules/utils/v1/handler"
	"starter-go-gin/modules/utils/v1/service"
	"starter-go-gin/sdk/fcm"
	"starter-go-gin/sdk/gcs"
	"starter-go-gin/utils"
)

var UtilsModule = fx.Options(
	fx.Provide(
		utils.NewRedisClient,
		// The code below is commented out because synchronization with Firestore is no longer needed
		// fx.Annotate(
		// 	firestore.NewFirestore,
		// 	fx.As(new(interfaces.FirestoreUseCase)),
		// ),
		fx.Annotate(
			fcm.NewFCM,
			fx.As(new(interfaces.FCMUseCase)),
		),
		fx.Annotate(
			service.NewUtilsCreator,
			fx.As(new(service.UtilsCreatorUseCase)),
		),
		gcs.NewGoogleCloudStorage,
		handler.NewUtilsCreatorHandler,
	),
	fx.Invoke(app.UtilsCreatorHTTPHandler),
)
