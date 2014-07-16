num-id
======

A number to alphanumeric conversion library for Go, similar to URL shortening services such as tinyurl.

This code is a Go port of the code discussed here: http://kvz.io/blog/2009/06/10/create-short-ids-with-php-like-youtube-or-tinyurl/

Credit should be given to Kevin van Zonneveld; I just ported it to Go.

Unit tests included.

Example:

    import (
        "fmt"
        "github.com/simon-whitehead/num-id"
    )

    result := numid.Encode(123456789)
    fmt.Println(result) // "8Kawi"

    original := numid.Decode("8Kawi")
    fmt.Println(original) // 123456789

