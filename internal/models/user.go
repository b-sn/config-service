package models

type User struct {
	Name     string `param:"user_name" json:"user_name"`
	Token    string `json:"user_token"`
	IsActive bool   `json:"is_active"`
}

type UserRepository interface {
	Create(user *User) error
	List() ([]*User, error)
	// FindByName(name string) (*User, error)
	// Deactivate() error
}
