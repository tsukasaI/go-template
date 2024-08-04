package usecase

import (
	"errors"
	"my-go-template/src/adapter/repository"
	customErrors "my-go-template/src/core/errors"
	"my-go-template/src/usecase/result"
)

type (
	getSampleUsecase struct {
		ur repository.SampleRepositoryInterface
	}
	GetSampleUsecaseInterface interface {
		Execute() (*result.GetSampleResponse, customErrors.ErrInterface)
	}
)

func NewGetSampleUsecase(ur repository.SampleRepositoryInterface) GetSampleUsecaseInterface {
	return &getSampleUsecase{ur: ur}
}

func (gpmu *getSampleUsecase) Execute() (*result.GetSampleResponse, customErrors.ErrInterface) {
	profile, err := gpmu.ur.FindOneBySub()
	if errors.Is(err, customErrors.NotFoundErr) {
		return &result.GetSampleResponse{}, nil
	}
	if err != nil {
		return nil, customErrors.NewErr(err, 500, customErrors.InternalServerErrCode, "internal server error")
	}
	return &result.GetSampleResponse{Name: profile.Name}, nil
}
