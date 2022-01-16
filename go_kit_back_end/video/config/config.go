package config

type RDSConfig struct {
	Dsn        string `json:"dsn"`
	StorageDir string `json:"storageDir"`
}
