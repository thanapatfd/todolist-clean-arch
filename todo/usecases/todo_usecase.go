package usecases

import (
	"context"

	"github.com/thanapatfd/todolist/todo/entity"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("usecases")

type todoRepository interface {
	CreateList(ctx context.Context, list entity.List) (entity.List, error)
	GetListByID(ctx context.Context, id string) (entity.List, error)
	GetLists(ctx context.Context, name string, status string) ([]entity.List, error)
	UpdateList(ctx context.Context, list entity.List, id string) (entity.List, error)
	PatchList(ctx context.Context, list entity.List, id string) (entity.List, error)
	DeleteList(ctx context.Context, id string) error
	SortListsByID(ctx context.Context) ([]entity.List, error)
}
