# HStore
this repo and its content is just for learning purpos. I used golang hash maps as a data store, connecting to it using redis protocol [RESP](https://redis.io/topics/protocol).
### Use It
```bash
go run main.go
redis-cli -p 2654 # Let the magic happends :)
redis-benchmark -p 2654
```
### Benchmarks
CPU: i7 8750H 2.20GHZ
RAM: 16 GP DDR4
Hard Drive: SSD 250GB
OS: Ubuntu Linux Distribution.
```
100000 requests completed in 2.71 seconds
36913.99 requests per second
```
