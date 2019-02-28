package lex

import "github.com/aws/aws-lambda-go/events"

type Event struct {
	events.LexEvent
}

type Response struct {
	DialogAction      events.LexDialogAction `json:"dialogAction,omitempty"`
	SessionAttributes map[string]string      `json:"sessionAttributes,omitempty"`
}

func (e *Event) Delegate() *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:  "Delegate",
			Slots: e.DialogAction.Slots,
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) Close(m string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:             "Close",
			FulfillmentState: "Fulfilled",
			Message: map[string]string{
				"content":     m,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) ConfirmIntent(m string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:       "ConfirmIntent",
			IntentName: e.DialogAction.IntentName,
			Slots:      e.DialogAction.Slots,
			Message: map[string]string{
				"content":     m,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: e.SessionAttributes,
	}
}

func (e *Event) ElicitSlot(m string) *Response {
	return &Response{
		DialogAction: events.LexDialogAction{
			Type:         "ElicitSlot",
			IntentName:   e.DialogAction.IntentName,
			Slots:        e.DialogAction.Slots,
			SlotToElicit: e.DialogAction.SlotToElicit,
			Message: map[string]string{
				"content":     m,
				"contentType": "PlainText",
			},
		},
		SessionAttributes: event.SessionAttributes,
	}
}

func (e *Event) ClearSlot(s string) {
	e.CurrentIntent.Slots[s] = nil
}
