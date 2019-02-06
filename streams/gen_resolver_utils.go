package streams

import (
	"context"
	"errors"
)

// ErrNoCallbackMatch indicates a Resolver could not match the ActivityStreams value to a callback function.
var ErrNoCallbackMatch error = errors.New("activity stream did not match the callback function")

// ErrUnhandledType indicates that an ActivityStreams value has a type that is not handled by the code that has been generated.
var ErrUnhandledType error = errors.New("activity stream did not match any known types")

// ErrPredicateUnmatched indicates that a predicate is accepting a type or interface that does not match an ActivityStreams value's type or interface.
var ErrPredicateUnmatched error = errors.New("activity stream did not match type demanded by predicate")

// errCannotTypeAssertType indicates that the 'type' property returned by the ActivityStreams value cannot be type-asserted to its interface form.
var errCannotTypeAssertType error = errors.New("activity stream type cannot be asserted to its interface")

// errCannotTypeAssertPredicate indicates that a predicate cannot be type-casted to an expected function signature.
var errCannotTypeAssertPredicate error = errors.New("predicate cannot be type asserted to a known function type")

// ActivityStreamsInterface represents any ActivityStream value code-generated by
// go-fed or compatible with the generated interfaces.
type ActivityStreamsInterface interface {
	// GetName returns the ActiivtyStreams value's type.
	GetName() string
}

// Resolver represents any TypeResolver.
type Resolver interface {
	// Resolve will attempt to resolve an untyped ActivityStreams value into a
	// Go concrete type.
	Resolve(ctx context.Context, o ActivityStreamsInterface) error
}

// IsUnmatchedErr is true when the error indicates that a Resolver was
// unsuccessful due to the ActivityStreams value not matching its callbacks or
// predicates.
func IsUnmatchedErr(err error) bool {
	return err == ErrPredicateUnmatched || err == ErrUnhandledType || err == ErrNoCallbackMatch
}
