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


## How to do it?

Use `reflect` package.

```go
t := reflect.TypeOf(user)
```

The `t` has `t.Field(i)` method that returns the `i`th field. Every field has `field.Tag.Get(tagName)` method that returns parsed tag.

To find out the number of fields in this type, you can use `t.NumField()` method.
