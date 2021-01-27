package entities

type PixKey struct {
	ID        int64  `json:"id"`
	Value     string `json:"value"`
	AccountID int    `json:"account_id"`
	KeyTypeID int    `json:"key_type_id"`
}
