package repositories

import (
	"context"
	"pixkeys/entities"
	"pixkeys/infra"
)

type pixKeyRepo struct{}

type PixKeyCreator interface {
	Create(entities.PixKey) (entities.PixKey, error)
}

func (ar pixKeyRepo) Create(pixKey entities.PixKey) (entities.PixKey, error) {
	ctx := context.Background()
	result, err := infra.DB.ExecContext(ctx,
		"INSERT INTO api_pix_keys.PixKeys (Value, AccountID, KeyTypeID) VALUES (?, ?, ?)",
		pixKey.Value,
		pixKey.AccountID,
		pixKey.KeyTypeID)

	if err != nil {
		return pixKey, err
	}

	id, _ := result.LastInsertId()
	pixKey.ID = id

	return pixKey, nil
}

var (
	PixKeyRepo = pixKeyRepo{}
)
