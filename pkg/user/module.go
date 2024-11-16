package user

import "gorm.io/gorm"

type Module struct {
	db         *gorm.DB
	repository Repository
	service    Service
	controller *Controller
}

func (m *Module) setupModule() {
	// 각 레이어 초기화하면서 의존성 주입
	m.repository = NewRepository(m.db)
	m.service = NewService(m.repository)
	m.controller = NewController(m.service)
}

func NewModule(db *gorm.DB) *Module {
	m := &Module{
		db: db,
	}
	m.setupModule()
	return m
}

// Controller 반환 - 라우터에서 사용
func (m *Module) Controller() *Controller {
	return m.controller
}

// Repository 반환 - 다른 모듈에서 필요할 경우 사용
func (m *Module) Repository() Repository {
	return m.repository
}
