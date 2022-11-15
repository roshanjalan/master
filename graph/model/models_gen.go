// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AdditionInput struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type AdditionResponse struct {
	Number3 *int `json:"number3"`
}

type CalInput struct {
	Operator *string `json:"operator"`
	Number1  *int    `json:"number1"`
	Number2  *int    `json:"number2"`
}

type CalResponse struct {
	Result *int `json:"result"`
}

type DivInput struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type DivResponse struct {
	Number3 *int `json:"number3"`
}

type ExcelOutput struct {
	Status *string `json:"status"`
}

type MulInput struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type MulResponse struct {
	Number3 *int `json:"number3"`
}

type StudentData struct {
	Inp *string `json:"inp"`
}

type StudentInput struct {
	ID         *string `json:"ID"`
	FirstName  string  `json:"FirstName"`
	LastName   *string `json:"LastName"`
	Gender     *string `json:"Gender"`
	Roll       *int    `json:"Roll"`
	BloodGroup string  `json:"BloodGroup"`
	Dob        string  `json:"DOB"`
	PhNo       *int    `json:"PhNo"`
}

type StudentOutput struct {
	ID         *string `json:"ID"`
	FirstName  *string `json:"FirstName"`
	LastName   *string `json:"LastName"`
	Gender     *string `json:"Gender"`
	Roll       *int    `json:"Roll"`
	BloodGroup *string `json:"BloodGroup"`
	Dob        *string `json:"DOB"`
	PhNo       *int    `json:"PhNo"`
}

type StudentResponse struct {
	ID     *string   `json:"ID"`
	Status []*string `json:"status"`
}

type SubInput struct {
	Number1 *int `json:"number1"`
	Number2 *int `json:"number2"`
}

type SubResponse struct {
	Number3 *int `json:"number3"`
}