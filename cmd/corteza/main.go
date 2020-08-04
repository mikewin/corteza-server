package main

import (
	"github.com/cortezaproject/corteza-server/app"
	"github.com/cortezaproject/corteza-server/pkg/cli"
)

func main() {
	cli.HandleError(app.New().Execute())
}
