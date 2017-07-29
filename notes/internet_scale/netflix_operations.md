Service Oriented Architecture at Netflix
==========================================
Netflix moved to the cloud.

The old system was a large java application on a tomcat container in a large data center. The problems with this were
* lots of code had to be checked in for deployment
* a single bug could mess with the deployments
* the centralized Dev Ops team was in charge of making sure no bugs are put into production, while the dev team was just trying to push out features which led to a difference in goals.

During cloud migration, Netflix turned their giant system into hundreds of smaller services. Each service had its own deployment schedule. The centralized DEV OPS team got split up and each member was in charge of their own services. They created a new Site Reliability Engineering team which understood the system as a whole and analyzed it for reliability and resiliency.

Having developers truly own a product by having them handle all the way from development to deployment of their code gets developers to write better code. They do not want to stay up till 4am on call so they try and write good code the first time around. This also made it so that a few people ended up being highly specialized in a few portions of the system.
