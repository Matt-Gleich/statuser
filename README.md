# statuser

📣 A user-friendly status outputting library for go

![test](https://github.com/Matt-Gleich/statuser/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/statuser/workflows/lint/badge.svg)

## 🚀 Installing

Simply run the command below in the root of your go project:

```bash
go get -u github.com/Matt-Gleich/statuser
```

## 📚 Documentation

### Emojis

By default, emojis are turned on. If you wish to disable them then just do `emojis = false` in your go code.

### `Error()`

Output an error to the user

#### Parameters

1. message
   - Type: `string`
   - Description: A human readable message to help the user
2. err
   - Type: `err`
   - Description: The golang error
3. exitCode
   - Type: `int`
   - Description: Exit code used to exit the program

#### Example

Outputs the following in red text:

```
░░░░░░░░░░░░░
░🚨 ERROR 🚨░
░░░░░░░░░░░░░

message

GOLANG ERROR (SHOW DEVELOPER):
err
exit status 1
```

### `ErrorMsg()`

Output an error to the user with just a message

#### Parameters

1. message
   - Type: `string`
   - Description: A human readable message to help the user
2. exitCode
   - Type: `int`
   - Description: Exit code used to exit the program

#### Example

Outputs the following in red text:

```
░░░░░░░░░░░░░
░🚨 ERROR 🚨░
░░░░░░░░░░░░░

message
exit status 1
```

### `Warning()`

Output a warning the user

#### Parameters

1. message
   - Type: `string`
   - Description: The warning message to display to the user

#### Example

Outputs the following text in yellow:

```
⚠️ WARNING ⚠️
message
```

### `Success()`

Output a warning the user

#### Parameters

1. message
   - Type: `string`
   - Description: The success message to display to the user

#### Example

Outputs the following text in green:

```
✓ message
```

## Contributors

1. Matthew Gleich ([@Matt-Gleich](http://www.github.com/Matt-Gleich))
