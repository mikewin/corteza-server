package app

import (
	"context"
	"fmt"
	cmpService "github.com/cortezaproject/corteza-server/compose/service"
	cmpEvent "github.com/cortezaproject/corteza-server/compose/service/event"
	msgService "github.com/cortezaproject/corteza-server/messaging/service"
	msgEvent "github.com/cortezaproject/corteza-server/messaging/service/event"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/corredor"
	"github.com/cortezaproject/corteza-server/pkg/db"
	"github.com/cortezaproject/corteza-server/pkg/eventbus"
	"github.com/cortezaproject/corteza-server/pkg/healthcheck"
	"github.com/cortezaproject/corteza-server/pkg/http"
	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/pkg/mail"
	"github.com/cortezaproject/corteza-server/pkg/monitor"
	"github.com/cortezaproject/corteza-server/pkg/scenario"
	"github.com/cortezaproject/corteza-server/pkg/scheduler"
	"github.com/cortezaproject/corteza-server/pkg/sentry"
	"github.com/cortezaproject/corteza-server/store/mysql"
	"github.com/cortezaproject/corteza-server/system/auth/external"
	sysService "github.com/cortezaproject/corteza-server/system/service"
	sysEvent "github.com/cortezaproject/corteza-server/system/service/event"
	"strings"
	"time"
)

const (
	bootLevelWaiting = iota
	bootLevelSetup
	bootLevelStoreInitialized
	bootLevelServicesInitialized
	bootLevelUpgraded
	bootLevelProvisioned
	bootLevelActivated
)

// Setup configures all required services
func (app *CortezaApp) Setup() (err error) {
	logger.Init()
	app.Log = logger.Default()

	if app.lvl >= bootLevelSetup {
		// Are basics already set-up?
		return nil
	}

	hcd := healthcheck.Defaults()
	hcd.Add(scheduler.Healthcheck, "Scheduler")
	hcd.Add(mail.Healthcheck, "Mail")
	hcd.Add(corredor.Healthcheck, "Corredor")
	hcd.Add(db.Healthcheck(), "Database")

	if err = sentry.Init(app.Opt.Sentry); err != nil {
		return fmt.Errorf("could not initialize Sentry: %w", err)
	}

	// Use Sentry right away to handle any panics
	// that might occur inside auth, mail setup...
	defer sentry.Recover()

	auth.SetupDefault(app.Opt.Auth.Secret, int(app.Opt.Auth.Expiry/time.Minute))
	mail.SetupDialer(app.Opt.SMTP.Host, app.Opt.SMTP.Port, app.Opt.SMTP.User, app.Opt.SMTP.Pass, app.Opt.SMTP.From)

	http.SetupDefaults(
		app.Opt.HTTPClient.HttpClientTimeout,
		app.Opt.HTTPClient.ClientTSLInsecure,
	)

	monitor.Setup(app.Log, app.Opt.Monitor)

	scheduler.Setup(app.Log, eventbus.Service(), 0)
	scheduler.Service().OnTick(
		sysEvent.SystemOnInterval(),
		sysEvent.SystemOnTimestamp(),
		cmpEvent.ComposeOnInterval(),
		cmpEvent.ComposeOnTimestamp(),
		msgEvent.MessagingOnInterval(),
		msgEvent.MessagingOnTimestamp(),
	)

	if err = corredor.Setup(app.Log, app.Opt.Corredor); err != nil {
		return err
	}

	app.lvl = bootLevelSetup
	return
}

// InitStore initializes store backend(s) and runs upgrade procedures
func (app *CortezaApp) InitStore(ctx context.Context) error {
	if app.lvl >= bootLevelStoreInitialized {
		// Is store already initialised?
		return nil
	} else if err := app.Setup(); err != nil {
		// Initialize previous level
		return err
	}

	defer sentry.Recover()

	// @todo this should be configurable
	s, err := mysql.New(ctx, app.Opt.DB.DSN)
	if err != nil {
		return err
	}

	// @todo store provision => store upgrade
	ctx = actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_APP_Upgrade)

	err = scenario.
		NewScenario(func(_ int, msg string) { app.Log.Debug(strings.TrimSpace(msg)) }).
		Play(s.Provision())

	if err != nil {
		return err
	}

	// deprecated connector
	// current state of Corteza (repos...) still requires it
	_, err = db.TryToConnect(ctx, app.Log, app.Opt.DB)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}

	app.lvl = bootLevelStoreInitialized
	return nil
}

// InitServices initializes all services used
func (app *CortezaApp) InitServices(ctx context.Context) (err error) {
	if app.lvl >= bootLevelServicesInitialized {
		return nil
	} else if err := app.InitStore(ctx); err != nil {
		return err

	}

	ctx = actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_APP_Init)
	defer sentry.Recover()

	if err = corredor.Service().Connect(ctx); err != nil {
		return
	}

	corredor.Service().SetUserFinder(sysService.DefaultUser)
	corredor.Service().SetRoleFinder(sysService.DefaultRole)

	// Initializes system services
	//
	// Note: this is a legacy approach, all services from all 3 apps
	// will most likely be merged in the future
	err = sysService.Initialize(ctx, app.Log, sysService.Config{
		ActionLog: app.Opt.ActionLog,
		Storage:   app.Opt.Storage,
	})

	if err != nil {
		return
	}

	// Initializes compose services
	//
	// Note: this is a legacy approach, all services from all 3 apps
	// will most likely be merged in the future
	err = cmpService.Initialize(ctx, app.Log, cmpService.Config{
		ActionLog: app.Opt.ActionLog,
		Storage:   app.Opt.Storage,
	})

	if err != nil {
		return
	}

	// Initializes messaging services
	//
	// Note: this is a legacy approach, all services from all 3 apps
	// will most likely be merged in the future
	err = msgService.Initialize(ctx, app.Log, msgService.Config{
		ActionLog: app.Opt.ActionLog,
		Storage:   app.Opt.Storage,
	})

	if err != nil {
		return
	}

	// Initialize external authentication (from default settings)
	external.Init()

	app.lvl = bootLevelServicesInitialized
	return
}

// Provision instance with configuration and settings
// by importing preset configurations and running autodiscovery procedures
func (app *CortezaApp) Provision(ctx context.Context) (err error) {
	if app.lvl >= bootLevelProvisioned {
		return
	}

	if err := app.InitServices(ctx); err != nil {
		return err

	}

	ctx = actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_APP_Provision)
	//defer sentry.Recover()
	//
	//ctx = auth.SetSuperUserContext(ctx)
	//Log := logger.Default()
	//
	//// ****************************************************************************************************************
	//// system
	//if err = provisionConfig(ctx, Log); err != nil {
	//	return
	//}
	//
	//// creates default applications that will appear in Unify/One
	//// @todo migrate this to provisioning/YAML
	//if err = makeDefaultApplications(ctx, Log); err != nil {
	//	return
	//}
	//
	//// auto-discovery auth.* settings
	//if err = authSettingsAutoDiscovery(ctx, app.Log, service.DefaultSettings); err != nil {
	//	return
	//}
	//
	//// external provider auto configuration
	//// creates: auth.external.providers.(google|linkedin|github|facebook).*
	//if err = authAddExternals(ctx); err != nil {
	//	return
	//}
	//
	//// OIDC provider auto configuration
	//// creates: auth.external.providers.openid-connect.*
	//if err = oidcAutoDiscovery(ctx, Log); err != nil {
	//	return
	//}
	//
	//ctx = auth.SetSuperUserContext(ctx)
	//
	//// ****************************************************************************************************************
	//// compose
	//if err = provisionConfig(ctx, Log); err != nil {
	//	return
	//}
	//
	//// ****************************************************************************************************************
	//// messaging
	//if err = provisionConfig(ctx, Log); err != nil {
	//	return
	//}

	app.lvl = bootLevelProvisioned
	return
}

// Activate start all internal services and watchers
func (app *CortezaApp) Activate(ctx context.Context) (err error) {
	if app.lvl >= bootLevelActivated {
		return
	} else if err := app.Provision(ctx); err != nil {
		return err
	}

	ctx = actionlog.RequestOriginToContext(ctx, actionlog.RequestOrigin_APP_Activate)
	defer sentry.Recover()

	// Start scheduler
	scheduler.Service().Start(ctx)

	// Load corredor scripts & init watcher (script reloader)
	corredor.Service().Load(ctx)
	corredor.Service().Watch(ctx)

	sysService.Watchers(ctx)
	cmpService.Watchers(ctx)
	msgService.Watchers(ctx)

	app.lvl = bootLevelActivated
	return nil
}
