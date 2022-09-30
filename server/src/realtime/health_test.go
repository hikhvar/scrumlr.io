package realtime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"scrumlr.io/server/realtime"
)

func TestRealtime_IsHealthy(t *testing.T) {
	//rt := realtime.New("foo")
	// assert.False(t, rt.IsHealthy())

	rt := realtime.New(SetupNatsContainer(t))
	assert.True(t, rt.IsHealthy())
}
