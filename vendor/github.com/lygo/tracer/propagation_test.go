package tracer

import (
	"bytes"
	"testing"

	"github.com/opentracing/opentracing-go"
)

func TestText(t *testing.T) {
	sp := &Span{
		raw: RawSpan{
			SpanContext: SpanContext{
				SpanID:   1,
				ParentID: 2,
				TraceID:  3,
				Flags:    FlagSampled,
				Baggage: map[string]string{
					"k1": "v1",
					"k2": "",
				},
			},
		},
	}

	carrier := opentracing.TextMapCarrier{}
	if err := textInjecter(sp.Context().(SpanContext), carrier); err != nil {
		t.Fatal("unexpected error: ", err)
	}
	context, err := textExtracter(carrier)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	if context.TraceID != sp.raw.TraceID ||
		context.ParentID != sp.raw.ParentID ||
		context.SpanID != sp.raw.SpanID ||
		context.Flags != sp.raw.Flags ||
		len(context.Baggage) != 2 ||
		context.Baggage["k1"] != "v1" ||
		context.Baggage["k2"] != "" {

		t.Errorf("got (%d, %d, %d, %d, %v), want (%d, %d, %d, %d, %v)",
			context.TraceID, context.ParentID, context.SpanID, context.Flags, context.Baggage,
			sp.raw.TraceID, sp.raw.ParentID, sp.raw.SpanID, sp.raw.Flags, sp.raw.Baggage)
	}
}

func TestBinary(t *testing.T) {
	sp := &Span{
		raw: RawSpan{
			SpanContext: SpanContext{
				SpanID:   1,
				ParentID: 2,
				TraceID:  3,
				Flags:    FlagSampled,
				Baggage: map[string]string{
					"k1": "v1",
					"k2": "",
				},
			},
		},
	}
	buf := &bytes.Buffer{}
	if err := binaryInjecter(sp.Context().(SpanContext), buf); err != nil {
		t.Fatal("unexpected error: ", err)
	}
	context, err := binaryExtracter(buf)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	if context.TraceID != sp.raw.TraceID ||
		context.ParentID != sp.raw.ParentID ||
		context.SpanID != sp.raw.SpanID ||
		context.Flags != sp.raw.Flags ||
		len(context.Baggage) != 2 ||
		context.Baggage["k1"] != "v1" ||
		context.Baggage["k2"] != "" {

		t.Errorf("got (%d, %d, %d, %d, %v), want (%d, %d, %d, %d, %v)",
			context.TraceID, context.ParentID, context.SpanID, context.Flags, context.Baggage,
			sp.raw.TraceID, sp.raw.ParentID, sp.raw.SpanID, sp.raw.Flags, sp.raw.Baggage)
	}
}
