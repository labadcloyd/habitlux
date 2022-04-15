package controllers

type ReqHabit struct {
	ID											uint		`json:"id"`
	Owner_ID								string	`json:"owner_id"`
	Date_Created						string	`json:"date_created" gorm:"unique"`
	Habit_Name							string	`json:"habit_name" validate:"required,min=1,max=32"`
	Comment									string	`json:"comment"`
	Required_Repeat_Count		uint		`json:"required_repeat_count"`
	Repeat_Count						uint		`json:"repeat_count"`
}