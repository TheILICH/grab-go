package repositories

import (
    "errors"
    "go_final/models"
    "gorm.io/gorm"

    "fmt"
)

type BasketRepository interface {
    GetBasket(userID uint) (models.Basket, error)
    UpdateBasket(userID uint, productID uint, cnt int) error
}

type basketRepository struct {
    connection *gorm.DB
}

func NewBasketRepository() BasketRepository {
    return &basketRepository{
        connection: DB(),
    }
}

func (db *basketRepository) GetBasket(userID uint) (models.Basket, error) {
    var basket models.Basket
    err := db.connection.Preload("BasketItems.Product").Where("user_id = ?", userID).First(&basket).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // Create a new basket if none exists
            basket = models.Basket{
                UserID: userID,
            }
            if err := db.connection.Create(&basket).Error; err != nil {
                return basket, err
            }
            return basket, nil
        }
        return basket, err
    }
    return basket, nil
}

func (db *basketRepository) UpdateBasket(userID uint, productID uint, cnt int) error {
    return db.connection.Transaction(func(tx *gorm.DB) error {
        var basket models.Basket
        err := tx.Where("user_id = ?", userID).First(&basket).Error
        if err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                // Create basket if not exists
                basket = models.Basket{
                    UserID: userID,
                }
                if err := tx.Create(&basket).Error; err != nil {
                    return err
                }
            } else {
                return err
            }
        }


        var basketItem models.BasketItem
        // Check if the product is already in the basket
        err = tx.Where("basket_id = ? AND product_id = ?", basket.ID, productID).First(&basketItem).Error
        if err != nil {
            fmt.Printf("basket_id = %d and product_id = %d\n", basket.ID, productID)
            if errors.Is(err, gorm.ErrRecordNotFound) {
                if cnt <= 0 {
                    // Cannot subtract from non-existing item
                    return nil
                }
                // Create new BasketItem
                basketItem = models.BasketItem{
                    BasketID:  basket.ID,
                    ProductID: productID,
                    Quantity:  cnt,
                }
                if err := tx.Create(&basketItem).Error; err != nil {
                    return err
                }
            } else {
                fmt.Printf("IT GOT HERE!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n\n\n\n")
                return err
            }
        } else {
            // Update existing BasketItem
            newQuantity := basketItem.Quantity + cnt
            if newQuantity <= 0 {
                // Remove the BasketItem
                if err := tx.Delete(&basketItem).Error; err != nil {
                    return err
                }
            } else {
                basketItem.Quantity = newQuantity
                if err := tx.Save(&basketItem).Error; err != nil {
                    return err
                }
            }
        }
        return nil
    })
}
