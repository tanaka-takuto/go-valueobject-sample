package main

// ValueObject is a generic type that holds a value.
type ValueObject[T any] struct {
	value T
}

// NewValueObject returns a new ValueObject.
func NewValueObject[T any](value T) ValueObject[T] {
	return ValueObject[T]{value}
}

// StringValueObject is a type that holds a string value.
type StringValueObject struct {
	ValueObject[string]
}

// NewStringValueObject returns a new StringValueObject.
func NewStringValueObject(value string) StringValueObject {
	return StringValueObject{NewValueObject(value)}
}

// BytesValueObject is a type that holds an encrypted byte slice.
type BytesValueObject struct {
	ValueObject[[]byte]
}

// NewBytesValueObject returns a new BetesValueObject.
func NewBytesValueObject(value []byte) BytesValueObject {
	return BytesValueObject{NewValueObject(value)}
}
