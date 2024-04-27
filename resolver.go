package yodel

import "context"

type (
	// Resolver resolves a response value for a given request value
	Resolver[Req, Res any] interface {
		Resolve(context.Context, Req) (Res, error)
	}

	// ResolverFunc provides a function implementation of Resolver.
	ResolverFunc[Req, Res any] func(context.Context, Req) (Res, error)

	// StaticResolver resolves a static value.
	StaticResolver[_, T any] struct {
		Value T
	}
)

// Resolve implements Resolver.
func (fn ResolverFunc[Req, Res]) Resolve(ctx context.Context, req Req) (Res, error) {
	return fn(ctx, req)
}

// Resolve implements Resolver.
func (resolver StaticResolver[Req, Res]) Resolve(context.Context, Req) (Res, error) {
	return resolver.Value, nil
}
