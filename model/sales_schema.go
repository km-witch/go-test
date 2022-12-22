package model

import (
	"pkg/configs"
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	Id           int       `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Product_id   int       `gorm:"column:product_id" json:"product_id" binding:"required"`
	Status       bool      `gorm:"column:status;default=false" json:"status"`
	Sale_type    int       `gorm:"column:sale_type" json:"sale_type"`
	Sale_count   int       `gorm:"column:sale_count" json:"sale_count"`
	Won_price    int       `gorm:"column:won_price" json:"won_price"`
	Doller_price int       `gorm:"column:doller_price" json:"doller_price"`
	Created_time time.Time `gorm:"autoCreateTime" json:"created_time"`
}

type Saleslog struct {
	Id           int `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Type         int `gorm:"column:type" json:"type" binding:"required"`
	Sale_id      int `gorm:"column:sale_id" json:"collection_id" binding:"required"`
	User_id      int `gorm:"column:user_id" json:"user_id" binding:"required"`
	Nft_id       int `gorm:"column:nft_id" json:"nft_id" binding:"required"`
	Amount       int `gorm:"column:amount" json:"amount" binding:"required"`
	Won_price    int `gorm:"column:won_price" json:"won_price"`
	Doller_price int `gorm:"column:doller_price" json:"doller_price"`
}

var SalesSchema *Sale
var SalesLogSchema *Saleslog

// 세일 조회 (ByID)
func (s *Sale) GetSalesById(db *gorm.DB, saleId string) (Sale, error) {
	var sales Sale
	configs.DB.Model(&sales).First(&sales, saleId)
	return sales, nil
}

// 세일 로그 조회 (길이)
func (sl *Saleslog) GetSalesLog(db *gorm.DB, sid string, uid string) int {
	var sales_logs []Saleslog
	configs.DB.Model(&sales_logs).Where("sale_id=? AND user_id=?", sid, uid).Find(&sales_logs)
	return len(sales_logs)
}

// SaleLog 생성
func (sl *Saleslog) CreateSalesLog(db *gorm.DB, salelog Saleslog) (Saleslog, error) {
	var userInput Saleslog
	userInput = salelog
	db.Create(&userInput)
	return userInput, nil
}
