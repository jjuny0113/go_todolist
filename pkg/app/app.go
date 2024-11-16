package app

import "gorm.io/gorm"

type Application struct {
	db *gorm.DB
}

func NewApplication(db *gorm.DB) *Application {
	app := &Application{
		db: db,
	}
	app.setModules()
	return app
}

func (a *Application) setModules() {

}
