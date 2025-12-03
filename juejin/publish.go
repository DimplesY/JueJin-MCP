package juejin

import (
	"context"
	"time"

	"github.com/go-rod/rod"
)

var (
	PUBLISH_URL = "https://juejin.cn/editor/drafts/new?v=2"
)

func Publish(page *rod.Page, ctx context.Context) error {
	p := page.Context(ctx)
	p.MustNavigate(PUBLISH_URL).MustWaitLoad()

	time.Sleep(10 * time.Second)
	return nil
}
