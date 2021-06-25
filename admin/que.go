package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	. "github.com/goplaid/x/vuetify"
	"github.com/theplant/go-que-admin/config"
	"github.com/theplant/go-que-admin/models"
	h "github.com/theplant/htmlgo"
	"github.com/tnclong/go-que"
)

func retryPolicyEditor(j *models.GoqueJob) h.HTMLComponent {
	return VContainer(
		VMenu(
			web.Slot(
				h.A().Text("Retry Policy").Attr("v-on:click", "vars.myMenuShow = true"),
			).Name("activator"),

			VCard(
				VDivider(),
				VList(
					VListItem(
						VListItemTitle(h.Text("Initial Interval")),
						VListItemAction(
							VTextField().
								FieldName("RetryPolicy.InitialInterval").
								Value(fmt.Sprint(int64(j.RetryPolicy.InitialInterval))),
						),
					),
					VListItem(
						VListItemAction(
							VSwitch().Color("purple").
								FieldName("EnableHints").
								InputValue(true),
						),
						VListItemTitle(h.Text("Enable hints")),
					),
				),

				VCardActions(
					VSpacer(),
					VBtn("Cancel").Text(true).
						On("click", "vars.myMenuShow = false"),
					VBtn("Save").Color("primary").
						Text(true).OnClick("submit"),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetY(true).
			Attr("v-model", "vars.myMenuShow"),
	).Attr(web.InitContextVars, `{myMenuShow: false}`)
}

func configQue(b *presets.Builder) {
	m := b.Model(&models.GoqueJob{})

	d := m.Detailing("ID", "Queue", "Actions")

	d.Field("Queue").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {

		j := obj.(*models.GoqueJob)
		return VCard(
			VCardTitle(h.Text(j.Queue)),
			VCardText(
				VChip(h.Text("5:30")),
				VChip(h.Text("6:30")),
			),
		)

	})

	d.Field("Actions").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {

		//j := obj.(*models.GoqueJob)
		return VBtn("Back").Href("/admin/goque-jobs")
	})

	_ = d

	l := m.Listing("ID", "Queue", "Args", "RunAt", "DoneAt", "RetryCount", "LastErrMsg", "UniqueID", "UniqueLifeCycle")

	l.Field("RunAt").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		runAt := obj.(*models.GoqueJob).RunAt
		return h.Td(h.Text(fmt.Sprint(runAt))).Style("width: 100px")
	})

	eb := m.Editing("Queue", "Args", "RetryPolicy")

	eb.Field("Queue").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		j := obj.(*models.GoqueJob)
		return h.Div().Text(j.Queue)
	})

	eb.Field("RetryPolicy").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		j := obj.(*models.GoqueJob)
		return retryPolicyEditor(j)
	})

	eb.SaveFunc(func(job interface{}, id string, ctx *web.EventContext) (err error) {
		q := config.TheQ

		j := job.(*models.GoqueJob)
		fmt.Println("j.RetryPolicy.InitialInterval = ", j.RetryPolicy.InitialInterval)
		var args []interface{}
		json.Unmarshal([]byte(j.Args), &args)
		//json.Unmarshal([]byte(j.RetryPolicy), &retryPolicy)
		_, err = q.Enqueue(context.Background(), nil, que.Plan{
			Queue:       "import_pdf",
			Args:        que.Args(args...),
			RunAt:       time.Now(),
			RetryPolicy: j.RetryPolicy.RetryPolicy,
		})
		return
	})
}
