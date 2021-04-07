package Controllers

import (
	"belajar_golang/Config"
	"belajar_golang/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var taksKaryawan []Models.TaksKaryawan

type Input struct {
	IDTaks     int    `json:"id_taks"`
	IDKaryawan string `json:"id_karyawan"`
}

func AssignTaks(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taksKaryawan := Models.TaksKaryawan{
		IDTaks:     input.IDTaks,
		IDKaryawan: input.IDKaryawan,
	}

	result := Config.DB.Create(&taksKaryawan)

	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}
	c.JSON(http.StatusOK, res)
}

func ReadAssignedTaks(c *gin.Context) {
	result := Config.DB.
		Preload("Taks.KategoriTaks").
		Preload("Karyawan").
		Find(&taksKaryawan)

	res := gin.H{
		"data":  result.Value,
		"error": result.Error,
	}

	c.JSON(http.StatusOK, res)
}

func DeleteAssignedTaks(c *gin.Context) {
	row := Config.DB.Where("id_taks = ? AND id_karyawan = ?", c.Param("idTaks"), c.Param("idKaryawan")).Find(&taksKaryawan).RowsAffected
	if row == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data not found"})
		return
	}
	Config.DB.Where("id_taks = ? AND id_karyawan = ?", c.Param("idTaks"), c.Param("idKaryawan")).Delete(&taksKaryawan)
	c.JSON(http.StatusBadRequest, gin.H{"data": true})
}
