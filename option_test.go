package loghook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	t.Run("NewOption", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			option := NewOption("discord")
			assert.Equal(t, "discord", option.Types())
		})
	})
}
