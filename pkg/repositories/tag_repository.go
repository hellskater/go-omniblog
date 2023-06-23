package repositories

import (
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) Create(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) FindByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) FindAll() ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) Update(tag *models.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) Delete(tag *models.Tag) error {
	return r.db.Delete(tag).Error
}

func (r *TagRepository) FindByName(name string) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("name = ?", name).First(&tag).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}
