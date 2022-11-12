package entities

type Employee struct {
	Id int64
	Name string `validate:"required"`
	Address string `validate:"required"`
	Position string `validate:"required"`
	Company string `validate:"required"`
	PresenceIn string `validate:"required" label:"Presence In"`
	PresenceOut string `validate:"required" label:"Presence Out"`
	Gender string `validate:"required"`
	EmployeeNumber string `validate:"required" label:"Employee Number"`
}