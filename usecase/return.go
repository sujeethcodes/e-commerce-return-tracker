package usecase

import (
	"e-commerce-return-tracker/constant"
	"e-commerce-return-tracker/entity"
	db "e-commerce-return-tracker/repository"
	"fmt"
)

type ReturnProduct struct {
	Mysql *db.MysqlCon
}

func (r *ReturnProduct) ReturnProducts(returnProducts entity.ReturnProducts) error {
	fmt.Println("Enter Return Products Usecase")
	if r.Mysql.Connection == nil {
		fmt.Println("DB coneection  failed")
	}
	err := r.Mysql.Connection.Table(constant.RETURN_PRODUCT_TABLE_NAME).Create(&returnProducts)
	if err != nil {
		fmt.Errorf("Return Products Insert failed")
	}
	return nil
}
