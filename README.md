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

## Advanced

### Building the Lua Scripts

The Lua scripts required for this project are shared between multiple
libraries. In order to build with them the bamboo-scripts submodule must be
initialized.

    - `git submodule init`
    - `git submodule update`
    - Run `go generate` in the project directory.

## Writing a Library in Another Language

### The Job Object

The Job object represents the a Job in the queueing system.

Fields:

    - Priority
    - JobID
    - Payload
    - Failures
    - DateAdded
    - DateFailed
    - Owner
    - ContentType
    - Encoding
    - State

Methods:

    - Serialization
        - ToStringArray(Job) -> []String
            IE. [key1, value1, key2, value2 ...]
        - FromStringArray([]String) -> Job

### The RJQ (Redis Job Queue) Object

Fields:

    - Namespace
    - Client Name (Worker Name)

Methods:

    - `Enqueue(job)`
        - Uses `bamboo-scripts/enqueue.lua`.
    - `Schedule(job, datetime)`
    - `Consume(job_id)`
        - Uses `bamboo-scripts/consume.lua`.
    - `Remove(...job_id)`
    - `Clear`
    - `Delete`
    - `Ack`
        - Uses `bamboo-scripts/ack.lua`.
    - `Fail`
        - Uses `bamboo-scripts/fail.lua`.
    - `RecoverAbandoned`
    - `ListJobs(queue)`
    - `GetCount(queue)`
    - `ListWorkers(queue)`
    - `Ping(expires)`
    - `MaxFailed()`
        - Set the maximum number of job failure for retry.
    - `MaxFailed(int)`
        - Retrieve the maximum number of job failures for retry.
    - `MaxJobs()`
        - Set the maximum number of simultaneously consumed jobs.
    - `MaxJobs(int)`
        - Retrieve the maximum number of simultaneously consumed jobs.

Parameters:

    - job: Job object.
    - datetime: Unix UTC seconds since epoch.
    - queue: One of queue, working, scheduled
    - expires: Seconds.
