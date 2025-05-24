package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
	assert.NotEqual(t, AddOne(2), 4, "AddOne(2) should not be 4")
	assert.Nil(t, nil, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not execute")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("executing")
}
