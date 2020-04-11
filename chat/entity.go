package chat

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/eberson/rootinha/helper/strs"

	"github.com/pkg/errors"
)

var (
	errParamNotFulFilled   = errors.New("given parameter is not fulfilled")
	errResultTypeNotMapped = errors.New("result type not mapped")
)

//Fulfilled returns true if given params has this struct's parameter. It will returns
//false if param does not exists or if param does not match possible values (when set)
func (e *Entity) Fulfilled(params Parameters) bool {
	hasDefaultValue := strs.IsNotEmpty(e.DefaultValue)
	param, exists := params[e.Name]

	if !exists {
		return hasDefaultValue
	}

	if len(e.Values) == 0 {
		return true
	}

	for _, value := range e.Values {
		if strings.EqualFold(value, fmt.Sprintf("%s", param)) {
			return true
		}
	}

	return hasDefaultValue
}

func (e *Entity) extract(params Parameters) (string, error) {
	if !e.Fulfilled(params) {
		return "", errParamNotFulFilled
	}

	value, exists := params[e.Name]

	if !exists {
		return e.DefaultValue, nil
	}

	return value, nil
}

func (e *Entity) ValueInto(params Parameters, result interface{}) error {
	v, err := e.extract(params)

	if err != nil {
		return err
	}

	resultValue := reflect.Indirect(reflect.ValueOf(result))

	switch resultValue.Kind() {
	case reflect.String:
		resultValue.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int:
		i, err := strconv.Atoi(v)

		if err != nil {
			return err
		}

		resultValue.Set(reflect.ValueOf(i))
		return nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return err
		}

		resultValue.Set(reflect.ValueOf(f))
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(v)

		if err != nil {
			return err
		}

		resultValue.Set(reflect.ValueOf(b))
		return nil
	}

	return errors.Wrap(errResultTypeNotMapped, resultValue.Kind().String())
}
