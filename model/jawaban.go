package model

import "gorm.io/gorm"

type Jawaban struct {
	gorm.Model
	SoalId      uint   `json:"soal_id"`
	KeyJawaban  string `gorm:"not null" json:"key_jawaban"`
	Content     string `gorm:"not null" json:"content"`
	ContentType string `gorm:"not null" json:"content_type"`
}
