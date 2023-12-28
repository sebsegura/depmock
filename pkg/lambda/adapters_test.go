package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdaptEvent(t *testing.T) {
	type Body struct {
		Content string
	}

	var (
		body = Body{
			Content: "iao",
		}
		tests = []struct {
			name    string
			give    json.RawMessage
			inType  any
			want    *Body
			wantErr assert.ErrorAssertionFunc
		}{
			{
				name:    "should adapt an API GW event",
				give:    prepareEvent(_apiEvent, body),
				want:    &body,
				wantErr: assert.NoError,
			},
			{
				name:    "should adapt a SQS event",
				give:    prepareEvent(_sqsEvent, body),
				want:    &body,
				wantErr: assert.NoError,
			},
			{
				name:    "should adapt a SNS event",
				give:    prepareEvent(_snsEvent, body),
				want:    &body,
				wantErr: assert.NoError,
			},
			{
				name:    "should adapt a Cloudwatch event",
				give:    prepareEvent(_cwEvent, body),
				want:    &body,
				wantErr: assert.NoError,
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := AdaptEvent[Body](tt.give)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, res)
		})
	}
}

func prepareEvent(evtType string, evt any) json.RawMessage {
	var raw json.RawMessage

	switch evtType {
	case _apiEvent:
		raw, _ = json.Marshal(evt)
	case _sqsEvent:
		body, _ := json.Marshal(evt)
		sqsEvt := events.SQSEvent{
			Records: []events.SQSMessage{
				{
					EventSource: _sqsEvent,
					Body:        string(body),
				},
			},
		}
		raw, _ = json.Marshal(sqsEvt)
	case _snsEvent:
		body, _ := json.Marshal(evt)
		snsEvt := events.SNSEvent{
			Records: []events.SNSEventRecord{
				{
					EventSource: _snsEvent,
					SNS:         events.SNSEntity{Message: string(body)},
				},
			},
		}
		raw, _ = json.Marshal(snsEvt)
	case _cwEvent:
		body, _ := json.Marshal(evt)
		cwEvt := events.CloudWatchEvent{
			Source: "source",
			Detail: body,
		}
		raw, _ = json.Marshal(cwEvt)
	}

	return raw
}
