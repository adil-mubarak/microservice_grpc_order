package models

type Order struct {
	ID             uint   `gorm:"primarykey"`
	ProductID      uint   `grom:"not null"`
	UserID         uint   `grom:"not null"`
	Quantity       int64  `gorm:"not null"`
	Order_status   string `gorm:"varchar(50);default:'Pending'"`
	Payment_status string `gorm:"varchar(50);default:'Pending'"`
}
