### Run cmd
```bash
docker exec -it redis sh
redis-cli
auth hbb161112
config get maxmemory
setex mykey 60 "Hello world"
keys *
mget mykey
```