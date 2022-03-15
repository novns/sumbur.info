package blog

import (
	"sumbur/views"

	"github.com/savsgio/atreugo/v11"
)

type Blog struct {
	*views.BasePage
}

func BlogGET(ctx *atreugo.RequestCtx) error {
	views.WritePage(ctx, &Blog{})
	return nil
}
