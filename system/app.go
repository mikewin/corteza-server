package system

import (
	"context"
	"github.com/cortezaproject/corteza-server/corteza/store/mysql"
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
	"github.com/cortezaproject/corteza-server/pkg/app"
	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/automation"
	"github.com/cortezaproject/corteza-server/pkg/corredor"
	"github.com/cortezaproject/corteza-server/pkg/scheduler"
	"github.com/cortezaproject/corteza-server/system/auth/external"
	"github.com/cortezaproject/corteza-server/system/commands"
	systemGRPC "github.com/cortezaproject/corteza-server/system/grpc"
	"github.com/cortezaproject/corteza-server/system/proto"
	"github.com/cortezaproject/corteza-server/system/rest"
	"github.com/cortezaproject/corteza-server/system/service"
	"github.com/cortezaproject/corteza-server/system/service/event"
	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	App struct {
		Opts *app.Options
		Log  *zap.Logger
	}
)

const SERVICE = "system"

func (app *App) Setup(log *zap.Logger, opts *app.Options) (err error) {
	app.Log = log.Named(SERVICE)
	app.Opts = opts

	scheduler.Service().OnTick(
		event.SystemOnInterval(),
		event.SystemOnTimestamp(),
	)

	corredor.Service().SetUserFinder(service.DefaultUser)
	corredor.Service().SetRoleFinder(service.DefaultRole)

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

	// Initialize external authentication (from default settings)
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

	// creates default applications that will appear in Unify/One
	// @todo migrate this to provisioning/YAML
	if err = makeDefaultApplications(ctx, app.Log); err != nil {
		return
	}

	// auto-discovery auth.* settings
	if err = authSettingsAutoDiscovery(ctx, app.Log, service.DefaultSettings); err != nil {
		return
	}

	// external provider auto configuration
	// creates: auth.external.providers.(google|linkedin|github|facebook).*
	if err = authAddExternals(ctx); err != nil {
		return
	}

	// OIDC provider auto configuration
	// creates: auth.external.providers.openid-connect.*
	if err = oidcAutoDiscovery(ctx, app.Log); err != nil {
		return
	}

	return
}

func (app *App) MountApiRoutes(r chi.Router) {
	rest.MountRoutes(r)
}

func (app *App) RegisterGrpcServices(server *grpc.Server) {
	proto.RegisterUsersServer(server, systemGRPC.NewUserService(
		service.DefaultUser,
		service.DefaultAuth,
		auth.DefaultJwtHandler,
		service.DefaultAccessControl,
	))

	proto.RegisterRolesServer(server, systemGRPC.NewRoleService(
		service.DefaultRole,
	))
}

func (app *App) RegisterCliCommands(p *cobra.Command) {
	p.AddCommand(
		commands.Importer(),
		commands.Exporter(),
		commands.Settings(),
		commands.Auth(),
		commands.Users(),
		commands.Roles(),
		commands.Sink(),
		commands.RBAC(),
		// temp command, will be removed in 2020.6
		automation.ScriptExporter(SERVICE),
	)
}
