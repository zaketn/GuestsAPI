package validation

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"regexp"
	"strconv"
	"strings"
)

func NotEmpty() Rule {
	return func(fieldValue string) error {
		if fieldValue == "" {
			return errors.New("the value missing or empty")
		}

		return nil
	}
}

func Length(min, max int) Rule {
	return func(fieldValue string) error {
		if len(fieldValue) < min || len(fieldValue) >= max {
			return errors.New(fmt.Sprintf("the value must be between %d and %d characters", min, max))
		}

		return nil
	}
}

func Numeric() Rule {
	return func(fieldValue string) error {
		_, err := strconv.Atoi(fieldValue)
		if err != nil {
			return errors.New("must contains letters characters only")
		}

		return nil
	}
}

func String() Rule {
	return func(fieldValue string) error {
		regex := `^[a-zA-Z\p{L}]+$`
		re := regexp.MustCompile(regex)
		if !re.MatchString(fieldValue) {
			return errors.New("must contains letters characters only")
		}

		return nil
	}
}

func Email() Rule {
	return func(fieldValue string) error {
		regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
		re := regexp.MustCompile(regex)
		if !re.MatchString(fieldValue) {
			return errors.New("email format is invalid")
		}

		return nil
	}
}

func Phone() Rule {
	return func(fieldValue string) error {
		regex := `^\+[\d\(\)\-]+$` // Для лучшей валидности нужно усложнить регулярку
		re := regexp.MustCompile(regex)
		if !re.MatchString(fieldValue) {
			return errors.New("phone number has invalid format")
		}

		return nil
	}
}

func CountryCode() Rule {
	return func(fieldValue string) error {
		countrySrc, err := ReadCountryWithPhones()
		if err != nil {
			return errors.New("failed to get the file with countries")
		}

		_, countryExists := countrySrc[strings.ToUpper(fieldValue)]

		if countryExists {
			return nil
		}

		return errors.New("country should be valid country code e.g US")
	}
}

func Exists(db *pgx.Conn, table, columnName string) Rule {
	return func(fieldValue string) error {
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s=$1)", table, columnName)

		var exists bool
		db.QueryRow(query, fieldValue).Scan(&exists)

		if exists == true {
			return nil
		}

		return errors.New(fmt.Sprintf("the %s value %s was not found", columnName, fieldValue))
	}
}

func DoesNotExist(db *pgx.Conn, table, columnName string) Rule {
	return func(fieldValue string) error {
		exists := Exists(db, table, columnName)(fieldValue)

		if exists == nil {
			return errors.New(fmt.Sprintf("the %s value %s already exists", columnName, fieldValue))
		}

		return nil
	}
}
