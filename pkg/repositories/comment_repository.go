package repositories

import (
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("Author").First(&comment, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) FindAllByPost(post *models.Post) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("Author").Where("post_id = ?", post.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) Delete(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}
