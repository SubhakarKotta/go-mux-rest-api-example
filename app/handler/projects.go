package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"go-mux-rest-api-example/app/model"
	"net/http"
)

func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllProjects")
}
func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProject")
}
func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}
func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateProject")
}
func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteProject")
}
