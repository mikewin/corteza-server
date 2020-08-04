package test_store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//

import (
	"context"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/store/provisioner"
	"github.com/stretchr/testify/require"
	"testing"
)

func testAllGenerated(t *testing.T, all interface{}) {
{{ range . }}
	// Run generated tests for {{ .Types.Base }}
	t.Run({{ printf "%q" .Types.Base }}, func(t *testing.T) {
		var s = all.(store.{{ pubIdent .Types.Plural }}Provisioned)

		t.Run("provision", func(t *testing.T) {
			var (
				ctx = context.Background()
				req = require.New(t)
			)

			req.NoError(provisioner.NewProvisioner(nil).Run(s.Provision{{ pubIdent .Types.Plural }}()))
			req.NoError(s.Truncate{{ pubIdent .Types.Plural }}(ctx))
		})


		test{{ pubIdent .Types.Base }}(t, s)
	})
{{ end }}
}
