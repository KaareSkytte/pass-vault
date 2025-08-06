package vault

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kaareskytte/pass-vault/internal/crypto"
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
		return fmt.Errorf("Entry already exists")
	}

	entry.Created = time.Now()
	entry.Updated = time.Now()

	v.Entries[name] = entry
	return nil
}

func (v *Vault) Save(password, filename string) error {
	if v.Salt == nil {
		salt, err := crypto.GenerateSalt()
		if err != nil {
			return err
		}
		v.Salt = salt
	}

	encryptionKey := crypto.GenerateEncryptionKey(password, v.Salt)

	vault, err := json.Marshal(v)
	if err != nil {
		return err
	}

	encryptedVault, err := crypto.Encrypt(vault, encryptionKey)
	if err != nil {
		return err
	}

	encryptedVaultFile := EncryptedVault{
		Version:       v.Version,
		Salt:          v.Salt,
		EncryptedData: encryptedVault,
	}

	fileData, err := json.Marshal(encryptedVaultFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, fileData, 0600)
	if err != nil {
		return err
	}

	return nil
}

func LoadVault(password, filename string) (*Vault, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var encryptedVaultFile EncryptedVault
	err = json.Unmarshal(jsonData, &encryptedVaultFile)
	if err != nil {
		return nil, err
	}

	encryptionKey := crypto.GenerateEncryptionKey(password, encryptedVaultFile.Salt)

	decryptedData, err := crypto.Decrypt(encryptedVaultFile.EncryptedData, encryptionKey)
	if err != nil {
		return nil, err
	}

	var vault Vault
	err = json.Unmarshal(decryptedData, &vault)
	if err != nil {
		return nil, err
	}

	return &vault, nil
}
