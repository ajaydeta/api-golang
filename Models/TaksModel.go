package Models

import (
	"time"
)

type Taks struct {
	ID         int       `json:"id"`
	Nama       string    `json:"nama"`
	DueDate    time.Time `json:"due_date"`
	Deskripsi  string    `json:"deskripsi"`
	IDKategori int       `json:"id_kategori" gorm:"foreignKey:id_kategori"`
}
