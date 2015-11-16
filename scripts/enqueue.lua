--[[
-- enqueuejob
-- Keys: <queue_key> <job_key>
-- Args: <priority> <job_id> <key> <val> [<key> <val> ...]
-- ]]

local fname = "enqueue"
local queue_key = KEYS[1]
local job_key   = KEYS[2]
local priority  = tonumber(ARGV[1])
local job_id    = ARGV[2]

local log_error = function (fname, message)
    redis.log(redis.LOG_WARNING, "<" .. fname .. ">" .. " " .. message)
end

local current_score = tonumber(redis.call("zscore", queue_key, job_id));

if current_score and priority >= current_score then
    msg = " Not enqueing item. "
          .. "An existing item has the same or lesser score.";
    redis.log(redis.LOG_WARNING, "<" .. fname .. ">" .. " " .. msg)
    return 0;
end

result = redis.pcall("HMSET", job_key, unpack(job_data));
result = redis.pcall("ZADD",  queue_key, priority, job_id);

return 1;
