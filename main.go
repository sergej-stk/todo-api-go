package main

import "github.com/google/uuid"

type todo struct {
	ID        uuid.UUID
	Item      string
	Completed bool
}
