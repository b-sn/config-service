package models

type App struct {
	ID       uint
	Name     string
	User     *User
	Token    string
	IsActive bool
}

type AppRepository interface {
	Create(app *App) error
	Get(app *App) error
	// FindByName(name string) (*App, error)
	// FindByToken(token string) (*App, error)
	// Update(params map[string]string) error
	// Deactivate() error
}
