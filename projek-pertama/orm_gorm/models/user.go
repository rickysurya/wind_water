package models

import "time"

type Product struct {
	ID			uint    `gorm:"primaryKey"` 
	Name		string	`gorm:"not null;unique;type:varchar(191)"`
	Brand		string	`gorm:"not null;unique;type:varchar(191)"`
	UserID		uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
type User struct {
	ID        	uint     ` gorm:"primaryKey"`
	Name      	string   `gorm:"not null;unique;type:varchar(191)"`
	Products	[]Product
	CreatedAt 	time.Time
	UpdatedAt	time.Time
}