package repository

import (
	"go-crud/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository.
func (*TagsRepositoryImpl) Delete(tagsId int) {
	panic("unimplemented")
}

// FindAll implements TagsRepository.
func (*TagsRepositoryImpl) FindAll() []model.Tags {
	panic("unimplemented")
}

// FindById implements TagsRepository.
func (*TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	panic("unimplemented")
}

// Save implements TagsRepository.
func (*TagsRepositoryImpl) Save(tags model.Tags) {
	panic("unimplemented")
}

// Update implements TagsRepository.
func (*TagsRepositoryImpl) Update(tags model.Tags) {
	panic("unimplemented")
}

