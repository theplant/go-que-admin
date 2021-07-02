package admin

import (
	"github.com/goplaid/x/presets"
	"github.com/theplant/go-que-admin/models"
)

func configQue(b *presets.Builder) {
	m := b.Model(&models.GoqueJob{})

	listing(m)
	detailing(m)
	editing(m)
}
