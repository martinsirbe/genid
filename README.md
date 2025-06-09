# genid — Generate Unique IDs

`genid` is a fast and simple command-line tool written in Go that generates various types of unique 
identifiers. It supports modern, widely used formats including ULID, UUID versions 4–7, KSUID, XID, 
NanoID, and Snowflake.

Perfect for scripting, testing, and inspecting different ID formats in distributed systems.

## Installation

### 📦 With Homebrew (Recommended)

If you use [Homebrew](https://brew.sh), you can install `genid` using a custom tap:

```sh
brew tap martinsirbe/clinkclank
brew install martinsirbe/clinkclank/genid
```

### Run Without Installing

If you have Go installed, you can run it directly:

```sh
go run idgen.go -type uuid4 -n 3
```

If you have Go installed, you can run the tool directly:

```sh
go run cmd/genid/main.go -type uuid4 -n 3
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
genid -type <id_type> -n <count> [-len <length>]
```

### Flags

| Flag      | Description                                  |
|-----------|----------------------------------------------|
| `-type`   | Type of ID to generate (see list below)      |
| `-n`      | Number of IDs to generate (default: 1)       |
| `-len`    | Length of NanoID (only applies to `nanoid`)  |

---

## 🔢 Supported ID Types

| Type        | Description                                 |
|-------------|---------------------------------------------|
| `ulid`      | Lexicographically sortable UUID-like ID     |
| `uuid4`     | Random UUID (RFC 4122)                      |
| `uuid5`     | Name-based UUID (SHA-1)                     |
| `uuid6`     | Reordered time-based UUID                   |
| `uuid7`     | Unix time-based UUID (RFC 9562)             |
| `ksuid`     | K-Sortable globally unique ID               |
| `xid`       | MongoDB-style globally unique ID            |
| `nanoid`    | Secure, URL-friendly random ID              |
| `snowflake` | Twitter-style time-based distributed ID     |


## Examples

Generate 5 UUIDv4s:

```sh
genid -type uuid4 -n 5
```

Generate 3 ULIDs:

```sh
genid -type ulid -n 3
```

Generate 2 NanoIDs of length 10:

```sh
genid -type nanoid -n 2 -len 10
```

Generate a Snowflake ID:

```sh
genid -type snowflake
```

## Notes

- UUIDv6 and UUIDv7 are based on modern drafts (such as RFC 9562) and implemented using reliable Go libraries.
- NanoID is customisable in length and safe for use in URLs.

# License

This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md).
