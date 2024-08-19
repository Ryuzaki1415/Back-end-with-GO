package types


//here we can have every type modular so that we can use them throught the project.

type RegisterUserPayload struct{ 
	Firstname string `json:"firstName"`
	LastName string  `json:"lastName"`
	Email string	 `json:"email"`
	Password string	 `json:"password"`

}