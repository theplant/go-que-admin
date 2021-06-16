package admin

import (
	"context"
	"fmt"
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/goplaid/x/presets/gormop"
	"github.com/goplaid/x/vuetify"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/theplant/go-que-admin/config"
	"github.com/theplant/go-que-admin/models"
	h "github.com/theplant/htmlgo"
	"github.com/tnclong/go-que"
	"time"
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
	m := b.Model(&models.GoqueJob{})
	l := m.Listing("ID", "Args", "RunAt", "DoneAt", "RetryPolicy","RetryCount", "LastErrMsg", "LastErrStack", "UniqueID", "UniqueLifeCycle")

	l.Field("RunAt").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		runAt := obj.(*models.GoqueJob).RunAt
		return h.Td(h.Text(fmt.Sprint(runAt)))
	})


	m.Editing("Args").SaveFunc(func(job interface{}, id string, ctx *web.EventContext) (err error) {
		q := config.TheQ

		j := job.(*models.GoqueJob)
		_, err = q.Enqueue(context.Background(), nil, que.Plan{
			Queue: "import_pdf",
			Args: que.Args(j.Args),
			RunAt: time.Now(),
		})
		return
	})
	_ = m
	// Use m to customize the model, Or config more models here.
	return
}
