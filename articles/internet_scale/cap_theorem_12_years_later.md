CAP Twelve Years Later: How the "Rules" Have Changed
=====================================================
CAP theorem says that any networked shared-data system can only have 2 of the three desirable properties. People are challenging this very much.

Desirable Properties
* C = Consistency = single up-to-date copy of data
* A = Availability = data is always available for use
* P = Partitions = can handle network partitions (so that you have tolerant data)


NoSQL databases focus on high availability and worry less about consistency.

ACID (atomicity, consistency, isolation, and durability) databases focus on consistency and worry less about high availability.


Why the CAP theorem is nonsense
* partitions are rare. Many systems don't have to partition their data. In this case, it is dumb to give up C or A
* choice between C and A can occur many times within a given system. Different layers of the system may make opposite choices. Specific users or specific bits of data may lead to different choices
* all three choices are more continuous than binary


ACID is
* Atomicity (A) - an operation fails completely or is successful completely.
* Consistency (C) - all database rules are followed (all constraints)
* Isolation (I) - 
