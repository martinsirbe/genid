### Features

* **Support for Modern ID Formats**: Generate a variety of unique identifiers, including ULID, KSUID, XID, NanoID, Snowflake, and UUID versions 4–7.
* **RFC-Compliant UUIDv7**: Implements UUIDv7 using a reliable community library, compatible with RFC 9562 for time-based UUIDs.

### Changes

* **Simplified CLI Interface**: Replaced the `flag`-based syntax (`-type`, `-n`, `-len`) with cleaner, positional arguments. You now specify the ID type as the first argument, an optional count as the second, and `--len=n` only if using `nanoid`.    
    **Before:**
    ```sh
    genid -type uuid4 -n 3
    genid -type nanoid -n 2 -len 10
    ```
    **Now:**
    ```sh
    genid uuid4 3
    genid nanoid 2 --len=10
    ```

* **Improved Help Output**: Added a `help` and `-h` command that displays usage instructions, argument descriptions, and examples.
* **NanoID Length Handling**: `--len=n` is now strictly scoped to `nanoid` only and must follow the count argument if provided. Invalid usage (e.g., `--len` with non-nanoid types) results in a clear error message.
* **Removed `flag` Dependency**: All command-line argument parsing is now handled via `os.Args`, resulting in a more intuitive and script-friendly UX.
* **Backwards-Incompatible Change**: This release breaks compatibility with previous flag-based CLI usage. Scripts or aliases using the old `-type`/`-n` flags must be updated to the new positional format.
