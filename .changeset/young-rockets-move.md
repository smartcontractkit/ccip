---
"ccip": patch
---

Remove dependency on pg.qopts.

Postgres opts are not supported anymore. We are replacing all the `pg.qopts` params
with `context.Context`, it should contain all the relevant information for the underlying implementations.
