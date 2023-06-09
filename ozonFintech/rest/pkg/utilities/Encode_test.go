package utilities_test

import (
	"github.com/stretchr/testify/assert"
	"ozonFintech/pkg/utilities"
	"testing"
)

func TestEncodeBase63(t *testing.T) {
	asrt := assert.New(t)
	testCases := []struct {
		name     string
		input    uint64
		expected string
	}{
		{"zero", 0, "aaaaaaaaaa"},
		{"empty", utilities.HashLink(""), "cGo9u9YJ2a"},
		{"small", utilities.HashLink("test.1u"), "qtZ3u2Q_5M"},
		{"large", utilities.HashLink("test.me/21142/check/check/check/check/c2113111j12rwqoeqo2"),
			"Hj7gy6vGhk"},
		{"big", utilities.HashLink("www.ozon.ru/product/parovaya-shvabra-kitfort-kt-1015-belyy-" +
			"funktsiya-vertikalnoy-parkovki-podacha-para-" +
			"20-30-g-598509052/?avtc=1&avte=1&avts=1686304021&sh=TVUNfEJPNw"), "h4LKbZnJ3R"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utilities.EncodeBase63(tc.input)
			asrt.Equal(tc.expected, result, "Unexpected result for %s", tc.name)
		})
	}
}

func TestHashLink(t *testing.T) {
	asrt := assert.New(t)
	testCases := []struct {
		name     string
		input    string
		expected uint64
	}{
		{name: "test1", input: "my.test/wwwEq", expected: uint64(3977575888100549945)},
		{name: "test2", input: "test.me/21142/check/check/check/check/chek.me11212124121114121122",
			expected: uint64(7378359860320362803)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utilities.HashLink(tc.input)
			asrt.Equal(tc.expected, result, "Unexpected result for HashLink", tc.name)
		})
	}
}
