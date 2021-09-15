package gowebtemplate

import (
	"context"

	"github.com/menggggggg/go-web-template/internal/app"
)

func RunGoWebTemplateServer() {
	app.Run(context.Background())
}
