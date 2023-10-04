package helper

import "github.com/google/uuid"

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g Generator) GenerateUUID() string {
	return uuid.NewString()
}
