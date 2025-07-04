package initializers

import (
	"log"

	"github.com/Kariqs/mesh-art-gallery-api/models"
)

func SyncDatabase() {
	DB.AutoMigrate(models.User{}, models.Product{}, models.Sale{})
	log.Println("Database synced successfully")
}
