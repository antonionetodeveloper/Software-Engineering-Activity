package operations

type Student struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	CPF    string `json:"CPF"`
	Phone  string `json:"phone"`
	Gender string `json:"gender"`
}

type Admin struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
