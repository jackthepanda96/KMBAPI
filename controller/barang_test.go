package controller

import (
	"bytes"
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
func (mm *MockModel) Delete(id int) {}
func (mm *MockModel) UpdateData(updatedData model.Barang) bool {
	return false
}
func (mm *MockModel) UpdateData2(updatedData model.Barang) bool {
	return false
}

type SuccessMockModel struct{}

func (smm *SuccessMockModel) Insert(newItem model.Barang) *model.Barang {
	return &newItem
}
func (smm *SuccessMockModel) GetAllBarang() []model.Barang {
	return nil
}
func (smm *SuccessMockModel) Delete(id int) {}
func (smm *SuccessMockModel) UpdateData(updatedData model.Barang) bool {
	return false
}
func (smm *SuccessMockModel) UpdateData2(updatedData model.Barang) bool {
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

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Nil(t, err)
	assert.NotNil(t, tmp)
	assert.Equal(t, "success", tmp.Message)
	assert.Nil(t, tmp.Data)
}

func TestInsert(t *testing.T) {
	var mdl = MockModel{}
	var ctl = NewBarangControllInterface(&mdl)
	var e = echo.New()
	e.POST("/barangs", ctl.Insert())

	t.Run("Bind Error", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodPost, "/barangs", bytes.NewReader([]byte(`{"name":"buku", "pemilik":"saya"}`)))

		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponseData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponseData{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "invalid user input", tmp.Message)
	})

	t.Run("Server error", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodPost, "/barangs", bytes.NewReader([]byte(`{"name":"buku", "pemilik":"saya"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponseData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponseData{}

		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "cannot process data, something happend", tmp.Message)
	})

	t.Run("Succes Insert", func(t *testing.T) {
		var mdl = SuccessMockModel{}
		var ctl = NewBarangControllInterface(&mdl)
		var e = echo.New()
		e.POST("/barangs", ctl.Insert())

		var req = httptest.NewRequest(http.MethodPost, "/barangs", bytes.NewReader([]byte(`{"nama":"buku", "pemilik":"saya"}`)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)

		type ResponseData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.NoError(t, err)
		assert.Equal(t, "success", tmp.Message)
		assert.NotNil(t, tmp)
		assert.Equal(t, "buku", tmp.Data["nama"])
	})
}

func TestDelete(t *testing.T) {
	var mdl = MockModel{}
	var ctl = NewBarangControllInterface(&mdl)

	var e = echo.New()
	e.DELETE("/barangs/:id", ctl.Delete())

	t.Run("Invalid ID", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/barangs/satu", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)
		type ResponseData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("Success Delete", func(t *testing.T) {
		var req = httptest.NewRequest(http.MethodDelete, "/barangs/1", nil)
		var res = httptest.NewRecorder()

		e.ServeHTTP(res, req)
		type ResponseData struct {
			Data    map[string]any `json:"data"`
			Message string         `json:"message"`
		}

		var tmp = ResponseData{}
		var resData = json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

}
