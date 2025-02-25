package testing

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyTitlecase(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	if got != want {
		t.Errorf("TitleCase(%q, %q) = %q, want %q", str, minor, got, want)
	}
}

func TestRussianTitlecase(t *testing.T) {
	const str, minor, want = "привет друг", "", "Привет Друг"
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
	require.Equal(t, want, got)
}

func TestCapsTitlecase(t *testing.T) {
	const str, minor, want = "ПРИВЕТ ДРУГ", "", "Привет Друг"
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
}

func TestTitlecase(t *testing.T) {
	testCases := []struct {
		str   string
		minor string
		want  string
	}{
		{"hello friend", "", "Hello Friend"},
		{"hello friend", "friend", "Hello friend"},
		{"", "", ""},
		{"HI HELLO", "", "Hi Hello"},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			got := titlecase.TitleCase(tc.str, tc.minor)
			assert.Equal(t, tc.want, got)
		})
	}
}
