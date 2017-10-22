package settings

import (
	"encoding/json"
	"os"
)

var Config = &struct{
	Database struct {
		Host     string
		Dbname   string
		User     string
		Password string
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		From     string
	}
	I18n i18n
}{}


type i18n struct {
	ENUS map[string]string `json:"en_us"`
}

func (i *i18n) Get(lang, key string) string {
	switch lang {
	case "en_us":
		return i.ENUS[key]
	}
	return ""
}

func Read(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	if err = json.NewDecoder(file).Decode(&Config); err != nil {
		return err
	}
	return nil
}
