package repository

import (
	"my-go-template/src/adapter/repository/model"
	"my-go-template/src/driver/db"
	"my-go-template/src/entity"
)

type (
	sampleRepository struct {
		sqlHandler *db.SQLHandler
	}

	SampleRepositoryInterface interface {
		FindOneBySub() (*entity.Sample, error)
	}
)

func NewSampleRepository(sqlHandler *db.SQLHandler) SampleRepositoryInterface {

	return &sampleRepository{sqlHandler: sqlHandler}
}

func (ur *sampleRepository) FindOneBySub() (*entity.Sample, error) {
	sampleModel := new(model.Sample)
	err := ur.sqlHandler.Conn.Select([]string{"id", "name"}).First(sampleModel).Error
	return ur.convertEntity(sampleModel), err
}

func (ur *sampleRepository) convertEntity(sample *model.Sample) *entity.Sample {
	if sample == nil {
		return nil
	}
	return &entity.Sample{ID: sample.ID, Name: sample.Name}
}
