package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeTrack(t *testing.T) {
	totalTime := TimeTrack(time.Now(), "MockTestFunc")

	require.Contains(t, totalTime, "MockTestFunc took")
}
