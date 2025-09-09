package repositories

import (
	"github.com/dimashky/go-appointment/app/database"
	"github.com/dimashky/go-appointment/app/models"
)

// UserRepository interface defines the contract for user operations
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	List(limit, offset int) ([]models.User, error)
}

// userRepository implements UserRepository interface
type userRepository struct {
	db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates an existing user
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete soft deletes a user by ID
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// List retrieves a paginated list of users
func (r *userRepository) List(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
