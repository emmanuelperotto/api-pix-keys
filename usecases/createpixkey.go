package usecases

import (
	"errors"
	"log"
	"pixkeys/entities"
	"pixkeys/infra/middlewares"
	"pixkeys/repositories"
)

func CreatePixKey(pixKey entities.PixKey) (entities.PixKey, error) {
	pixKey.AccountID = middlewares.CurrentAccountID

	//FIXME: extract validation to validations pkg
	if pixKey.AccountID == 0 {
		log.Println("[CreatePixKey Error] AccountID can't be null/empty")
		return pixKey, errors.New("AccountID can't be null/empty")
	}

	pixKey, err := repositories.PixKeyRepo.Create(pixKey)

	if err != nil {
		log.Println("[CreatePixKey Error]", err)
		return pixKey, err
	}

	log.Println("PixKey Created")
	return pixKey, nil
}
