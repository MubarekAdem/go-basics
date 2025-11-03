package database

type Model struct {
	Users UserModel
	Events EventModel
	Attendees AttendeeModel
}

func NewModels(db *sql.DB)models {
	return models{
		Users: UserModel{DB: db},
		Events: EventModel{DB: db},
		Attendees: AttendeeModel{DB: db},
	}
}