package Controllers

import (
	"belajar_golang/Config"
	"belajar_golang/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var karyawan []Models.Karyawan

// GetKaryawan untuk get semua data karyawan di table karyawans
func GetKaryawan(c *gin.Context) {
	result := Config.DB.Find(&karyawan)

	res := gin.H{
		"error": result.Error,
		"data":  result.Value,
	}

	c.JSON(http.StatusOK, res)
}

//FindKaryawan cari karyawan dari id
func FindKaryawan(c *gin.Context) {
	result := Config.DB.First(&karyawan, "uuid = ?", c.Param("id"))
	if result.RowsAffected != 0 {
		res := gin.H{
			"error": result.Error,
			"data":  result.Value,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res := gin.H{
		"error": "Request id not found",
		"data":  result.Value,
	}
	c.JSON(http.StatusBadRequest, res)
}

//CreateKaryawan
type CreateKaryawanInput struct {
	UUID         string         `json:"id" binding:"required"`
	Nama         string         `json:"nama" binding:"required"`
	TglLahir     string         `json:"tgl_lahir" binding:"required"`
	JenisKelamin string         `json:"jenis_kelamin" binding:"required"`
	StatusAktif  int8           `json:"status_aktif" binding:"required"`
	Alamat       string         `json:"alamat" binding:"required"`
	CreatedAt    time.Time      `json:"created_at" binding:"required"`
	UpdatedAt    time.Time      `json:"updated_at" binding:"required"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" binding:"required"`
}

func CreateKaryawan(c *gin.Context) {
	var input CreateKaryawanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	karyawan := Models.Karyawan{
		UUID:         input.UUID,
		Nama:         input.Nama,
		TglLahir:     input.TglLahir,
		JenisKelamin: input.JenisKelamin,
		StatusAktif:  input.StatusAktif,
		Alamat:       input.Alamat,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
		DeletedAt:    input.DeletedAt,
	}
	result := Config.DB.Create(&karyawan)

	res := gin.H{
		"error": result.Error,
		"data":  karyawan,
	}

	c.JSON(http.StatusOK, res)
}

type UpdateKaryawanInput struct {
	Nama         string `json:"nama"`
	TglLahir     string `json:"tgl_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	StatusAktif  int8   `json:"status_aktif"`
	Alamat       string `json:"alamat"`
}

func UpdateKaryawan(c *gin.Context) {
	result := Config.DB.First(&karyawan, "uuid = ?", c.Param("id"))
	if result.RowsAffected == 0 {
		res := gin.H{
			"error": "Request id not found",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Validate input
	var input UpdateKaryawanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Config.DB.Model(&karyawan).Where("uuid = ?", c.Param("id")).Update(input)

	res := gin.H{
		"data": true,
	}
	c.JSON(http.StatusOK, res)
	return
}

//DeleteKaryawan
func DeleteKaryawan(c *gin.Context) {
	result := Config.DB.First(&karyawan, "uuid = ?", c.Param("id"))
	if result.RowsAffected != 0 {
		Config.DB.Delete(&karyawan, "uuid = ?", c.Param("id"))
		res := gin.H{
			"error": result.Error,
			"data":  true,
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res := gin.H{
		"error": "Request id not found",
		"data":  result.Value,
	}
	c.JSON(http.StatusBadRequest, res)
}
