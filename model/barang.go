package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	Nama    string `json:"nama" gorm:"type:varchar(255)"`
	Pemilik string `gorm:"type:varchar(20)"`
}

type BarangModel struct {
	db *gorm.DB
}

func (bm *BarangModel) Init(db *gorm.DB) {
	bm.db = db
}

func (bm *BarangModel) Insert(newItem Barang) *Barang {
	if err := bm.db.Create(&newItem).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newItem
}

func (bm *BarangModel) GetAllBarang() []Barang {
	var listBarang = []Barang{}

	if err := bm.db.Find(&listBarang).Error; err != nil {
		logrus.Error("Model : Get data error, ", err.Error())
		return nil
	}

	return listBarang
}

func (bm *BarangModel) Delete(id int) {
	var deletdData = Barang{}
	deletdData.ID = uint(id)
	if err := bm.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}

func (bm *BarangModel) UpdateData(updatedData Barang) bool {
	var qry = bm.db.Save(&updatedData)
	if err := qry.Error; err != nil {
		logrus.Error("Model : Update error, ", err.Error())
		return false
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return false
	}

	return true
}

func (bm *BarangModel) UpdateData2(updatedData Barang) bool {
	var qry = bm.db.Table("barangs").Where("id = ?", updatedData.ID).Update("nama", updatedData.Nama)
	if err := qry.Error; err != nil {
		logrus.Error("Model : Update error, ", err.Error())
		return false
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return false
	}

	return true
}
