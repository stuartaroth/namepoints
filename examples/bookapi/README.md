# example:bookapi

Here are the current HTTP requests and responses to this example API

### web root, empty name, or unsupported name request/response
```shell
curl --location --request POST 'http://localhost:8080' \
--data ''
```

```json
{
    "addBook": {
        "requestSchema": {
            "$schema": "https://json-schema.org/draft/2020-12/schema",
            "$ref": "#/$defs/addBookRequest",
            "$defs": {
                "addBookRequest": {
                    "properties": {
                        "book": {
                            "$ref": "#/$defs/book"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "book"
                    ]
                },
                "book": {
                    "properties": {
                        "title": {
                            "type": "string"
                        },
                        "author": {
                            "type": "string"
                        },
                        "genre": {
                            "type": "string"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "title",
                        "author",
                        "genre"
                    ]
                }
            }
        },
        "responseSchema": {
            "$schema": "https://json-schema.org/draft/2020-12/schema",
            "$ref": "#/$defs/addBookResponse",
            "$defs": {
                "addBookResponse": {
                    "properties": {
                        "book": {
                            "$ref": "#/$defs/book"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "book"
                    ]
                },
                "book": {
                    "properties": {
                        "title": {
                            "type": "string"
                        },
                        "author": {
                            "type": "string"
                        },
                        "genre": {
                            "type": "string"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "title",
                        "author",
                        "genre"
                    ]
                }
            }
        },
        "updatesData": true
    },
    "getAllBooks": {
        "requestSchema": null,
        "responseSchema": {
            "$schema": "https://json-schema.org/draft/2020-12/schema",
            "$ref": "#/$defs/getAllBooksResponse",
            "$defs": {
                "book": {
                    "properties": {
                        "title": {
                            "type": "string"
                        },
                        "author": {
                            "type": "string"
                        },
                        "genre": {
                            "type": "string"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "title",
                        "author",
                        "genre"
                    ]
                },
                "getAllBooksResponse": {
                    "properties": {
                        "books": {
                            "items": {
                                "$ref": "#/$defs/book"
                            },
                            "type": "array"
                        }
                    },
                    "additionalProperties": false,
                    "type": "object",
                    "required": [
                        "books"
                    ]
                }
            }
        },
        "updatesData": false
    }
}
```

### name=getAllBooks

```shell
curl --location --request POST 'http://localhost:8080?name=getAllBooks' \
--data ''
```

```json
{
    "books": [
        {
            "title": "It",
            "author": "Stephen King",
            "genre": "Horror"
        },
        {
            "title": "Ringworld",
            "author": "Larry Niven",
            "genre": "Science Fiction"
        },
        {
            "title": "A Game of Thrones",
            "author": "George R.R. Martin",
            "genre": "Fantasy"
        },
        {
            "title": "Murder on the Orient Express",
            "author": "Agatha Christie",
            "genre": "Mystery"
        },
        {
            "title": "In Cold Blood",
            "author": "Truman Capote",
            "genre": "True Crime"
        }
    ]
}
```

### name=addBook with valid data

```shell
curl --location 'http://localhost:8080?name=addBook' \
--header 'Content-Type: application/json' \
--data '{
    "book": {
        "title": "The Hobbit",
        "author": "J.R.R. Tolkien",
        "genre": "Fantasy"
    }
}'
```

```json
{
    "book": {
        "title": "The Hobbit",
        "author": "J.R.R. Tolkien",
        "genre": "Fantasy"
    }
}
```

### name=addBook with invalid data
```shell
curl --location 'http://localhost:8080?name=addBook' \
--header 'Content-Type: application/json' \
--data '{"book":{}}'
```

```json
{
    "error": "`title` and `author` cannot be empty strings"
}
```
