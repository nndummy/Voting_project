package services

import (
	"voting_system/app/database"

	"github.com/jinzhu/gorm"
)

var votingdb *gorm.DB

func init() {
	votingdb = database.GetConnection()
}
