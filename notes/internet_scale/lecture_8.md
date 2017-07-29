Lecture 8 - Databases cont.
===========================

### Spanner
* sharding: data is shard by KEY (every row has unique key)
* replication: each shard can be replicated
  * clients talk to directory servers to determine which shard a given piece of data is stored on
  * shards copy data consistently to replicas
  * data is not considered "stored" until all replicas confirm data is stored
  * one replica is elected "leader" and coordinates each update across other replicas
  * if replica fails, then you use a leader to determine what the true value is
    * if leader fails, you use majority to determine true value
* you need very precise time on server and timestamps on each update to make spanner work
  * create time line that all updates occur on ad make sure none happen at exactly same time



##### Challenges
* consistent update of replicas slows down system
  * 10-100x slower than cassandra
  * but consistency makes it super easy for programmers
* dependent on all machines having same idea of current time
  * distribution of time through low latency networks
  * use GPS and atomic clocks throughout data centers to synchronize time
