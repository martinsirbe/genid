### Features

* **Support for Modern ID Formats**: Generate a variety of unique identifiers, including ULID, KSUID, XID, NanoID, Snowflake, and UUID versions 4–7.
* **RFC-Compliant UUIDv7**: Implements UUIDv7 using a reliable community library, compatible with RFC 9562 for time-based UUIDs.
* **Flexible Output Options**: Easily configure the number of IDs to generate with `-n` and, for NanoID, customise the length with `-len`.
* **Simple CLI Design**: Intuitive flags and output make `genid` ideal for shell scripting, debugging, or quickly generating mock data.
* **Cross-Platform Friendly**: Written in pure Go with no external dependencies, works seamlessly on macOS, Linux, and Windows.
