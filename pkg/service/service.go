package service

import (
	"context"

	"github.com/fesyunoff/phone-book/pkg/service/dto"
)

type Service interface {
	Get(ctx context.Context, task dto.Entry) (out []*dto.Entry, msg string, err error)
	Add(ctx context.Context, task dto.Entry) (msg string, err error)
	Update(ctx context.Context, task dto.Entry) (msg string, err error)
	Delete(ctx context.Context, task dto.Entry) (msg string, err error)
}
