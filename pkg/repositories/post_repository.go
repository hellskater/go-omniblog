package repositories

import (
	"github.com/hellskater/omniblog/pkg/models"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("Author").Preload("Tags").Preload("Categories").First(&post, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("Author").Preload("Tags").Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(post *models.Post) error {
	return r.db.Delete(post).Error
}

func (r *PostRepository) FindByAuthor(author *models.User) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("Author").Preload("Tags").Preload("Categories").Where("author_id = ?", author.ID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
