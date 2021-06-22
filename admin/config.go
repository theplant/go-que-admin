package admin

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	"github.com/goplaid/x/vuetify"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	configTest(b)
	configQue(b)

	return
}
