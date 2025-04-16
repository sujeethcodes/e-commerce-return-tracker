package usecase

import (
	"e-commerce-return-tracker/constant"
	"e-commerce-return-tracker/entity"
	db "e-commerce-return-tracker/repository"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type ReturnProduct struct {
	Mysql *db.MysqlCon
}

func (r *ReturnProduct) ReturnProducts(returnProducts entity.ReturnProducts) error {
	fmt.Println("Enter Return Products Usecase")

	if r.Mysql.Connection == nil {
		zap.L().Info("Database connection failed")
		return errors.New("database connection not initialized")
	}

	if dbResult := r.Mysql.Connection.Debug().Table(constant.RETURN_PRODUCT_TABLE_NAME).Create(&returnProducts).Error; dbResult != nil {
		zap.L().Error("product return error", zap.Error(dbResult))
		return dbResult
	}

	fmt.Println("Inserted data:", returnProducts)
	return nil
}
