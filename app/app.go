package app

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/app"
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	httpApiServer interface {
		MountRoutes(mm ...func(chi.Router))
		Serve(ctx context.Context)
	}

	grpcServer interface {
		RegisterServices(func(server *grpc.Server))
		Serve(ctx context.Context)
	}

	CortezaApp struct {
		Opt *app.Options
		lvl int
		Log *zap.Logger

		// CLI Commands
		Command *cobra.Command

		// Servers
		HttpServer httpApiServer
		GrpcServer grpcServer
	}
)

func New() *CortezaApp {
	app := &CortezaApp{
		Opt: app.NewOptions(),
		lvl: bootLevelWaiting,
	}

	app.InitCLI()

	return app
}
