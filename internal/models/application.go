package models

type Application struct {
	Name     string
	User     *User
	Token    string
	IsActive bool
}

type AppRepository interface {
	Create(name string, user *User) (*Application, error)
	FindByName(name string) (*Application, error)
	FindByToken(token string) (*Application, error)
	Update(params map[string]string) error
	Deactivate() error
}
