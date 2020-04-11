package chat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntent_Matches(t *testing.T) {
	a := assert.New(t)

	it := Intent{
		Expression: []string{
			"^how many prs are (?P<state>\\w+) in (?P<context>[A-Za-z-_]+)\\??$",
			"^what are (?P<state>\\w+) prs in (?P<context>[A-Za-z-_]+)\\??$",
		},
	}

	err := it.Validate()

	a.Nil(err)

	testCases := []struct {
		text   string
		result bool
	}{
		{
			text:   "how many prs are open in digital-services",
			result: true,
		},
		{
			text:   "what are open prs in digital-services?",
			result: true,
		},
		{
			text:   "how many prs are in digital-services",
			result: false,
		},
		{
			text:   "how many prs are open in digital-services?",
			result: true,
		},
		{
			text:   "how many prs are open",
			result: false,
		},
		{
			text:   "what are open prs in digital-services",
			result: true,
		},
		{
			text:   "what are open prs?",
			result: false,
		},
		{
			text:   "what are prs in digital-services?",
			result: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			r := it.Matches(tc.text)
			a.Equal(tc.result, r)
		})
	}
}

func TestIntent_Parameters(t *testing.T) {
	a := assert.New(t)

	it := Intent{
		Expression: []string{
			"^how many prs are (?P<state>\\w+) in (?P<context>[A-Za-z-_]+)\\??$",
			"^what are (?P<state>\\w+) prs in (?P<context>[A-Za-z-_]+)\\??$",
		},
	}

	err := it.Validate()

	a.Nil(err)

	testCases := []struct {
		text     string
		expected Parameters
	}{
		{
			text: "how many prs are open in digital-services",
			expected: map[string]string{
				"state":   "open",
				"context": "digital-services",
			},
		},
		{
			text: "what are closed prs in digital-services?",
			expected: map[string]string{
				"state":   "closed",
				"context": "digital-services",
			},
		},
		{
			text: "how many prs are open in digital-tiff?",
			expected: map[string]string{
				"state":   "open",
				"context": "digital-tiff",
			},
		},
		{
			text: "what are closed prs in digital-tiff",
			expected: map[string]string{
				"state":   "closed",
				"context": "digital-tiff",
			},
		},
		{
			text:     "what are prs in digital-tiff",
			expected: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			r := it.Parameters(tc.text)
			a.Equal(tc.expected, r)
		})
	}
}
