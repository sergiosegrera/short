# Short
Minimal performant url shortener. Created using microservice standards, can be easily modified to add different transports, databases, middlewares.

## Structure
```
cmd/             -- main.go
config/          -- Env vars manager
db/redisdb/      -- Redis database implementation
models/          -- Data models
service/         -- Business Logic
transports/http/ -- Http interface
```

## Endpoints
```
GET /{id}                                    -- Gets url by id
POST /create { "url": "https://google.com" } -- Create new url
```

## TODO
* Url verification
* Http transport testing
* Front end?
