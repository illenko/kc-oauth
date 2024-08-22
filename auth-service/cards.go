package main

type Card struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	Owner  string `json:"owner"`
}

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

func (s *CardService) GetCards() []Card {
	return []Card{
		{ID: "1", Number: "1234-5678-9012-3456", Owner: "John Doe"},
		{ID: "2", Number: "2345-6789-0123-4567", Owner: "Jane Smith"},
	}
}
