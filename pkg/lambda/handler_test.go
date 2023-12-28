package lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testEvent struct {
	Amount string `json:"amount"`
}

func TestAsyncHandler_EventHandler(t *testing.T) {
	fn := func(ctx context.Context, in *testEvent) error {
		if in.Amount != "10" {
			return fmt.Errorf("error")
		}
		return nil
	}

	h := NewAsyncHandler[testEvent](fn)

	assert.NoError(t, h.EventHandler(context.TODO(), makeSQSEvent()))
}

func makeSQSEvent() json.RawMessage {
	evt := testEvent{Amount: "10"}
	body, _ := json.Marshal(evt)
	sqsEvt := events.SQSEvent{
		Records: []events.SQSMessage{
			{
				EventSource: _sqsEvent,
				Body:        string(body),
			},
		},
	}
	raw, _ := json.Marshal(sqsEvt)
	return raw
}
