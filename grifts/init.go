package grifts

import (
	"github.com/francocorreasosa/francocorreacom/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
