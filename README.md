# Fluent HTMX

An HTMX extension for [Fluent](https://github.com/jpl-au/fluent). Wrap any Fluent element to add HTMX attributes through method chaining, plus server-side helpers for handling HTMX requests and responses.

## Choosing a version

HTMX 4 changes enough (attributes, request/response headers, config keys, events, swap strategies) that one package cannot serve both major versions cleanly. The library is split by htmx major version; import the one that matches the htmx you ship:

| htmx version | import path | docs |
|--------------|-------------|------|
| htmx 2 | `github.com/jpl-au/fluent-htmx/htmx2` | [README](htmx2/README.md) · [AGENTS](htmx2/AGENTS.md) |
| htmx 4 | `github.com/jpl-au/fluent-htmx/htmx4` | [README](htmx4/README.md) · [AGENTS](htmx4/AGENTS.md) |

Each package is versioned independently with its own Go module with its own tags (`htmx2/vX.Y.Z`, `htmx4/vX.Y.Z`), You change versions by changing the import path. The package is named `htmx` in both, so call sites read `htmx.New(...)` and only the import path carries the version.

## Licence

MIT
