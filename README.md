# Bamboo

A reliable work queuing library backed by Redis.

## Purpose

1. Priority queue implementation for work prioritization.
2. Reliability (Ack/fail semantics) for guaranteed job processing.
3. Distributed without race conditions among workers.
4. Ability to view and modify queue contents (cancel job/clear queue).
5. Global queue controls. IE. Setting the maximum number of
simultaneously consumed jobs (MAXJOBS).
6. Clients that can be written in any language.

Most queuing libraries only have a subset of these properties.

## Namespaces

Namespaces are colon-separated string segments under which a queue set
will exist. For example, an application may have 1 set of queues under
`MYAPP:AJOBS` and another queue set under `MYAPP:BJOBS`.
