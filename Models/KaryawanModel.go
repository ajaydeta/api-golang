package Models

import (
	"gorm.io/gorm"
	"time"
)

type Karyawan struct {
	UUID         string         `json:"id" gorm:"primary_key"`
	Nama         string         `json:"name"`
	TglLahir     string         `json:"tgl_lahir"`
	JenisKelamin string         `json:"jenis_kelamin"`
	StatusAktif  int8           `json:"status_aktif"`
	Alamat       string         `json:"alamat"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
