package di

import (
	"my-go-template/src/adapter/controller"
	"my-go-template/src/adapter/repository"
	"my-go-template/src/driver/db"
	"my-go-template/src/usecase"
)

func GetSampleController() controller.SampleControllerInterface {
	return controller.NewSampleController(
		usecase.NewGetSampleUsecase(
			repository.NewSampleRepository(
				db.NewMysqlSQLHandler(),
			),
		),
	)
}

func GetSampleEchoController() controller.SampleEchoControllerInterface {
	return controller.NewSampleEchoController(
		usecase.NewGetSampleUsecase(
			repository.NewSampleRepository(
				db.NewMysqlSQLHandler(),
			),
		),
	)
}
