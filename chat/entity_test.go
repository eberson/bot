package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityFullFilled(t *testing.T) {

	t.Run("without values", func(t *testing.T) {
		a := assert.New(t)

		e := Entity{
			Name: "test",
		}

		a.Equal(true, e.Fulfilled(map[string]string{
			"test":  "xpto",
			"other": "abc",
		}))

		a.Equal(false, e.Fulfilled(map[string]string{
			"other": "abc",
		}))
	})

	t.Run("with default value", func(t *testing.T) {
		a := assert.New(t)

		e := Entity{
			Name:         "test",
			DefaultValue: "a",
		}

		a.Equal(true, e.Fulfilled(map[string]string{
			"other": "abc",
		}))
	})

	t.Run("with values", func(t *testing.T) {
		a := assert.New(t)

		e := Entity{
			Name: "test",
			Values: []string{
				"a",
				"b",
				"c",
			},
		}

		a.Equal(true, e.Fulfilled(map[string]string{
			"test":  "b",
			"other": "abc",
		}))

		a.Equal(false, e.Fulfilled(map[string]string{
			"test":  "d",
			"other": "abc",
		}))

		a.Equal(false, e.Fulfilled(map[string]string{
			"other": "abc",
		}))
	})
}

func TestEntityValueInto(t *testing.T) {
	params := Parameters{
		"a": "test",
		"b": "1",
		"c": "1.09",
		"d": "true",
	}

	a := assert.New(t)

	t.Run("string success", func(t *testing.T) {
		e := Entity{
			Name: "a",
		}

		var r string

		err := e.ValueInto(params, &r)

		a.Nil(err)
		a.Equal(params["a"], r)
	})

	t.Run("int success", func(t *testing.T) {
		e := Entity{
			Name: "b",
		}

		var r int

		err := e.ValueInto(params, &r)

		a.Nil(err)
		a.Equal(1, r)
	})

	t.Run("float success", func(t *testing.T) {
		e := Entity{
			Name: "c",
		}

		var r float64

		err := e.ValueInto(params, &r)

		a.Nil(err)
		a.Equal(1.09, r)
	})

	t.Run("bool success", func(t *testing.T) {
		e := Entity{
			Name: "d",
		}

		var r bool

		err := e.ValueInto(params, &r)

		a.Nil(err)
		a.Equal(true, r)
	})

	t.Run("fail not exists", func(t *testing.T) {
		e := Entity{
			Name: "x",
		}

		var r string

		err := e.ValueInto(params, &r)

		a.NotNil(err)
		a.Equal(errParamNotFulFilled, err)
	})

	t.Run("success has default value", func(t *testing.T) {
		e := Entity{
			Name:         "x",
			DefaultValue: "123",
		}

		var r string

		err := e.ValueInto(params, &r)

		a.Nil(err)
		a.Equal(e.DefaultValue, r)
	})

	t.Run("int fail", func(t *testing.T) {
		e := Entity{
			Name: "a",
		}

		var r int

		err := e.ValueInto(params, &r)

		a.NotNil(err)
	})

	t.Run("float fail", func(t *testing.T) {
		e := Entity{
			Name: "a",
		}

		var r float64

		err := e.ValueInto(params, &r)

		a.NotNil(err)
	})

	t.Run("bool fail", func(t *testing.T) {
		e := Entity{
			Name: "a",
		}

		var r bool

		err := e.ValueInto(params, &r)

		a.NotNil(err)
	})
}
