# Bamboo

A reliable work queuing library backed by Redis.

## Purpose

1. Priority queue implementation for work prioritization.
2. Reliability (Ack/fail semantics) for guaranteed job processing.
3. Distributed without race conditions among workers.
4. Ability to view and modify queue contents (cancel job/clear queue).

Most queuing libraries only have a subset of these properties.
