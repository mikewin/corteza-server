package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// {{ .Source }}

import (
	"context"
	"github.com/cortezaproject/corteza-server/store/provisioner"
{{- range .Import }}
	{{ normalizeImport . }}
{{- end }}
)

type (
	{{ $Types := .Types }}
	{{ $Fields := .Fields }}

	{{ pubIdent .Types.Plural }} interface {

	{{ if not .Search.Disable }}
		Search{{ pubIdent $Types.Plural }}(ctx context.Context, f {{ $Types.GoFilterType }}) ({{ $Types.GoSetType }}, {{ $Types.GoFilterType }}, error)
	{{- end }}
	{{- range .Lookups }}
		Lookup{{ pubIdent $Types.Singular }}By{{ pubIdent .Suffix }}(ctx context.Context{{- range $field := .Fields }}, {{ cc2underscore $field }} {{ ($field | $Fields.Find).Type  }}{{- end }}) (*{{ $Types.GoType }}, error)
	{{- end }}
		Create{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		Update{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		PartialUpdate{{ pubIdent $Types.Singular }}(ctx context.Context, onlyColumns []string, rr ... *{{ $Types.GoType }}) error
		Remove{{ pubIdent $Types.Singular }}(ctx context.Context, rr ... *{{ $Types.GoType }}) error
		Remove{{ pubIdent $Types.Singular }}By{{ template "primaryKeySuffix" $Fields }}(ctx context.Context {{ template "primaryKeyArgs" $Fields }}) error

		Truncate{{ pubIdent $Types.Plural }}(ctx context.Context) error
	}

	{{ pubIdent .Types.Plural }}Provisioned interface {
		{{ pubIdent .Types.Plural }}
		Provision{{ pubIdent $Types.Plural }}() provisioner.Executor
	}
)
