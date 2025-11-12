package juejin

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (j *JueJin) Login() error {
	bin, _ := launcher.LookPath()
	u := launcher.New().Bin(bin).Set("headless").Delete("--headless").MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("https://juejin.cn/")
	page.MustElementX(`//*[@id="juejin"]/div[1]/div/header/div/nav/ul/ul/li[3]/div/button`).MustClick()
	page.MustWaitStable().MustScreenshot("login.png")
	return nil
}
