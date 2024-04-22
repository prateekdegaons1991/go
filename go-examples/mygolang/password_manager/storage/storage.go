package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PasswordEntry struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func SavePasswordsToFile(entries []PasswordEntry, filename string) error {
	data, err := json.MarshalIndent(entries, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
func LoadPasswordsFromFile(filename string) ([]PasswordEntry, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []PasswordEntry{}, nil
		}
		return nil, err
	}
	var entries []PasswordEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}
