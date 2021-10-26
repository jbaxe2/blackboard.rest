package errors

/**
 * The [RestError] interface provides the base type of all REST API typed errors.
 */
type RestError interface {
  Code() string

  Description() string

  error
}
