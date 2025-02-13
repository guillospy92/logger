package logger

import (
	"context"
	"log/slog"
)

type ctxKey string

const (
	slogFields ctxKey = "slog_fields"
)

// ContextHandler implements interface slog.SaveHandler
type ContextHandler struct {
	slog.Handler
}

// Handle customize the default andler to make attributes traces
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}

	return h.Handler.Handle(ctx, r)
}

// AppendCtx creates a context for the attributes that are propagated
func AppendCtx(parent context.Context, attrs ...slog.Attr) context.Context {
	ctx := parent
	if ctx == nil {
		ctx = context.Background()
	}

	if v, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attrs...)
		return context.WithValue(ctx, slogFields, v)
	}

	var v []slog.Attr
	v = append(v, attrs...)

	return context.WithValue(parent, slogFields, v)
}
