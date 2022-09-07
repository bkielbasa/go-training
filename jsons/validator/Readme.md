# Validator

## Task

Write a validator that will use build tags to perform a simple validation.

```go
type User struct {
    Name string `validate:"required"`
    Email string `validate:"required"`
    FullName string
}
```

For the given struct, it should check the `validate` build tag and return an error if an instance of the struct doesn't meet it's requrements.

The validator's API should look like this:

```go
validator := NewValidator()

err := validator.Validate(user)
```
