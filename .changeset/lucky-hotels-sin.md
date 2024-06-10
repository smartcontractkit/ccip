---
"ccip": major
---

Make sure the Execution Plugin is resilient to zk overflow by

- reducing the batch size to 1
- avoiding infinite retries
