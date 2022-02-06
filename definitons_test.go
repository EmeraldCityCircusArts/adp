package adp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssociateOIDType_v02_String(t *testing.T) {
	assert.Equal(t, AssociateOIDType_v02("G15TFFD8R0YMSAA7").String(), "G15TFFD8R0YMSAA7")
}
