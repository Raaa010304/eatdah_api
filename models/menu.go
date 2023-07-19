package models

type Menu struct {
	IdMenu      int64  `gorm:"primarykey" json:"id_menu"`
	Kategori    string `gorm:"type:varchar(300)" json:"kategori"`
	NamaMenu    string `gorm:"type:varchar(300)" json:"nama_menu"`
	Warung      string `gorm:"type:varchar(300)" json:"warung"`
	Harga       string `gorm:"type:varchar(300)" json:"harga"`
	HargaDiskon string `gorm:"type:varchar(300)" json:"harga_diskon"`
	Rating      string `gorm:"type:varchar(300)" json:"rating"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
