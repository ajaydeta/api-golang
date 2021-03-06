package Routes

import (
	"belajar_golang/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	grup1 := r.Group("/api")
	{
		grup1.GET("/karyawan", Controllers.GetKaryawan)
		grup1.GET("/karyawan/:id", Controllers.FindKaryawan)
		grup1.POST("/karyawan", Controllers.CreateKaryawan)
		grup1.DELETE("/karyawan/:id", Controllers.DeleteKaryawan)
		grup1.PATCH("/karyawan/:id", Controllers.UpdateKaryawan)

		grup1.GET("/taks", Controllers.GetTaks)
		grup1.GET("/taks/:id", Controllers.FindTaks)
		grup1.POST("/taks", Controllers.CreateTaks)
		grup1.PATCH("/taks/:id", Controllers.UpadateTaks)
		grup1.DELETE("/taks/:id", Controllers.DeleteTaks)

		grup1.GET("/kategori", Controllers.GetKategori)
		grup1.POST("/kategori", Controllers.CreateKategori)
		grup1.PATCH("/kategori/:id", Controllers.UpdateKategori)
		grup1.DELETE("/kategori/:id", Controllers.DeleteKategori)

		grup1.POST("/assign-new-taks", Controllers.AssignTaks)
		grup1.GET("/assigned-taks", Controllers.ReadAssignedTaks)
		grup1.DELETE("/assigned-taks/:idTaks/:idKaryawan", Controllers.DeleteAssignedTaks)
	}
	return r
}
