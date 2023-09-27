package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restEcho1/configs"
	"restEcho1/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBarangs(t *testing.T) {

	// Setup Controller

	var cfg = configs.InitConfig()
	var gorm = model.InitModel(*cfg)

	var mdl = model.BarangModel{}
	mdl.Init(gorm)
	var ctl = BarangController{}
	ctl.InitUserController(mdl)

	var e = echo.New()

	var req = httptest.NewRequest(http.MethodGet, "/", nil)
	var res = httptest.NewRecorder()

	var c = e.NewContext(req, res)
	c.SetPath("/barangs")

	var tmp = map[string]any{}

	var resData = res.Body.Bytes()
	err := json.Unmarshal(resData, &tmp)

	// var resData = json.NewDecoder(res.Body)
	// resData.Decode(&tmp)

	assert.Equal(t, 200, res.Code)
	assert.Error(t, err)

}
