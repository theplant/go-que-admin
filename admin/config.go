package admin

import (
	"github.com/goplaid/web"
	"github.com/theplant/go-que-admin/models"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	"github.com/goplaid/x/vuetify"
	h "github.com/theplant/htmlgo"
)

func NewConfig() (b *presets.Builder) {
	db := ConnectDB()

	b = presets.New()
	b.URIPrefix("/admin").
		BrandTitle("go-que-admin").
		DataOperator(gormop.DataOperator(db)).
		HomePageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
			r.Body = vuetify.VContainer(
				h.H1("Home"),
				h.P().Text("Change your home page here"))
			return
		})
	m := b.Model(&models.Post{})
	_ = m
	// Use m to customize the model, Or config more models here.
	return
}
