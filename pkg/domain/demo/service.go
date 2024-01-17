package demo

import "github.com/MR5356/go-template/pkg/middleware/database"

type Service struct {
	db *database.BaseMapper[*Demo]
}

func NewService() *Service {
	return &Service{
		db: database.NewMapper(database.GetDB(), &Demo{}),
	}
}

func (s *Service) Initialize() error {
	err := s.db.DB.AutoMigrate(&Demo{})
	if err != nil {
		return err
	}

	_ = s.db.Insert(&Demo{ID: 1, Title: "demo"})
	return nil
}
