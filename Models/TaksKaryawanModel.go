package Models

type TaksKaryawan struct {
	IDTaks     int      `json:"id_taks"`
	IDKaryawan string   `json:"id_karyawan"`
	Taks       Taks     `json:"taks" gorm:"Foreignkey:IDTaks;association_foreignkey:ID"`
	Karyawan   Karyawan `json:"karyawan" gorm:"Foreignkey:IDKaryawan;association_foreignkey:UUID"`
}
