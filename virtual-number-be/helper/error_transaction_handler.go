package helper

import (
	"fmt"
	"log"

	"github.com/imnzr/virtual-number-service/models"
)

func ErrorTransaction(err error) (models.User, error) {
	if err != nil {
		log.Printf("transaction error: %v", err)
		return models.User{}, fmt.Errorf("failed to start transaction: %w", err)
	}
	return models.User{}, nil
}
