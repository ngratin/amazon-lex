# Amazon Lex

[![GoDoc](https://godoc.org/github.com/seqr/lex?status.svg)](https://godoc.org/github.com/seqr/lex)

Amazon Lex client for Lambda functions with Go

## Installation
```
go get github.com/seqr/lex
```

## Usage
```go
package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/seqr/lex"
)

func handler(ctx context.Context, lexEvent events.LexEvent) (*lex.Response, error) {
	evt := lex.Event{LexEvent: &lexEvent}

	switch evt.CurrentIntent.Name {
	case "Hello":
		return evt.Close("Hello, world!"), nil
	default:
		return evt.Delegate(), nil
	}
}

func main() {
	lambda.Start(handler)
}
```

## API
- `(e *Event) Delegate() *Response`
- `(e *Event) Close(msg string) *Response`
- `(e *Event) ConfirmIntent(msg string) *Response`
- `(e *Event) ElicitSlot(slot, msg string) *Response`
- `(e *Event) ClearSlot(s string)`
