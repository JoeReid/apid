# APID

Friendlier UIUDs for APIs

APIDs are Base62 encoded UUIDs with a string prefix

```
Before: "ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"

After: "user_6hoOwWlutwzIKWFCp54MUb"
```

## TODO

- [x] CLI application for generating and translating APIDs
- [ ] Improve documentation of CLI command
- [ ] Add apid generation to CLI
- [ ] Improve CLI to use cobra style verbs (`apid gen`, `apid enc`, `apid dec`)
- [ ] Add support for custom separators to cli
