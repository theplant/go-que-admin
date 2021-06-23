package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/theplant/go-que-admin/config"
	"github.com/theplant/go-que-admin/models"
	h "github.com/theplant/htmlgo"
	"github.com/tnclong/go-que"
)

func configQue(b *presets.Builder) {
	m := b.Model(&models.GoqueJob{})
	l := m.Listing("ID", "Queue", "Args", "RunAt", "DoneAt", "RetryCount", "LastErrMsg", "UniqueID", "UniqueLifeCycle")

	l.Field("RunAt").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		runAt := obj.(*models.GoqueJob).RunAt
		return h.Td(h.Text(fmt.Sprint(runAt)))
	})

	eb := m.Editing("Queue", "Args", "RetryPolicy")

	eb.Field("Queue").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		j := obj.(*models.GoqueJob)
		return h.Div().Text(j.Queue)
	})

	eb.SaveFunc(func(job interface{}, id string, ctx *web.EventContext) (err error) {
		q := config.TheQ

		j := job.(*models.GoqueJob)
		var args []interface{}
		json.Unmarshal([]byte(j.Args), &args)
		var retryPolicy que.RetryPolicy
		json.Unmarshal([]byte(j.RetryPolicy), &retryPolicy)
		_, err = q.Enqueue(context.Background(), nil, que.Plan{
			Queue:       "import_pdf",
			Args:        que.Args(args...),
			RunAt:       time.Now(),
			RetryPolicy: retryPolicy,
		})
		return
	})
}
