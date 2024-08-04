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
		FindOneBySub(sub string) (*entity.Sample, error)
	}
)

func NewSampleRepository(sqlHandler *db.SQLHandler) SampleRepositoryInterface {

	return &sampleRepository{sqlHandler: sqlHandler}
}

func (ur *sampleRepository) FindOneBySub(sub string) (*entity.Sample, error) {
	sampleEntity := new(model.Sample)
	err := ur.sqlHandler.Conn.Model(&model.Sample{}).Select([]string{"id", "name"}).Where("sub = ?", sub).First(sampleEntity).Error
	return ur.convertEntity(sampleEntity), err
}

func (ur *sampleRepository) convertEntity(sample *model.Sample) *entity.Sample {
	if sample == nil {
		return nil
	}
	return &entity.Sample{ID: sample.ID, Name: sample.Name}
}
