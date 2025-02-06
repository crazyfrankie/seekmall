local key = KEYS[1]
local cntKey = key..":cnt"

local val = ARGV[1]

local ttl = tonumber(redis.call("ttl",key))

if ttl == -1 then
--  有人误操作导致 key 冲突
    return -2
elseif ttl == -2 or ttl < 540 then
--  后续如果验证码有不同的过期时间，要在这里优化
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", cntKey, 600)
    return 0
else
--  已经发送了一个验证码，但是还不到一分钟
    return -1
end

