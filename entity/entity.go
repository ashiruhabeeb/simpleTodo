package entity

type Todo struct {
	TodoID		int		`json:"todo_id"`
	UserID		int		`json:"user_id"`
	Title		string	`json:"title"`
	Description string	`json:"description"`
	Completed	bool	`json:"completed"`
	StartAt		string	`json:"start_at"`
	EndAt		string	`json:"end_at"`
}

type User struct {
	UserID		int		`json:"user_id,omitempty"`
	Username	string	`json:"user_name"`
	FullName	string	`json:"full_name"`
	Email		string	`json:"e-mail"`
	Password	string	`json:"password"`
	Phone		string	`json:"phone"`
	Address		[]Address	`json:"address,omitempty"`
	Avatar		string	`json:"avatar,omitempty"`
	DOB			string	`json:"dob,omitempty"`
	CreatedAt	string	`json:"created_at"`
	UpdatedAt	string	`json:"updated_at,omitempty"`
}

type Address struct {
	AddressID	int		`json:"address_id"`
	HouseNumber	string	`json:"house_number"`
	StreetName	string	`json:"street_name"`
	LocalArea	string	`json:"local_area"`
	State		string	`json:"state"`
	Country		string	`json:"country"`
}
