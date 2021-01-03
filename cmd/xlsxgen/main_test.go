package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetRandString(t *testing.T) {
	expected := "abc"
	actual := getRandString()
	assert.IsType(t, expected, actual)
}

func TestGetRandInt(t *testing.T) {
	min, max := 0, 100
	expected := 1
	actual := getRandInt()
	assert.IsType(t, expected, actual)
	assert.GreaterOrEqual(t, actual, min)
	assert.LessOrEqual(t, actual, max)
}

func TestGetRandBool(t *testing.T) {
	expected := true
	actual := getRandBool()
	assert.IsType(t, expected, actual)
}
