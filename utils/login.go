package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var (
	c_user   = "61567492240323"
	datr     = "ezY8Z8_WTr2Dtoo3L81_1KU0"
	fr       = ""
	presence = "C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1733476138677%2C%22v%22%3A1%7D"
	sb       = "cH4OZ290dVLwqIBKm5vyjIw9"
	wd       = "1366x681"
	xs       = "29%3AYx9WHvHL_Q2ZaA%3A2%3A1733475937%3A-1%3A13560"
)

func Login() (*rod.Browser, *rod.Page) {
	dir := "~/.config/google-chrome"

	u := launcher.New().UserDataDir(dir).Leakless(true).NoSandbox(true).Headless(false).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect().NoDefaultDevice()

	page := browser.MustPage("https://web.facebook.com/").MustWaitLoad().MustWindowMaximize()

	pageHasLoginButton := page.MustHas(`button[name="login"]`)

	switch {
	case pageHasLoginButton:
		// Define the session cookies
		cookies := []*proto.NetworkCookieParam{
			{
				Name:     "c_user",
				Value:    c_user,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "datr",
				Value:    datr,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "fr",
				Value:    fr,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "presence",
				Value:    presence,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "sb",
				Value:    sb,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "wd",
				Value:    wd,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
			{
				Name:     "xs",
				Value:    xs,
				Domain:   ".facebook.com",
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "None",
				Priority: "Medium",
			},
		}

		// Inject the session cookie
		err := browser.SetCookies(cookies)
		if err != nil {
			fmt.Println("Failed to set session cookie:", err)
			return nil, nil
		}

		// check if cookies are valid
		page = page.MustNavigate("https://web.facebook.com/").MustWaitLoad()

		pageHasLoginForm, _, err := page.Has(`form[data-testid="royal_login_form"]`)
		if err != nil {
			log.Println("Error checking if page has login form:", err)
			return nil, nil
		}
		switch {
		case pageHasLoginForm:
			fmt.Println("Invalid or expired cookies 😞")
			os.Exit(1)
		default:
			fmt.Println("Log in complete 😊")
			page.MustScreenshot("home.png")
			return browser, page
		}
	default:
		fmt.Println("User is logged in 😊")
		page.MustScreenshot("home.png")
	}

	return browser, page
}
