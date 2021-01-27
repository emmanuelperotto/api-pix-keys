package usecases

import (
	"log"
	"pixkeys/entities"
	"pixkeys/repositories"
)

func CreatePixKey(pixKey entities.PixKey) (entities.PixKey, error) {
	pixKey, err := repositories.PixKeyRepo.Create(pixKey)

	if err != nil {
		log.Println("[CreatePixKey Error]", err)
		return pixKey, err
	}

	log.Println("PixKey Created")
	return pixKey, nil
}
