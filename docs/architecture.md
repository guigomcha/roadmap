# Implementation decisions

## Database

### Redis 

Redis core has the key-value store but is compatible with multiple "modules". Additionaly, redis work as a async task broker.

- https://youtu.be/OqCK95AS-YE
- https://redis.io/nosql/document-databases/
- https://www.dragonflydb.io/guides/redis-memory-and-performance-optimization


Nice to have:
- Redis monitoring of key-value changes in real time + Client Side Caching
- Timeseries
- Automated snapshots and revocery (RBD + AOF)
- Sharding for distributed scaling.
- Lua scripting allows you to run complex operations directly on the Redis server. 

Drawbacks:

- RAM memory is always more expensive than disk. 
- The app is not concerned about real-time.
- Main storage will be relational data + multi-media which might not be ideal for redis

Decision:

- Could still be used as message broker and slowly adopted for client side cache 

### PostgreSQL

For relational data. 

### Multi-media storage

I need an object storage in disk to optimize storage.

- https://github.com/minio/minio
- https://ceph.io/en/discover/technology/
- https://www.tritondatacenter.com/triton/object-storage