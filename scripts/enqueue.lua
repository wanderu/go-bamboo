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

local log_error = function (message)
    redis.log(redis.LOG_WARNING, "<" .. fname .. ">" .. " " .. message)
end

local log_info = function (message)
    redis.log(redis.LOG_NOTICE, "<" .. fname .. ">" .. " " .. message)
end

log_info(cjson.encode(ARGV))
log_info("priority: " .. tostring(priority))

local msg, result;
local argc = tonumber(table.getn(ARGV));

local job_data = {};
for i = 3, argc do
    table.insert(job_data, ARGV[i]);
end

-- How many key/value paramters?
local n = tonumber(table.getn(job_data));

local current_score = tonumber(redis.call("zscore", queue_key, job_id));

if current_score and priority >= current_score then
    msg = " Not enqueing item. "
          .. "An existing item has the same or lesser score.";
    log_error(msg)
    return 0;
end

-- result = redis.pcall("HMSET", job_key, "priority", priority);
-- result = redis.pcall("HMSET", job_key, unpack(job_data));
log_info("queue_key: " .. tostring(queue_key)
         .. " priority: " .. tostring(priority)
         .. " job_id: " .. job_id)
result = redis.pcall("ZADD",  queue_key, priority, job_id);

return 1;
