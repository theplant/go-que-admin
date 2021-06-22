package admin

import (
	"github.com/goplaid/x/presets"
	"github.com/theplant/go-que-admin/models"
)

func configTest(b *presets.Builder) {
	b.Model(&models.TestUser{})
}
