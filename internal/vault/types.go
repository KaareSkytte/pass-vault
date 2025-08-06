package vault

import "time"

type Entry struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	URL      string    `json:"url"`
	Notes    string    `json:"notes"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

type Vault struct {
	Version int              `json:"version"`
	Salt    []byte           `json:"salt"`
	Entries map[string]Entry `json:"entries"`
}

type EncryptedVault struct {
	Version       int    `json:"version"`
	Salt          []byte `json:"salt"`
	EncryptedData []byte `json:"encrypted_data"`
}
