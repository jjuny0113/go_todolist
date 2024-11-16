package app

import (
	"gorm.io/gorm"
	"todoList/pkg/user"
)

type Application struct {
	db         *gorm.DB
	UserModule *user.Module
}

func NewApplication(db *gorm.DB) *Application {
	app := &Application{
		db: db,
	}
	app.setModules()
	return app
}

func (a *Application) setModules() {
	a.UserModule = user.NewModule(a.db)
}
