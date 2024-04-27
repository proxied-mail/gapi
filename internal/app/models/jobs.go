package models

type Jobs struct {
	ID          int
	Queue       string
	Payload     string
	Attempts    int
	ReversedAt  int
	AvailableAt int
	CreatedAt   int
}
