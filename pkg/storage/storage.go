package storage

import (
	"github.com/fesyunoff/phone-book/pkg/service/dto"
)

type Book interface {
	CreateEntry(t dto.Entry) (msg string, err error)
	DisplayEntry(t dto.Entry) (out []*dto.Entry, err error)
	UpdateEntry(t dto.Entry) (msg string, err error)
	DeleteEntry(t dto.Entry) (msg string, err error)
}
