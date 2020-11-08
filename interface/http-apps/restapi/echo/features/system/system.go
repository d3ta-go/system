package system

import (
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	captcha "github.com/d3ta-go/system/interface/http-apps/restapi/echo/features/system/captcha"
	response "github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewSystem new FSystem
func NewSystem(h *handler.Handler) (*FSystem, error) {

	f := new(FSystem)
	f.SetHandler(h)

	return f, nil
}

// FSystem represent FSystem
type FSystem struct {
	features.BaseFeature
}

// HealthCheck display system health check
func (f *FSystem) HealthCheck(c echo.Context) error {
	data := map[string]interface{}{"serviceStatus": "OK"}
	return response.OKWithData(data, c)
}

// GenerateCaptcha generate Captcha
func (f *FSystem) GenerateCaptcha(c echo.Context) error {
	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}

	resp := captcha.GenerateCaptchaID(cfg, c)

	return response.OKWithData(resp, c)
}

// GenerateCaptchaImage generate CaptchaImage
func (f *FSystem) GenerateCaptchaImage(c echo.Context) error {
	cfg, err := f.GetHandler().GetDefaultConfig()
	if err != nil {
		return err
	}

	return captcha.CaptchaServeHTTP(cfg, c)
}
