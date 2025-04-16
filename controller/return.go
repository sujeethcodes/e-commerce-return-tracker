package controller

import (
	"e-commerce-return-tracker/constant"
	"e-commerce-return-tracker/entity"
	db "e-commerce-return-tracker/repository"
	"e-commerce-return-tracker/usecase"
	"e-commerce-return-tracker/utils"
	"net/http"

	"fmt"

	"github.com/labstack/echo/v4"
)

type ReturnProduct struct {
	Mysql *db.MysqlCon
}

func (r *ReturnProduct) ReturnProducts(c echo.Context) error {
	fmt.Println("Enter Return Products controller")
	returnProdutUsecase := usecase.ReturnProduct{
		Mysql: r.Mysql,
	}
	req := entity.ReturnProducts{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadGateway, entity.Response{
			Status:  401,
			Message: constant.INVAILD_PARAMS,
			Error:   err.Error(),
		})
	}
	returnDate, err := utils.FormateDate()
	if err != nil {
		return c.JSON(http.StatusBadGateway, entity.Response{
			Status:  401,
			Message: constant.RETURN_DATE_REQUIRED,
			Error:   err.Error(),
		})
	}

	req.ReturnDate = returnDate

	if err := returnProdutUsecase.ReturnProducts(req); err != nil {
		return c.JSON(http.StatusBadGateway, entity.Response{
			Status:  401,
			Message: constant.PRODUCT_RETURN_INSERT_FAILED,
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.Response{
		Status:  200,
		Message: constant.PRODUCT_RETURN_SUCCESS,
	})
}
