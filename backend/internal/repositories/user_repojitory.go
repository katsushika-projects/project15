package repositories

import (
	"errors"

	"github.com/moto340/project15/backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateAccessToken(user *models.User, refreshToken string) error {
	user.RefreshToken = refreshToken
	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateRefreshToken(user_id string) error {
	if err := r.db.Model(&models.User{}).Where("id = ?", user_id).Update("refresh_token", "").Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) PostBlackList(access_token string) error {
	blacklist := models.BlackList{
		AccessToken: access_token,
	}
	if err := r.db.Model(&models.BlackList{}).Create(&blacklist).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) AuthBlackList(token string) error {
	var blacklist models.BlackList
	if err := r.db.Model(&models.BlackList{}).Where("access_token = ?", token).First(&blacklist).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードが見つからない場合は正常ケース
			return nil
		}
		// その他のエラーはそのまま返す
		return err
	}

	// レコードが見つかった場合
	return errors.New("token already exists")
}
