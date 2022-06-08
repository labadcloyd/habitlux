package models

type User struct {
	ID						uint				`json:"id"`
	Username			string			`json:"username"`
	Password			[]byte			`json:"-"` // putting a minus means not returning field
}


