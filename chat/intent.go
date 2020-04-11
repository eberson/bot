package chat

import (
	"regexp"

	"github.com/pkg/errors"
)

func (it *Intent) Validate() error {
	it.regex = make([]*regexp.Regexp, len(it.Expression))

	for i, exp := range it.Expression {
		r, err := regexp.Compile(exp)

		if err != nil {
			return errors.Wrap(err, "error compiling intent expressions")
		}

		it.regex[i] = r
	}

	return nil
}

func (it *Intent) Matches(text string) bool {
	for _, re := range it.regex {
		if re.Match([]byte(text)) {
			return true
		}
	}

	return false
}

func (it *Intent) Parameters(text string) Parameters {
	extractEntities := func(re *regexp.Regexp, text string) Parameters {
		match := re.FindStringSubmatch(text)

		result := make(Parameters)

		if len(match) == 0 {
			return result
		}

		for i, name := range re.SubexpNames() {
			if i != 0 && name != "" {
				if match[i] != "" {
					result[name] = match[i]
				}
			}
		}

		return result
	}

	for _, re := range it.regex {
		if re.Match([]byte(text)) {
			return extractEntities(re, text)
		}
	}

	return nil
}
