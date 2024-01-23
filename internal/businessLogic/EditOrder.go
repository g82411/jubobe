package businessLogic

import (
	"JuboTest/internal/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func EditOrder(ctx context.Context, previousOrderId int, prevOrderMessage, newMessage string) error {

	db := ctx.Value("db").(*gorm.DB)
	tx := db.Begin()
	var prevOrder models.Order
	prevTx := tx.Model(&models.Order{}).Where("id = ?", previousOrderId).Find(&prevOrder)
	if prevTx.Error != nil {
		tx.Rollback()
		return prevTx.Error
	}
	if prevOrder.Text != prevOrderMessage {
		tx.Rollback()
		return fmt.Errorf("previousOrder.Text is not equal to prevOrder.Text")
	}
	var orderChangeLog models.OrderLog
	orderChangeLog.OrderID = previousOrderId
	orderChangeLog.Order = prevOrder
	orderChangeLog.Text = prevOrder.Text
	err := tx.Model(&models.OrderLog{}).Create(&orderChangeLog)
	if err.Error != nil {
		tx.Rollback()
		return fmt.Errorf("error creating orderChangeLog: %v", err)
	}
	err = tx.Model(&models.Order{}).Where("id = ?", previousOrderId).Update("text", newMessage)
	if err.Error != nil {
		tx.Rollback()
		return fmt.Errorf("error updating order: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		// 如果提交事务失败，也要回滚
		tx.Rollback()
		return fmt.Errorf("tx.Commit error: %v", err)
	}
	return tx.Error
}
