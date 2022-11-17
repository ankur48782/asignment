package controllers

type UserInput struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}
