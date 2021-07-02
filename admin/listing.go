package admin

import (
	"fmt"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	"github.com/theplant/go-que-admin/models"
	h "github.com/theplant/htmlgo"
)

func listing(m *presets.ModelBuilder) {
	l := m.Listing("ID", "Queue", "Args", "RunAt", "DoneAt", "RetryCount", "LastErrMsg", "UniqueID", "UniqueLifeCycle")

	l.Field("RunAt").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		runAt := obj.(*models.GoqueJob).RunAt
		return h.Td(h.Text(fmt.Sprint(runAt))).Style("width: 100px")
	})
}
