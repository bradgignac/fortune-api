# fortune-api

A HTTP API for generating fortunes, inspired by [fortune](http://en.wikipedia.org/wiki/Fortune_%28Unix%29).

## API

### `GET /fortunes`

Return a list of all fortunes in the fortune database.

**Sample Request:**

```json
GET /fortunes
```

**Sample Response:**

```json
Status: 200 OK

[
  {
    "id": "8843d7f",
    "data": "\"Failure is the opportunity to begin again more intelligently.\"\n  ~Henry Ford"
  }
]
```

### `GET /fortunes/:id`

Return a specific fortune from the database.

**Sample Request:**

```json
GET /fortunes/8843d7f
```

**Sample Response:**

```json
Status: 200 OK

{
  "id": "8843d7f",
  "data": "\"Failure is the opportunity to begin again more intelligently.\"\n  ~Henry Ford"
}
```

### `GET /random`

Return a random fortune from the database.

**Sample Request:**

```json
GET /random
```

**Sample Response:**

```json
Status: 302 Found
Location: /fortunes/8843d7f
```
