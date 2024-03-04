package unitOfWorkAbstraction

import (
	"context"

	"github.com/omarkamel02/SharedPackages/external/dbAbstraction"
)

// IUnitOfWork is an interface that defines the methods that a UnitOfWork should implement
type IUnitOfWork interface {
	GetDbTransaction() dbAbstraction.IDbTransaction
	BeginTransaction(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
