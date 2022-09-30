package realtime_test

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"scrumlr.io/server/realtime"
)

func TestRealtime_GetBoardSessionRequestChannel(t *testing.T) {

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	rt := realtime.New(SetupNatsContainer(t))
	testBoard := uuid.New()
	testUser := uuid.New()

	testEvents := []realtime.BoardSessionRequestEventType{
		realtime.RequestRejected,
		realtime.RequestAccepted,
		"some undefined event",
	}

	eventChannel := rt.GetBoardSessionRequestChannel(testBoard, testUser)
	readEvents := []realtime.BoardSessionRequestEventType{}
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			case ev := <-eventChannel:
				assert.NotNil(t, ev)
				readEvents = append(readEvents, *ev)
				wg.Done()
			case <-ctx.Done():
				return
			}
		}
	}()

	for _, ev := range testEvents {
		err := rt.BroadcastUpdateOnBoardSessionRequest(testBoard, testUser, ev)
		assert.Nil(t, err)
		wg.Add(1)
	}

	wg.Wait()
	assert.Equal(t, testEvents, readEvents)

}
