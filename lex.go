package lex

import "github.com/aws/aws-lambda-go/events"

type Event struct {
	*events.LexEvent
}

type Response struct {
	DialogAction      events.LexDialogAction `json:"dialogAction,omitempty"`
	SessionAttributes map[string]string      `json:"sessionAttributes,omitempty"`
}

func (e *Event) Delegate() *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:  "Delegate",
			Slots: e.CurrentIntent.Slots,
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) Close(msg string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:             "Close",
			FulfillmentState: "Fulfilled",
			Message: map[string]string{
				"content":     msg,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) ConfirmIntent(msg string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:       "ConfirmIntent",
			IntentName: e.CurrentIntent.Name,
			Slots:      e.CurrentIntent.Slots,
			Message: map[string]string{
				"content":     msg,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) ElicitSlot(slot, msg string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:         "ElicitSlot",
			IntentName:   e.CurrentIntent.Name,
			Slots:        e.CurrentIntent.Slots,
			SlotToElicit: slot,
			Message: map[string]string{
				"content":     msg,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) ClearSlot(s string) {
	e.CurrentIntent.Slots[s] = nil
}
