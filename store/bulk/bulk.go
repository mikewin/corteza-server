package bulk

import (
	"context"
	"github.com/cortezaproject/corteza-server/store"
)

type (
	Job interface {
		Do(ctx context.Context, s store.Interface) error
	}
)
