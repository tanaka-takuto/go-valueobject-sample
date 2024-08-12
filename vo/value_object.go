package vo

// valueObject is a generic type that holds a value.
type valueObject[T any] struct {
	value T
}

// Value returns the value.
func (vo valueObject[T]) Value() T {
	return vo.value
}

// NewValueObject returns a new ValueObject.
func NewValueObject[T any](value T) valueObject[T] {
	return valueObject[T]{value}
}

// StringValueObject is a type that holds a string value.
type StringValueObject struct {
	valueObject[string]
}

// NewStringValueObject returns a new StringValueObject.
func NewStringValueObject(value string) StringValueObject {
	return StringValueObject{NewValueObject(value)}
}

// BytesValueObject is a type that holds an encrypted byte slice.
type BytesValueObject struct {
	valueObject[[]byte]
}

// NewBytesValueObject returns a new BetesValueObject.
func NewBytesValueObject(value []byte) BytesValueObject {
	return BytesValueObject{NewValueObject(value)}
}

// IntValueObject is a type that holds an integer value.
type IntValueObject struct {
	valueObject[int]
}

// NewIntValueObject returns a new IntValueObject.
func NewIntValueObject(value int) IntValueObject {
	return IntValueObject{NewValueObject(value)}
}
