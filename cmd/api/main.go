package main
import "fmt"

func main() {
	db ,err := sql.Open("sqlite3", "./data.db")
	if err!= nil{
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
}