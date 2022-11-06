package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMongoClient(t *testing.T) {
	_, err := GetMongoClient()
	require.NoError(t, err)
}
