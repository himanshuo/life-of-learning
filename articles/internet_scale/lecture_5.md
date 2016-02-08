Lecture 5 - single system databases
=====================

### Overview
Databases store data in tables with fixed schemas.
* all tables have an PK id
* joining is combining two tables

### Database transactions
Transactions are a group of reads/writes to the database
* make sure either they all happen or none happen
* make sure concurrent users never see an "in between state"
  * never see half the updates done
* these grouped updates are called transactions

### Scaling issues with transactions
There are two main components for how transactions are implemented: *logging* and *locks*.

Logging is a way to make sure the updates in a transaction all happen or not. The specific type of problem that logging helps to solve is:
* want to update two tables - table 1 and table 2
* update table 1 and then db hits a bug and crashes
* on restart, database is "inconsistent" as it shows the update to table 1 but not to table 2. This should NOT be allowed based on the rules of transactions.

### Consistency requirements 
