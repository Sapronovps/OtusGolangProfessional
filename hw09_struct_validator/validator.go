package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const separatorRule = "|"

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var resultErr string
	for _, item := range v {
		resultErr += fmt.Sprintf("field: %s | err: %s\n", item.Field, item.Err)
	}

	return resultErr
}

func Validate(v interface{}) error {
	var errSlice ValidationErrors

	typeStruct := reflect.TypeOf(v)
	valueStruct := reflect.ValueOf(v)

	if typeStruct.Kind().String() != "struct" {
		errSlice = append(errSlice, ValidationError{"none", fmt.Errorf("value is not struct")})
		return errSlice
	}

	// Итерируемся по полям структуры
	for i := 0; i < typeStruct.NumField(); i++ {
		typeField := typeStruct.Field(i)
		valueField := valueStruct.Field(i)

		// Получаем теги поля
		validateTag := typeField.Tag.Get("validate")

		// Пропускам валидацию, если нет тега validate
		if validateTag == "" {
			continue
		}
		// Получаем правила валидации для текущего поля
		rules := getRules(validateTag)
		err := validateField(valueField, typeField, rules)
		if err != nil {
			errSlice = append(errSlice, err...)
		}
	}

	return errSlice
}

func validateField(valueField reflect.Value, typeField reflect.StructField, rules map[string]string) ValidationErrors {
	var errSlice ValidationErrors

	switch valueField.Kind().String() {
	case "int":
		err := validateInt(valueField.Int(), typeField.Name, rules)
		if err != nil {
			errSlice = append(errSlice, err...)
		}
	case "string":
		err := validateString(valueField.String(), typeField.Name, rules)
		if err != nil {
			errSlice = append(errSlice, err...)
		}
	case "slice":
		switch valueField.Type().String() {
		case "[]int":
			for _, v := range valueField.Interface().([]int) {
				err := validateInt(int64(v), typeField.Name, rules)
				if err != nil {
					errSlice = append(errSlice, err...)
				}
			}
		case "[]string":
			for _, v := range valueField.Interface().([]string) {
				err := validateString(v, typeField.Name, rules)
				if err != nil {
					errSlice = append(errSlice, err...)
				}
			}
		default:
			errSlice = append(errSlice, ValidationError{
				Field: typeField.Name,
				Err:   fmt.Errorf("unsupported type field"),
			})
		}
	default:
		errSlice = append(errSlice, ValidationError{
			Field: typeField.Name,
			Err:   fmt.Errorf("unsupported type field"),
		})
	}

	return errSlice
}

func validateInt(value int64, field string, rules map[string]string) ValidationErrors {
	var errSlice ValidationErrors

	for ruleName, ruleValue := range rules {
		validationErr := ValidationError{Field: field}
		switch ruleName {
		case "min":
			if err := validateMinMaxInt(value, ruleValue, true); err != nil {
				validationErr.Err = err
				errSlice = append(errSlice, validationErr)
			}
		case "max":
			if err := validateMinMaxInt(value, ruleValue, false); err != nil {
				validationErr.Err = err
				errSlice = append(errSlice, validationErr)
			}
		case "in":
			isNumberInList := false
			for _, item := range strings.Split(ruleValue, ",") {
				numStr := strings.TrimSpace(item)
				num, err := strconv.Atoi(numStr)
				if err != nil {
					validationErr.Err = err
					errSlice = append(errSlice, validationErr)
					continue
				}
				if value == int64(num) {
					isNumberInList = true
					break
				}
			}

			if !isNumberInList {
				validationErr.Err = fmt.Errorf("in error number %d not contained in %s", value, ruleValue)
				errSlice = append(errSlice, validationErr)
			}

		default:
			validationErr.Err = fmt.Errorf("unsupported rule: %s:%s", ruleName, ruleValue)
			errSlice = append(errSlice, validationErr)
		}
	}

	return errSlice
}

func validateMinMaxInt(value int64, valueRule string, isMin bool) error {
	minMaxValue, err := strconv.Atoi(valueRule)
	if err != nil {
		return fmt.Errorf("atoi error - %w", err)
	}

	if isMin && value < int64(minMaxValue) {
		return fmt.Errorf("value must be no less then %d, current: %d", minMaxValue, value)
	}
	if !isMin && value > int64(minMaxValue) {
		return fmt.Errorf("value must be no greater then %d, current: %d", minMaxValue, value)
	}
	return nil
}

func validateString(value string, field string, rules map[string]string) ValidationErrors {
	var errSlice ValidationErrors

	for ruleName, ruleValue := range rules {
		validationErr := ValidationError{Field: field}
		switch ruleName {
		case "len":
			maxLength, err := strconv.Atoi(ruleValue)
			if err != nil {
				validationErr.Err = fmt.Errorf("atoi error - %w", err)
				errSlice = append(errSlice, validationErr)
				continue
			}
			if len(value) != maxLength {
				validationErr.Err = fmt.Errorf("len must be equal to %d, current: %d", maxLength, len(value))
				errSlice = append(errSlice, validationErr)
			}
		case "regexp":
			matchString, err := regexp.MatchString(ruleValue, value)
			if err != nil {
				validationErr.Err = fmt.Errorf("regexp error - %w", err)
				errSlice = append(errSlice, validationErr)
				continue
			}
			if !matchString {
				validationErr.Err = fmt.Errorf("regexp mismatch error - %s, current: %s", ruleValue, value)
				errSlice = append(errSlice, validationErr)
			}
		case "in":
			isContains := false
			for _, item := range strings.Split(ruleValue, ",") {
				if strings.Contains(value, item) {
					isContains = true
				}
			}
			if !isContains {
				validationErr.Err = fmt.Errorf("in contains error - %s, current: %s", ruleValue, value)
				errSlice = append(errSlice, validationErr)
			}
		default:
			validationErr.Err = fmt.Errorf("unsupported rule: %s:%s", ruleName, ruleValue)
			errSlice = append(errSlice, validationErr)
		}
	}

	return errSlice
}

func getRules(validateTag string) map[string]string {
	rules := make(map[string]string)
	ruleParts := strings.Split(validateTag, separatorRule)

	for _, part := range ruleParts {
		// Разделяем имя правила и значение
		kv := strings.SplitN(part, ":", 2)
		if len(kv) == 2 {
			rules[kv[0]] = kv[1]
		}
	}

	return rules
}
