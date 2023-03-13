package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func Generate() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 104,
			B: 125,
			A: 100,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return id, b64s, err

}

func VerifyCaptcha(id string, value string) bool {
	if store.Verify(id, value, true) {
		return true
	} else {
		return false
	}
}
