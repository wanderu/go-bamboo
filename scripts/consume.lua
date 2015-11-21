-- [[
-- consume
-- Keys: <ns>
--  ns: Namespace under which queue data exists.
-- Args: <client_name> <datetime> [<key> <val> ...]

local fname = "consumejob"
local sep = ":"

local log_notice = function (message)
    redis.log(redis.LOG_NOTICE, "<" .. fname .. ">" .. " " .. message)
end

local log_verbose = function (message)
    redis.log(redis.LOG_VERBOSE, "<" .. fname .. ">" .. " " .. message)
end

local is_error = function(result)
    return type(result) == 'table' and result.err
end

local ns = KEYS[1]
local kqueue   = ns .. sep .. "QUEUED"   -- Primary Job queue
local kworking = ns .. sep .. "WORKING"  -- Jobs that have been consumed
local kmaxjobs = ns .. sep .. "MAXJOBS"  -- Max number of jobs allowed
local kworkers = ns .. sep .. "WORKERS"  -- Worker IDs
local kactive  = ns .. sep .. "ACTIVE"   -- Workers with active jobs

local client_name = ARGV[1]
local dtutcnow    = ARGV[2]
local job_id      = ARGV[3]
local exp         = ARGV[4]
local kclient     = kactive .. sep .. client_name

local max_jobs  = tonumber(redis.pcall("GET", kmaxjobs))
local njobs = tonumber(redis.pcall("ZCARD", kworking))

-- Don't consume more than the max number of jobs
if njobs and max_jobs and (njobs >= max_jobs) then
    return redis.error_reply("MAXJOBS")
end

-- The client name must be non-blank.
-- Blank indicates initial job state with no owner.
if client_name == "" then
    msg = "Invalid client name."
    log_error(fname, msg)
    return redis.error_reply(msg)
end

-- Add this worker to the active worker set
redis.call("SADD", kactive, client_name)
-- Make a key for this worker that expires.
redis.call("SET", kclient, 1)
redis.call("EXPIRE", kclient, tonumber(exp))


local kjob = ns .. sep .. "JOBS" .. sep .. job_id
return redis.call('HGETALL', kjob)
