package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//


type (
	// Interface combines interfaces of all supported store interfaces
	Interface interface {
	{{ range . -}}
		{{ pubIdent .Types.Plural }}
	{{ end }}
	}

	// Interface combines interfaces of all supported store interfaces
	Provisioned interface {
	{{ range . -}}
		{{ pubIdent .Types.Plural }}Provisioned
	{{ end }}
	}
)
