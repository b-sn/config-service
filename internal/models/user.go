package models

type User struct {
	ID       uint   `json:"-" gorm:"primaryKey"`
	Name     string `param:"user_name" json:"user_name" gorm:"uniqueIndex"`
	Token    string `json:"user_token"`
	IsActive bool   `json:"is_active"`
}

type UserRepository interface {
	Create(user *User) error
	Get(user *User) error
	Find(user *User) ([]*User, error)
	// List() ([]*User, error)
	// GetUserByName(user *User) error
	// Deactivate() error
}
