package settings

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	log *logrus.Logger
}

type StatusResp struct {
	Module string `json:"module"`
	Time   string `json:"time"`
}

func NewService(db *gorm.DB, log *logrus.Logger) *Service {
	return &Service{db: db, log: log}
}

func (s *Service) Status() StatusResp {
	return StatusResp{
		Module: "settings",
		Time:   time.Now().Format(time.RFC3339),
	}
}
