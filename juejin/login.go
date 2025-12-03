package juejin

import (
	"context"
	"time"

	"github.com/go-rod/rod"
)

func Login(page *rod.Page, ctx context.Context) error {
	p := page.Context(ctx)
	p.MustNavigate("https://juejin.cn/").MustWaitLoad()

	time.Sleep(1 * time.Second)

	page.MustElementX(`//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[2]/div/button`).MustClick()
	page.MustWaitStable().MustScreenshot("login.png")
	return nil
}
