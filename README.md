# goeql Package

## Overview

The `goeql` package provides a set of helper functions to assist with serializing and deserializing various data types into a format that is compatible with CipherStash Proxy. This package enables seamless encryption and decryption of values and facilitates search operations on encrypted data without exposing the plaintext.

The package is designed to support [CipherStash’s Encrypt Query Language (EQL)](https://github.com/cipherstash/encrypt-query-language), which is a language for querying encrypted data in a PostgreSQL database.

## Installation

To install `goeql`, use `go get`:

```bash
go get github.com/cipherstash/goeql
```

## Data Format

EQL requires data to be serialized in the following JSON format:

```json
{
  "k": "pt",
  "p": "a string representation of the plaintext that is being encrypted",
  "i": {
    "t": "table",
    "c": "column"
  },
  "v": 1
}
```

For more information about this format, refer to the [Encrypt Query Language documentation](https://github.com/cipherstash/encrypt-query-language#data-format).

## Supported Types

The `goeql` package supports the following data types for serialization and deserialization:

- `EncryptedText`: Represents a `string` value.
- `EncryptedJsonb`: Represents a `jsonb` object (map).
- `EncryptedInt`: Represents an `int` value.
- `EncryptedBool`: Represents a `bool` value.

## Usage

### Serialization

Each supported type provides a `Serialize` method that converts the value into an EQL-compatible JSON format for CipherStash Proxy:

```go
text := EncryptedText("secret value")
data, err := text.Serialize("users", "password")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))  // JSON-encoded EQL object
```

### Deserialization

The `Deserialize` method allows converting a JSON payload from CipherStash Proxy back into the corresponding type:

```go
var text EncryptedText
err := text.Deserialize(data)
if err != nil {
    log.Fatal(err)
}
fmt.Println(text)  // Decrypted plaintext value
```

### Query Serialization

The package provides helper functions to serialize queries that interact with encrypted data in various ways:

- `MatchQuery`: Serializes a plaintext value for an equality query.
- `OreQuery`: Serializes a value for order-preserving encryption queries (range).
- `UniqueQuery`: Serializes a value for a unique constraint check.
- `JsonbQuery`: Serializes a value for JSONB vector-based queries.

Example:

```go
queryData, err := MatchQuery("search term", "users", "username")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(queryData))
```

## Functions

### `Serialize()`

- **Description**: Converts a plaintext value into an encrypted JSON payload that conforms to the EQL format.
- **Parameters**: 
  - `table`: The name of the table.
  - `column`: The name of the column.
- **Returns**: `[]byte` (serialized JSON), `error`

### `Deserialize()`

- **Description**: Converts an encrypted JSON payload back into its plaintext representation.
- **Parameters**: 
  - `data`: JSON payload from CipherStash Proxy.
- **Returns**: Decrypted value, `error`

### `MatchQuery()`

- **Description**: Serializes a value for use in an equality query in EQL.
- **Parameters**: 
  - `value`: The plaintext value to query.
  - `table`: The name of the table.
  - `column`: The name of the column.
- **Returns**: Serialized query, `error`

### `OreQuery()`, `UniqueQuery()`, `JsonbQuery()`

These functions work similarly to `MatchQuery()`, but are used for different query types, such as range, unique, and JSONB queries.

## Example

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // Encrypt a text value
    text := EncryptedText("example plaintext")
    data, err := text.Serialize("users", "email")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Serialized data:", string(data))
    
    // Decrypt the value
    var decryptedText EncryptedText
    decryptedText, err = decryptedText.Deserialize(data)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Decrypted text:", decryptedText)
}
```

## Contributing

We welcome contributions! Feel free to open an issue or submit a pull request if you find a bug or have suggestions for improvement.

## License

This project is licensed under the MIT License.