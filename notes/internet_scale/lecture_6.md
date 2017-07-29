Lecture 6 - Databases, pt2
============================

### Types of solutions
![](lecture_6-images/e64468a9420fecc0d40ea2edc1f1dbe1.png)


![](lecture_6-images/0f463d2a60c401df3281c180628c94fb.png)
sharded - how split apart the database is


### SQL vs NoSQL
SQL is a query language

### Sharding
Sharding is dividing data across many machines

Clients need strategy to know which machines store which data
  * fixed mapping, hashing

Want to avoid "hot spots" by sending equal load to all shards
  * eg. make sure highly accessed users/listings on different shards

![](lecture_6-images/eb8904fcef95798695423637bbafa371.png)
The image is a sample partitioning system to determine which data is on which shard.

Problems can occur
* different clients have different idea of what machine to use


### Locking / Transactions
Ensure consistency by making sure sets of updates are ACID.
  * ACID: Atomic, Consistent, Isolated, Durable
  * Atomic - change occurs or does not occur. No in between.
  * Consistent - only valid data is written to db. If a transaction is executed that violates the db consistency rules, the entire transaction will be rolled back.
  * Isolated - multiple transactions occurring at the same time cannot impact each others execution. If two transactions come at the same time and affect the same time, then the transactions should be done sequentially.
  * Durable - any transaction committed to the database will not be lost. This is done via backups and transaction logs that restore committed transactions in the case of software or hardware failures

Usually enforced by locking rows that are affected by query/update
* locks out other users so lowers performance

![](lecture_6-images/1ad2f61b8c26115d756c77334ba9012b.png)

### Replication
make copies of data

techniques to keep copies in sync
  * block reads until all copied data is updated
  * allow low consistency reads

### Columnar databases
Columnar Databases: store columns contiguously thus scanning multiple columns
