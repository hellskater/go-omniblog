package repositories

import (
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
)

type FollowerRepositoryInterface interface {
	Follow(follower *models.Follower) error
	Unfollow(followerID, followingID uint) error
	FindByFollowerID(followerID uint) ([]models.Follower, error)
	FindByFollowingID(followingID uint) ([]models.Follower, error)
}

type FollowerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) *FollowerRepository {
	return &FollowerRepository{db}
}

func (r *FollowerRepository) Follow(follower *models.Follower) error {
	return r.db.Create(follower).Error
}

func (r *FollowerRepository) Unfollow(followerID, followingID uint) error {
	return r.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&models.Follower{}).Error
}

func (r *FollowerRepository) FindByFollowerID(followerID uint) ([]models.Follower, error) {
	var followers []models.Follower
	err := r.db.Where("follower_id = ?", followerID).Find(&followers).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return followers, nil
}

func (r *FollowerRepository) FindByFollowingID(followingID uint) ([]models.Follower, error) {
	var followers []models.Follower
	err := r.db.Where("following_id = ?", followingID).Find(&followers).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return followers, nil
}
