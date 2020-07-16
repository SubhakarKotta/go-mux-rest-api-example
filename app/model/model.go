package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	Title string `gorm:"unique" json:"title"`
	Tasks []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
	ProjectID uint       `json:"project_id"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
