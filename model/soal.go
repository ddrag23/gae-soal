package model

import "gorm.io/gorm"

type Soal struct {
	gorm.Model
	UserId       string    `gorm:"not null" json:"user_id"`
	Title        string    `gorm:"not null" json:"title"`
	TypeSoal     string    `gorm:"not null" json:"type_soal"`
	Image        string    `gorm:"not null" json:"image"`
	KunciJawaban string    `gorm:"not null" json:"kunci_jawaban"`
	Answers      []Jawaban `gorm:"foreignKey:SoalId" json:"answers"`
}
