package main

import (
	"errors"
	"github.com/zaketn/GuestsAPI/pkg/validation"
	"sort"
	"strings"
)

type cp struct {
	CountryCode, PhoneCode string
}

func matchCountryFromPhone(phone string) (string, error) {
	countryPhonesJson, err := validation.ReadCountryWithPhones()
	if err != nil {
		return "", errors.New("failed to get the file with phone codes")
	}

	var countryPhones []cp
	for country, phoneCode := range countryPhonesJson {
		countryPhones = append(countryPhones, cp{country, phoneCode})
	}

	sort.SliceStable(countryPhones, func(i, j int) bool {
		return len(countryPhones[i].PhoneCode) > len(countryPhones[j].PhoneCode)
	})

	for _, cp := range countryPhones {
		if strings.HasPrefix(phone, cp.PhoneCode) {
			return cp.CountryCode, nil
		}
	}

	return "", nil
}
