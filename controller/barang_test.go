package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restEcho1/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockModel struct{}

func (mm *MockModel) Insert(newItem model.Barang) *model.Barang {
	return nil
}
func (mm *MockModel) GetAllBarang() []model.Barang {
	return nil
}
func (mm *MockModel) Delete(id int) {

}
func (mm *MockModel) UpdateData(updatedData model.Barang) bool {
	return false
}
func (mm *MockModel) UpdateData2(updatedData model.Barang) bool {
	return false
}

func TestGetBarangs(t *testing.T) {

	// Setup Controller

	var mdl = MockModel{}

	var ctl = NewBarangControllInterface(&mdl)

	var e = echo.New()
	e.GET("/barangs", ctl.GetBarangs())

	var req = httptest.NewRequest(http.MethodGet, "/barangs", nil)
	var res = httptest.NewRecorder()
	e.ServeHTTP(res, req)

	type ResponseData struct {
		Data    map[string]any `json:"data"`
		Message string         `json:"message"`
	}

	var tmp = ResponseData{}

	var resData = json.NewDecoder(res.Result().Body)
	err := resData.Decode(&tmp)

	assert.Equal(t, 200, res.Code)
	assert.Nil(t, err)
	assert.NotNil(t, tmp)
	assert.Equal(t, "success", tmp.Message)
	assert.Nil(t, tmp.Data)

}
