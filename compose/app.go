package compose

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/commands"
	"github.com/cortezaproject/corteza-server/compose/rest"
	"github.com/cortezaproject/corteza-server/compose/service"
	"github.com/cortezaproject/corteza-server/compose/service/event"
	"github.com/cortezaproject/corteza-server/corteza/store/mysql"
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
	"github.com/cortezaproject/corteza-server/pkg/app"
	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/automation"
	"github.com/cortezaproject/corteza-server/pkg/corredor"
	"github.com/cortezaproject/corteza-server/pkg/scheduler"
	"github.com/cortezaproject/corteza-server/system/auth/external"
	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type (
	App struct {
		Opts *app.Options
		Log  *zap.Logger
	}
)

const SERVICE = "compose"

func (app *App) Setup(log *zap.Logger, opts *app.Options) (err error) {
	app.Log = log.Named(SERVICE)
	app.Opts = opts

	scheduler.Service().OnTick(
		event.ComposeOnInterval(),
		event.ComposeOnTimestamp(),
	)

	// @todo Wire in cross-service JWT maker for Corredor
	corredor.Service().SetUserFinder(nil)
	corredor.Service().SetRoleFinder(nil)

	return
}

func (app *App) Upgrade(ctx context.Context) (err error) {
	s, err := mysql.New(ctx, app.Opts.DB.DSN)
	if err != nil {
		return
	}

	return provisioner.
		NewProvisioner(func(_ int, msg string) { app.Log.Info(msg) }).
		Run(s.ProvisionCompose())
}

// Initialized
func (app *App) Initialize(ctx context.Context) (err error) {
	// Connects to all services it needs to
	err = service.Initialize(ctx, app.Log, service.Config{
		ActionLog: app.Opts.ActionLog,
		Storage:   app.Opts.Storage,
	})

	if err != nil {
		return
	}

	// Initialize external (from default settings)
	external.Init()
	return
}

func (app *App) Activate(ctx context.Context) (err error) {
	if err = service.Activate(ctx); err != nil {
		return
	}

	service.Watchers(ctx)

	return
}

func (app *App) Provision(ctx context.Context) (err error) {
	ctx = auth.SetSuperUserContext(ctx)

	if err = provisionConfig(ctx, app.Log); err != nil {
		return
	}

	return
}

func (app *App) MountApiRoutes(r chi.Router) {
	rest.MountRoutes(r)
}

func (app *App) RegisterCliCommands(p *cobra.Command) {
	p.AddCommand(
		commands.Importer(),
		commands.Exporter(),
		commands.NGImporter(),
		// temp command, will be removed in 2020.6
		automation.ScriptExporter(SERVICE),
	)
}
