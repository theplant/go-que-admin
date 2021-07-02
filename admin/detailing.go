package admin

import (
	"github.com/goplaid/web"
	"github.com/goplaid/x/presets"
	. "github.com/goplaid/x/vuetify"
	"github.com/theplant/go-que-admin/models"
	h "github.com/theplant/htmlgo"
)

func detailing(m *presets.ModelBuilder) {
	d := m.Detailing("Actions", "ID", "Queue")

	d.Field("Actions").ComponentFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) h.HTMLComponent {
		return VBtn("Back").Href("/admin/goque-jobs")
	})

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

}
