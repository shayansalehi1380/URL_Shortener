package Database

import (
	"log"
	"urlshortener/Models"
)

func ModelMigration() {

	err := DB.AutoMigrate(&Models.URL{})

	if err != nil {
		
		log.Println("Migration Faild...")
	}
	log.Println("Migration Success...")
}