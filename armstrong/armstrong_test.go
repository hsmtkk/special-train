package armstrong_test

import (
	"testing"

	"github.com/hsmtkk/special-train/armstrong"
	"github.com/stretchr/testify/assert"
)

func TestIsArmstrong(t *testing.T) {
	assert.True(t, armstrong.IsArmstrong(9))
	assert.False(t, armstrong.IsArmstrong(10))
	assert.True(t, armstrong.IsArmstrong(153))
	assert.False(t, armstrong.IsArmstrong(154))
}

func TestNumberToDigit(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, armstrong.NumberToDigit(12345))
}
