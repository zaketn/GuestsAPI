package validation

import (
	"encoding/json"
	"log"
	"os"
)

func ReadCountryWithPhones() (map[string]string, error) {
	byt, err := os.ReadFile("./pkg/validation/storage/country_phone.json")
	if err != nil {
		return nil, err
	}

	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Println(err)
		return nil, err
	}

	return dat, nil
}
