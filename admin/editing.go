package admin

import (
	"context"
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

func editing(m *presets.ModelBuilder) {
	queues := config.MustGetQueues()

	eb := m.Editing("Queue", "Args", "RetryPolicy")

	eb.Field("Queue").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		ctx.Hub.RegisterEventFunc("updateArgsEditor", updateArgsEditor)

		j := obj.(*models.GoqueJob)
		// return VSelect().
		// 	Label("Queue").
		// 	Items(queues).
		// 	ItemText("Name").
		// 	ItemValue("Name").
		// 	Value(j.Queue).
		// 	FieldName("Queue")

		var options []h.HTMLComponent
		for _, q := range queues {
			opt := h.Option(q.Name).Value(q.Name)
			if q.Name == j.Queue {
				opt.Checked(true)
			}
			options = append(options, opt)
		}
		return web.Bind(h.Select(
			options...,
		)).OnInput("updateArgsEditor")
	})

	eb.Field("Args").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		queues := config.MustGetQueues()

		j := obj.(*models.GoqueJob)
		var queue string
		if j.Queue == "" {
			queue = queues[0].Name
		}
		return web.Portal(
			argsEditor(queue),
		).Name("argsEditor")
	})

	eb.Field("RetryPolicy").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		j := obj.(*models.GoqueJob)
		return retryPolicyEditor(j)
	})

	eb.SaveFunc(func(job interface{}, id string, ctx *web.EventContext) (err error) {
		q := config.TheQ

		j := job.(*models.GoqueJob)
		_, err = q.Enqueue(context.Background(), nil, que.Plan{
			Queue:       j.Queue,
			Args:        que.Args(j.Args...),
			RunAt:       time.Now(),
			RetryPolicy: j.RetryPolicy.RetryPolicy,
		})
		return
	})
}

func retryPolicyEditor(j *models.GoqueJob) h.HTMLComponent {
	return VContainer(
		VMenu(
			web.Slot(
				h.A().Text("Retry Policy").Attr("v-on:click", "vars.retryPolicyEditorShow = true"),
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
						VListItemTitle(h.Text("MaxInterval")),
						VListItemAction(
							VTextField().
								FieldName("RetryPolicy.MaxInterval").
								Value(fmt.Sprint(int64(j.RetryPolicy.MaxInterval))),
						),
					),
					VListItem(
						VListItemTitle(h.Text("NextIntervalMultiplier")),
						VListItemAction(
							VTextField().
								FieldName("RetryPolicy.NextIntervalMultiplier").
								Value(fmt.Sprint(j.RetryPolicy.NextIntervalMultiplier)),
						),
					),
					VListItem(
						VListItemTitle(h.Text("IntervalRandomPercent")),
						VListItemAction(
							VTextField().
								FieldName("RetryPolicy.IntervalRandomPercent").
								Value(fmt.Sprint(j.RetryPolicy.IntervalRandomPercent)),
						),
					),
					VListItem(
						VListItemTitle(h.Text("MaxRetryCount")),
						VListItemAction(
							VTextField().
								FieldName("RetryPolicy.MaxRetryCount").
								Value(fmt.Sprint(j.RetryPolicy.MaxRetryCount)),
						),
					),
				),

				VCardActions(
					VSpacer(),
					VBtn("Save").Text(true).Color("primary").
						On("click", "vars.retryPolicyEditorShow = false"),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetY(true).
			Attr("v-model", "vars.retryPolicyEditorShow"),
	).Attr(web.InitContextVars, `{retryPolicyEditorShow: false}`)
}

func updateArgsEditor(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.UpdatePortals = append(er.UpdatePortals, &web.PortalUpdate{
		Name: "argsEditor",
		Body: argsEditor(ctx.Event.Value),
	})
	return
}

func argsEditor(queue string) h.HTMLComponent {
	queues := config.MustGetQueues()
	var argsCfg []*config.QueueArg
	for _, q := range queues {
		if q.Name == queue {
			argsCfg = q.Args
			break
		}
	}

	var argItems []h.HTMLComponent
	for i, a := range argsCfg {
		argItems = append(
			argItems,
			VListItem(
				VListItemTitle(h.Text(a.Name)),
				VListItemAction(
					VTextField().
						FieldName(fmt.Sprintf("Args[%d]", i)),
				),
			),
		)
	}
	return VContainer(
		VMenu(
			web.Slot(
				h.A().Text("Args").Attr("v-on:click", "vars.argsEditorShow = true"),
			).Name("activator"),

			VCard(
				VDivider(),
				VList(
					argItems...,
				),

				VCardActions(
					VSpacer(),
					VBtn("Save").Text(true).Color("primary").
						On("click", "vars.argsEditorShow = false"),
				),
			),
		).CloseOnContentClick(false).
			NudgeWidth(200).
			OffsetY(true).
			Attr("v-model", "vars.argsEditorShow"),
	).Attr(web.InitContextVars, `{argsEditorShow: false}`)
}
