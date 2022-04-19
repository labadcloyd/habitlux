package models

type User struct {
	ID						uint				`json:"id"`
	Username			string			`json:"username" gorm:"unique"`
	Password			[]byte			`json:"-"` // putting a minus means not returning field
	// Habit_Names 	[]string		`json:"habit_names"`
}


