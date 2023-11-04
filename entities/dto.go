package entities

type BookInput struct {
	Book        string `json:"book"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
