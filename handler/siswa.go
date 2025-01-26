package handler

import (
	"log"
	"myapp/helper"
	"myapp/model"
	"myapp/repo"
	"net/http"

	"github.com/labstack/echo/v4"
)


func CreateSiswa(c echo.Context) error {
    var siswa model.Siswa

    // Bind the incoming JSON to the Siswa struct
    if err := c.Bind(&siswa); err != nil {
        log.Printf("Error binding request body: %v", err)
        return c.JSON(http.StatusBadRequest, helper.Error("Invalid request body", http.StatusBadRequest))
    }
    createdSiswa, err := repo.CreateSiswa(siswa)
    if err != nil {
        log.Printf("Error creating Siswa: %v", err)
        return c.JSON(http.StatusInternalServerError, helper.Error("Failed to create Siswa", http.StatusInternalServerError))
    }
    return c.JSON(http.StatusOK, helper.Succes("Siswa created successfully", http.StatusOK, createdSiswa))
}

func GetAllSiswa(c echo.Context) error {
    siswa, err := repo.GetAllSiswa()
    if err != nil {
        log.Printf("Error getting Siswa: %v", err)
    }

    return c.JSON(http.StatusOK, helper.Succes("Siswa retrieved successfully", http.StatusOK, siswa))
}


func DeleteSiswa(c echo.Context) error {
    id := c.Param("id")

    if !existsSiswa(id) {
        return c.JSON(http.StatusNotFound, helper.Error("Siswa not found", http.StatusNotFound))
    }
    
    if err := repo.DeleteSiswa(id); err != nil {
        log.Printf("Error deleting Siswa: %v", err)
        return c.JSON(http.StatusInternalServerError, helper.Error("Failed to delete Siswa", http.StatusInternalServerError))
    }

    return c.JSON(http.StatusOK, helper.Succes("Siswa deleted successfully", http.StatusOK, nil))
}


func DetailSiswa(c echo.Context) error {
    id := c.Param("id")

    siswa, err := repo.DetailSiswa(id)
     
    if err != nil {
        log.Printf("Error getting Siswa: %v", err)
        return c.JSON(http.StatusInternalServerError, helper.Error("Failed to get Siswa", http.StatusInternalServerError))
    }

    return c.JSON(http.StatusOK, helper.Succes("Siswa retrieved successfully", http.StatusOK, siswa))
}


func UpdateSiswa(c echo.Context) error {
    id := c.Param("id")


    if !existsSiswa(id) {
        return c.JSON(http.StatusNotFound, helper.Error("Siswa not found", http.StatusNotFound))
    }
    var siswa model.Siswa
    if err := c.Bind(&siswa); err != nil {
        log.Printf("Error binding request body: %v", err)
        return c.JSON(http.StatusBadRequest, helper.Error("Invalid request body", http.StatusBadRequest))
    }

    updateSiswa, err := repo.UpdateSiswa(id, siswa)
    if err != nil {
        log.Println("Error updating Siswa:", err)
        return c.JSON(http.StatusInternalServerError, helper.Error("Failed to update Siswa", http.StatusInternalServerError))
    }

    return c.JSON(http.StatusOK, helper.Succes("Siswa updated successfully", http.StatusOK, updateSiswa))
    

}


func existsSiswa(id string) bool {
    _, err := repo.DetailSiswa(id)
    return err == nil
}