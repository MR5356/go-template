package demo

import "github.com/MR5356/go-template/pkg/middleware/database"

type Demo struct {
	ID    uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Title string `json:"title" gorm:"not null"`

	database.BaseModel
}

func (*Demo) TableName() string {
	return "demo_table_name"
}
