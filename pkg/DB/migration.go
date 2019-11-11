package DB

import (
	"CASystem/pkg/infra"
	"CASystem/pkg/infra/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	DB = infra.GetDBConnection()
	dbMigration()
}

func dbMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.People{})
	DB.AutoMigrate(&model.Group{})
	DB.AutoMigrate(&model.Faculty{})
	DB.AutoMigrate(&model.Role{})
	DB.AutoMigrate(&model.Authority{})
	DB.AutoMigrate(&model.LendingHistory{})
	DB.AutoMigrate(&model.Equipment{})
	DB.AutoMigrate(&model.UserProject{})
	DB.AutoMigrate(&model.Project{})
	DB.AutoMigrate(&model.ProjectProgress{})
	DB.AutoMigrate(&model.ProjectStatus{})
	DB.AutoMigrate(&model.UsersEvent{})
	DB.AutoMigrate(&model.Event{})
	DB.AutoMigrate(&model.EventCategory{})
	DB.AutoMigrate(&model.QuietionnaireContent{})
	DB.AutoMigrate(&model.QuestionType{})
	DB.AutoMigrate(&model.Questionnaire{})
	DB.AutoMigrate(&model.Budget{})
}
