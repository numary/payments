package task

import (
	"context"
)

type StateResolver interface {
	ResolveTo(ctx context.Context, v any) error
}
type StateResolverFn func(ctx context.Context, v any) error

func (fn StateResolverFn) ResolveTo(ctx context.Context, v any) error {
	return fn(ctx, v)
}

func ResolveTo[State any](ctx context.Context, resolver StateResolver, to *State) (*State, error) {
	err := resolver.ResolveTo(ctx, to)
	if err != nil {
		return nil, err
	}
	return to, nil
}

func MustResolveTo[State any](ctx context.Context, resolver StateResolver, to State) State {
	state, err := ResolveTo[State](ctx, resolver, &to)
	if err == ErrNotFound {
		return to
	}
	if err != nil {
		panic(err)
	}
	return *state
}
