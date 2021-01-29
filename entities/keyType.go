package entities

type KeyType struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	Format                 string `json:"format"`
	AutomaticallyGenerated bool   `json:"automatic_generated"`
}
