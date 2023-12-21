package lambda

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/tidwall/gjson"
)

const (
	_snsEvent = "aws:sns"
	_sqsEvent = "aws:sqs"
	_cwEvent  = "aws:cw"
	_apiEvent = "aws:apigw"
)

var eventKeys = []string{
	"Records.0.eventSource",
	"Records.0.EventSource",
}

func AdaptEvent[I any](raw json.RawMessage) (*I, error) {
	switch eventSource(raw) {
	case _sqsEvent:
		return adaptSQS[I](raw)
	case _snsEvent:
		return adaptSNS[I](raw)
	case _cwEvent:
		return adaptCW[I](raw)
	default:
		return adaptJSON[I](raw)
	}
}

func eventSource(raw json.RawMessage) string {
	jsonStr := string(raw)

	for _, key := range eventKeys {
		if source := gjson.Get(jsonStr, key); source.Exists() {
			return source.String()
		}
	}

	if source := gjson.Get(jsonStr, "source"); source.Exists() {
		return _cwEvent
	}

	return _apiEvent
}

func adaptSQS[I any](raw json.RawMessage) (*I, error) {
	var (
		evt events.SQSEvent
		in  I
	)

	if err := json.Unmarshal(raw, &evt); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(evt.Records[0].Body), &in); err != nil {
		return nil, err
	}

	return &in, nil
}

func adaptSNS[I any](raw json.RawMessage) (*I, error) {
	var (
		evt events.SNSEvent
		in  I
	)

	if err := json.Unmarshal(raw, &evt); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(evt.Records[0].SNS.Message), &in); err != nil {
		return nil, err
	}

	return &in, nil
}

func adaptCW[I any](raw json.RawMessage) (*I, error) {
	var (
		evt events.CloudWatchEvent
		in  I
	)

	if err := json.Unmarshal(raw, &evt); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(evt.Detail), &in); err != nil {
		return nil, err
	}

	return &in, nil
}

func adaptJSON[I any](raw json.RawMessage) (*I, error) {
	var in I
	if err := json.Unmarshal(raw, &in); err != nil {
		return nil, err
	}
	return &in, nil
}
