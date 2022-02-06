package adp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAssociateOIDType_v02_String(t *testing.T) {
	assert.Equal(t, AssociateOIDType_v02("G15TFFD8R0YMSAA7").String(), "G15TFFD8R0YMSAA7")
}

func TestDateTimeType_v01_String(t *testing.T) {
	assert.Equal(t, DateTimeType_v01("2008-05-11T15:30:00-06:00").String(), "2008-05-11T15:30:00-06:00")
}

func TestDateTimeType_v01_Time(t *testing.T) {
	test1, _ := time.Parse("2006-01-02T15:04:05-07:00", "2008-05-11T15:30:00-06:00")

	assert.Equal(t, DateTimeType_v01("2008-05-11T15:30:00-06:00").Time(), test1)
	assert.Panics(t, func() { DateTimeType_v01("2008-05-11").Time() })
}

func TestDateType_v01_String(t *testing.T) {
	assert.Equal(t, DateType_v01("2008-05-11").String(), "2008-05-11")
}

func TestDateType_v01_Time(t *testing.T) {
	test1, _ := time.Parse("2006-01-02", "2008-05-11")

	assert.Equal(t, DateType_v01("2008-05-11").Time(), test1)
	assert.Panics(t, func() { DateType_v01("2008-05-11T15:30:00-06:00").Time() })
}
