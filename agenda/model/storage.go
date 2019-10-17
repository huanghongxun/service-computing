package model

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Path string
}

func (s *Storage) Load(v interface{}) error {
	file, err := os.Open(s.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(v)
}

func (s *Storage) Save(v interface{}) error {
	file, err := os.Create(s.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(v)
}
