package validation

import (
	"errors"
	"fmt"
	"net/http"
)

type FormValidator struct {
	Request *http.Request
}

type Ruleset struct {
	Rules *map[string][]Rule
}

type Rule func(fieldValue string) error

func (fv FormValidator) Validate(rs *Ruleset) error {
	err := fv.Request.ParseForm()
	if err != nil {
		return err
	}

	for field, rules := range *rs.Rules {
		fieldValue := fv.Request.PostForm.Get(field)

		if fv.Request.Method == "GET" || fv.Request.Method == "DELETE" {
			fieldValue = fv.Request.Form.Get(field)
		}

		for _, rule := range rules {
			vErr := rule(fieldValue)
			if vErr != nil {
				return errors.New(fmt.Sprintf("%s: %s", field, vErr.Error()))
			}
		}
	}

	return nil
}
