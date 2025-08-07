package lex

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func strp(s string) *string {
	return &s
}

func TestDelegate(t *testing.T) {
	slots := events.Slots{
		"baz": strp("qux"),
	}
	evt := &Event{
		LexEvent: &events.LexEvent{
			SessionAttributes: map[string]string{
				"foo": "bar",
			},
			CurrentIntent: &events.LexCurrentIntent{
				Slots: slots,
			},
		},
	}

	expected := &Response{
		DialogAction: events.LexDialogAction{
			Type:  "Delegate",
			Slots: slots,
		},
		SessionAttributes: evt.SessionAttributes,
	}

	if !reflect.DeepEqual(evt.Delegate(), expected) {
		t.Error("failed")
	}
}

func TestClose(t *testing.T) {
	evt := &Event{
		LexEvent: &events.LexEvent{
			SessionAttributes: map[string]string{
				"foo": "bar",
			},
		},
	}

	expected := &Response{
		DialogAction: events.LexDialogAction{
			Type:             "Close",
			FulfillmentState: "Fulfilled",
			Message: map[string]string{
				"content":     "test",
				"contentType": "PlainText",
			},
		},
		SessionAttributes: evt.SessionAttributes,
	}

	if !reflect.DeepEqual(evt.Close("test"), expected) {
		t.Error("failed")
	}
}

func TestConfirmIntent(t *testing.T) {
	slots := events.Slots{
		"baz": strp("qux"),
	}
	evt := &Event{
		LexEvent: &events.LexEvent{
			SessionAttributes: map[string]string{
				"foo": "bar",
			},
			CurrentIntent: &events.LexCurrentIntent{
				Name:  "intent",
				Slots: slots,
			},
		},
	}

	expected := &Response{
		DialogAction: events.LexDialogAction{
			Type:       "ConfirmIntent",
			IntentName: "intent",
			Slots:      slots,
			Message: map[string]string{
				"content":     "test",
				"contentType": "PlainText",
			},
		},
		SessionAttributes: evt.SessionAttributes,
	}

	if !reflect.DeepEqual(evt.ConfirmIntent("test"), expected) {
		t.Error("failed")
	}
}

func TestElicitSlot(t *testing.T) {
	slots := events.Slots{
		"baz": strp("qux"),
	}
	evt := &Event{
		LexEvent: &events.LexEvent{
			SessionAttributes: map[string]string{
				"foo": "bar",
			},
			CurrentIntent: &events.LexCurrentIntent{
				Name:  "intent",
				Slots: slots,
			},
		},
	}

	expected := &Response{
		DialogAction: events.LexDialogAction{
			Type:         "ElicitSlot",
			IntentName:   "intent",
			Slots:        slots,
			SlotToElicit: "slot",
			Message: map[string]string{
				"content":     "test",
				"contentType": "PlainText",
			},
		},
		SessionAttributes: evt.SessionAttributes,
	}

	if !reflect.DeepEqual(evt.ElicitSlot("slot", "test"), expected) {
		t.Error("failed")
	}
}

func TestClearSlot(t *testing.T) {
	evt := &Event{
		LexEvent: &events.LexEvent{
			CurrentIntent: &events.LexCurrentIntent{
				Slots: events.Slots{
					"baz": strp("qux"),
				},
			},
		},
	}

	evt.ClearSlot("baz")

	if evt.CurrentIntent.Slots["baz"] != nil {
		t.Error("failed")
	}
}
