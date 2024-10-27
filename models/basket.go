package models

import (
    "gorm.io/gorm"
)

type Basket struct {
    gorm.Model
    UserID      uint         `json:"user_id"`
    User        User         `json:"user"`
    BasketItems []BasketItem `gorm:"foreignKey:BasketID" json:"basket_items"`
}

type BasketItem struct {
    gorm.Model
    BasketID  uint    `json:"basket_id"`
    Basket    Basket  `gorm:"foreignKey:BasketID" json:"-"`
    ProductID uint    `json:"product_id"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int     `json:"quantity"`
}
