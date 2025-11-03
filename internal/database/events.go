package database
import "database/sql"

type EventModel struct{
	DB *sql.DB
}

type Event struct{
	Id int `json:"id"`
    OwnerId int `json:"ownerId"  binding:"required"`
	Name string `json:"name" binding:"required",min=3`
	Title string `json:"title"`
    Description string `json:"description" binding:"required",min=10`,
    Date time.Time `json:"date" binding:"required",`
    location string `json:"location" binding:"required",min=3`
}