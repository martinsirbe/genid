# genid — Generate Unique IDs

`genid` is a fast and simple command-line tool written in Go that generates various types of unique 
identifiers. It supports modern, widely used formats including ULID, UUID versions 4–7, KSUID, XID, 
NanoID, and Snowflake.

Perfect for scripting, testing, and inspecting different ID formats in distributed systems.

## Installation

### With Homebrew (Recommended)

If you use [Homebrew](https://brew.sh), you can install `genid` using a custom tap:

```sh
brew tap martinsirbe/clinkclank
brew install martinsirbe/clinkclank/genid
````

### Run Without Installing

If you have Go installed, you can run it directly:

```sh
go run cmd/genid/main.go nanoid 3 --len=10
```

### Build From Source

```sh
git clone https://github.com/martinsirbe/genid.git
cd genid/cmd/genid
go build -o genid
```

Optionally move the binary to your `$PATH`:

```sh
mv genid /usr/local/bin/
```

## Usage

```sh
genid <id_type> [count] [--len=n]
```

### Arguments

| Argument    | Description                                                    |
| ----------- | -------------------------------------------------------------- |
| `<id_type>` | Type of ID to generate (see list below)                        |
| `[count]`   | Optional number of IDs to generate (default: 1)                |
| `[--len=n]` | Optional NanoID length (applies only to `nanoid`, default: 21) |

> Use `genid help` or `genid -h` to print usage instructions.

## Supported ID Types

| Type        | Description                             |
| ----------- | --------------------------------------- |
| `ulid`      | Lexicographically sortable UUID-like ID |
| `uuid4`     | Random UUID (RFC 4122)                  |
| `uuid5`     | Name-based UUID (SHA-1)                 |
| `uuid6`     | Reordered time-based UUID               |
| `uuid7`     | Unix time-based UUID (RFC 9562)         |
| `ksuid`     | K-Sortable globally unique ID           |
| `xid`       | MongoDB-style globally unique ID        |
| `nanoid`    | Secure, URL-friendly random ID          |
| `snowflake` | Twitter-style time-based distributed ID |

## Examples

Generate 5 UUIDv4s:

```sh
genid uuid4 5
```

Generate 3 ULIDs:

```sh
genid ulid 3
```

Generate 2 NanoIDs of length 10:

```sh
genid nanoid 2 --len=10
```

Generate a Snowflake ID:

```sh
genid snowflake
```

## Notes

* UUIDv6 and UUIDv7 are based on modern drafts (such as RFC 9562) and implemented using reliable Go libraries.
* NanoID is customisable in length and safe for use in URLs.
* `--len=n` must be used *after* the count and only with the `nanoid` type.

## License

This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md).
