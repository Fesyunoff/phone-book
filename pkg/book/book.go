package book

import (
	"context"

	"github.com/fesyunoff/phone-book/pkg/service"
	"github.com/fesyunoff/phone-book/pkg/service/dto"

	db "github.com/fesyunoff/phone-book/pkg/storage/db"
)

type Service struct {
	db *db.PostgreBookStorage
}

var _ service.Service = (*Service)(nil)

func (s *Service) Add(ctx context.Context, p dto.Entry) (msg string, err error) {

	msg, err = s.db.CreateEntry(p)

	return
}

func (s *Service) Get(ctx context.Context, p dto.Entry) (out []*dto.Entry, msg string, err error) {

	out, err = s.db.DisplayEntry(p)

	return
}

func (s *Service) Update(ctx context.Context, p dto.Entry) (msg string, err error) {

	msg, err = s.db.UpdateEntry(p)

	return
}

func (s *Service) Delete(ctx context.Context, p dto.Entry) (msg string, err error) {

	msg, err = s.db.DeleteEntry(p)

	return
}

func NewService(db *db.PostgreBookStorage) *Service {
	return &Service{
		db: db,
	}
}
