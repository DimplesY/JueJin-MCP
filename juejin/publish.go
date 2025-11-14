package juejin

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (j *JueJin) PublishArticle() error {
	bin, _ := launcher.LookPath()
	u := launcher.New().Bin(bin).Set("headless").Delete("--headless").MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://juejin.cn/editor/drafts/new?v=2")
	page.MustWaitLoad()
	return nil
}
