package helper

import (
	"bot/db"
	"log"
)

// DB Create database connection
var DB = NewDatabase()

// NewDatabase creates and returns a new Database instance.
func NewDatabase() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(`Unable to connect to database`, err)
	}
	return client
}
