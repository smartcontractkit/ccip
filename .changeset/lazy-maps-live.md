---
"ccip": patch
---

Implementing batched blessing checks (performance optimization).

Exec plugin has a new field, a cache that expires every 4 hours.
On each round we make a batch call to the commit store for all the commit roots that are not present in the cache.
We keep in the cache only the blessed roots - this essentially makes the cache a `set`.

Apart from that, the logic for fetching the executable observations is slightly changed.
Instead of iterating through the reports and skipping invalid/already-executed, first we filter-out all this reports
and then we start the processing.

Using this approach, fetching the destination chain pool rate limits, adding jobs to the token data worker and
performing the blessing checks essentially ignore irrelevant reports which makes the execution plugin more robust.
