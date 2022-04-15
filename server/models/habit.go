package models

type Habit struct {
	ID											uint		`json:"id"`
	Owner_ID								string	`json:"owner_id"`
	Date_Created						string	`json:"date_created"` //gorm:"unique"
	Habit_Name							string	`json:"habit_name"`
	Comment									string	`json:"comment"`
	Required_Repeat_Count		uint		`json:"required_repeat_count"`
	Repeat_Count						uint		`json:"repeat_count"`
}


