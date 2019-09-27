package demoservice

import (
	"github.com/server-may-cry/go-workers-test/internal/db"
	"github.com/server-may-cry/go-workers-test/internal/entity"
)

type Service struct {
	myEntityRepository *db.MyEntryRepository
}

func New (r *db.MyEntryRepository) *Service {
	return &Service{
		myEntityRepository: r,
	}
}

func (s *Service) InsertObject(entry *entity.MyEntry) error {
	return s.myEntityRepository.Insert(entry)
}

func (s *Service) ReadObject() (entity.MyEntry, bool) {
	return s.myEntityRepository.SearchLastInserted()
}
