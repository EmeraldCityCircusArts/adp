package adp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailUri_String(t *testing.T) {
	assert.Equal(t, EmailUri("test@example.com").String(), "test@example.com")
}
