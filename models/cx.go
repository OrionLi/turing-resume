package models

import (
	"time"
)

type Cx struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(4);not null" json:"name" validate:"required,max=4"`
	StudentID   string    `gorm:"type:char(12);unique;not null" json:"student_id" validate:"required,len=12,numeric"`
	Major       string    `gorm:"type:varchar(16);not null" json:"major" validate:"required,max=16"`
	Direction   string    `gorm:"type:varchar(8);not null" json:"direction" validate:"required,oneof=前端 后端 计算机视觉 自然语言处理"`
	PhoneNumber string    `gorm:"type:char(11);not null" json:"phone_number" validate:"required,len=11,numeric"`
	Email       string    `gorm:"type:varchar(32);not null" json:"email" validate:"required,email"`
	Evaluation  string    `gorm:"type:varchar(2048)" json:"evaluation"`
	Skills      string    `gorm:"type:varchar(2048)" json:"skills"`
	Expectation string    `gorm:"type:varchar(2048)" json:"expectation"`
	Experience  string    `gorm:"type:varchar(2048)" json:"experience"`
	Others      string    `gorm:"type:varchar(2048)" json:"others"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"update_time"`
}
