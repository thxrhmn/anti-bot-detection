package main

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func main() {
	standard()
	stealthMode()
}

func standard() {
	launch := launcher.New().Headless(false)
	url := launch.MustLaunch()

	browser := rod.New().ControlURL(url).Timeout(time.Minute).MustConnect()
	defer browser.Close()

	page := browser.MustPage("https://bot.sannysoft.com/")
	page.MustWaitDOMStable().MustScreenshot("images/screenshot-standard-1.png")

	page.MustNavigate("http://www.browserscan.net/bot-detection")
	page.MustWaitDOMStable().MustScreenshot("images/screenshot-standard-2.png")
}

func stealthMode() {
	launch := launcher.New().Headless(false)
	url := launch.MustLaunch()

	browser := rod.New().ControlURL(url).Timeout(time.Minute).MustConnect()
	defer browser.Close()

	page := stealth.MustPage(browser)

	page.MustNavigate("https://bot.sannysoft.com/")
	page.MustWaitDOMStable().MustScreenshot("images/screenshot-stealth-1.png")

	page.MustNavigate("https://www.browserscan.net/bot-detection")
	page.MustWaitDOMStable().MustScreenshot("images/screenshot-stealth-2.png")
}
