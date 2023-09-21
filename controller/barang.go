package controller

import (
	"net/http"
	"restEcho1/helper"
	"restEcho1/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BarangController struct {
	model model.BarangModel
}

func (bc *BarangController) InitUserController(bm model.BarangModel) {
	bc.model = bm
}

func (bc *BarangController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Barang{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var res = bc.model.Insert(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}

func (bc *BarangController) GetBarangs() echo.HandlerFunc {
	return func(c echo.Context) error {
		var res = bc.model.GetAllBarang()

		return c.JSON(http.StatusOK, helper.FormatResponse("success", res))
	}
}

func (bc *BarangController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		bc.model.Delete(cnv)

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (bc *BarangController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramId = c.Param("id")
		cnv, err := strconv.Atoi(paramId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid id", nil))
		}

		var input = model.Barang{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}
		input.ID = uint(cnv)

		var res = bc.model.UpdateData2(input)

		if !res {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}
