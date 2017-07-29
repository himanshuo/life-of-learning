Adopting Microservices at Netflix: Lessons for Architectural Design
=====================
A Microservices architecture is a service-oriented architecture composed of loosely coupled elements that have bounded contexts.

loosely-coupled = update services independently. Each service should seriously should NOT require other services to also do stuff. They need to be very independent. A key thing to split up is the database. You do so by normalizing the database.

bounded contexts = self-contained service. You can understand and update the micro-service without knowing anything about the internals of its peers. A micro-service interacts with its peers entirely through APIs. Because of this, micro-services also don't share data structures, or database schemas.

Each micro-service must be stable and its older API should still be supported, even when new versions come out.

### Best Practices for designing a micro-services architecture
##### Create a separate data store for each micro-service
Each service should have its own data store which cannot be affected by others.

This can be difficult because data can get out of sync or become inconsistent. Thus you should use a tool that performs master data management (MDM) by operating in the background to find and fix inconsistencies. For example, a MDM tool will examine every database that stores subscriber IDs to verify that the same IDS exists in all of them. This ensures that there aren't missing or extra IDS in any one database.

##### Keep all code in a micro-service at a similar level of maturity and stability
If you need to add or rewrite some of the code in a deployed microservice thats working well, then the best approach is to create a new microservice for the new or changed code. You leave the old microservice in place until this new microservice is bug free and maximally efficient and then just replace. This reduces risk of failer or performance degradation in the existing microservice.

This is called the *immutable infrastructure* principle.

##### Do a separate build for each micro-service
A separate build for each micro-service allows you to pull in the appropriate versions of files. This will certainly lead to you pulling in the same file with different versions. This will make it difficult to remove old files as you have to make sure no microservice is using it. However, the benefit of easily being able to deploy something quickly by simply adding whatever version of the files that you need make this feature worth it.

##### Deploy in containers
This allows you to use a single tool to deploy everything.

##### Treat Servers as Stateless
Think of servers as interchangeable members of a group. You should not be concerned about a specific server. A given server should also not perform specialized work.
