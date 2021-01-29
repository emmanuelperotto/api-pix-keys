package usecases

import (
	"errors"
	"github.com/google/uuid"
	"log"
	"pixkeys/entities"
	"pixkeys/repositories"
)

func CreatePixKey(pixKey entities.PixKey) (entities.PixKey, error) {
	//FIXME: extract validation to validations pkg
	if pixKey.AccountID == 0 {
		log.Println("[CreatePixKey Error] AccountID can't be null/empty")
		return pixKey, errors.New("AccountID can't be null/empty")
	}

	keyType, err := repositories.KeyTypeRepo.FindByID(pixKey.KeyTypeID)
	if err != nil {
		log.Println("error keytype", err)
		return pixKey, err
	}

	if keyType.AutomaticallyGenerated {
		log.Println("Setting pix key value to a random uuid")
		pixKey.Value = uuid.NewString()
	}

	pixKey, err = repositories.PixKeyRepo.Create(pixKey)

	if err != nil {
		log.Println("[CreatePixKey Error]", err)
		return pixKey, err
	}

	log.Println("PixKey Created")
	return pixKey, nil
}
