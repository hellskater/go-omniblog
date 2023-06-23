package repositories

import (
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{db}
}

func (r *LikeRepository) Create(like *models.Like) error {
	return r.db.Create(like).Error
}

func (r *LikeRepository) FindByID(id uint) (*models.Like, error) {
	var like models.Like
	err := r.db.Preload("User").First(&like, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &like, nil
}

func (r *LikeRepository) FindByUserAndPost(user *models.User, post *models.Post) (*models.Like, error) {
	var like models.Like
	err := r.db.Preload("User").Where("user_id = ? AND post_id = ?", user.ID, post.ID).First(&like).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &like, nil
}

func (r *LikeRepository) Update(like *models.Like) error {
	return r.db.Save(like).Error
}

func (r *LikeRepository) Delete(like *models.Like) error {
	return r.db.Delete(like).Error
}
