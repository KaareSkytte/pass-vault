package vault

import (
	"fmt"
	"time"
)

func NewVault() *Vault {
	entries := make(map[string]Entry)
	vault := Vault{
		Version: 1,
		Salt:    nil,
		Entries: entries,
	}
	return &vault
}

func (v *Vault) AddEntry(name string, entry Entry) error {
	if _, exists := v.Entries[name]; exists {
		return fmt.Errorf("Entry already exists: %v")
	}

	entry.Created = time.Now()
	entry.Updated = time.Now()

	v.Entries[name] = entry
	return nil
}

func (v *Vault) Save(password, filename string) error {

}

func LoadVault(password, filename string) (*Vault, error) {

}
