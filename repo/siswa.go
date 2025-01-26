package repo

import (
	"fmt"
	"myapp/model"
    "myapp/db"
)


func CreateSiswa(siswa model.Siswa) (model.Siswa, error) {
    if db.DB == nil {
        return model.Siswa{}, fmt.Errorf("koneksi database belum diinisialisasi")
    }
    if err := db.DB.Create(&siswa).Error; err != nil {
        return model.Siswa{}, err
    }
    return siswa, nil
}

func GetAllSiswa() ([]model.Siswa, error) {
    var siswa []model.Siswa


    if err := db.DB.Find(&siswa).Error; err != nil {
        return nil, err
    }

    return siswa, nil
}

func DeleteSiswa(id string) error {
    if db.DB == nil {
        return fmt.Errorf("koneksi database belum diinisialisasi")
    }
    if err := db.DB.Where("id = ?", id).Delete(&model.Siswa{}).Error; err != nil {
        return err
    }
    return nil
}


func DetailSiswa(id string) (model.Siswa, error) {
    var siswa model.Siswa
    if err := db.DB.Where("id = ?", id).First(&siswa).Error; err != nil {
        return model.Siswa{}, err
    }

return siswa, nil
}


func UpdateSiswa(id string, siswa model.Siswa) (model.Siswa, error) {
    if db.DB == nil {
        return model.Siswa{}, fmt.Errorf("koneksi database belum diinisialisasi")
    }
    if err := db.DB.Model(&siswa).Where("id = ?", id).Updates(&siswa).Error; err != nil {
        return model.Siswa{}, err
    }
    return siswa, nil
}