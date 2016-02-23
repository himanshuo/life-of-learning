Lecture 7
============

### DNS (Domain Name System)
Converts url to ip address.

URL -> 128.96.13.5 (ip address)

### How to determine IP address of the server you want to contact


* you input some url into the browser(http client)
* http client -> local DNS client in order to convert url
* local DNS client -> local DNS server
* local DNS server -> local DNS client
* local DNS client -> http client

A few key things are omitted - how does your local DNS server know about 


* DNS client(also called resolver) needs to know ip address of DNS server.
* ISP gives the DNS client the ip address of the DNS server.
  * when you connect ISP with DNS for the first time, then your DNS client will be configured to store the DNS servers ip address  
* ttl (time to line) is used to determine whether the cached DNS server ip address is
  * ttl is just how long the ip address of the dns server
* The DNS server then converts the URL to the ip address of the resulting  



country name (top most level)
* ca - china
* ca - canada
* fr - france
* us - usa

if you are in usa and you are sending to usa, then you can omit usa

nature of organization (secondary-level)
* aero - aerospace/airlines
* com - company
* biz - business
* coop - cooperative
* edu
* gov
* info
* int - international organization
* mil - military
* name - individuals
* net - companies providing network services  
* org - nonprofits
* pro - professional organizations
* museum

<server_name>.com
server names must be unique.


Two ways to resolve name:
* recursive
* iterative

### Recursive name resolution
DIAGRAM
### iterative name resolution


nslookup command gives you ip address.
if you are on cs.uva.edu, then nslookup will return cs.uva.edu
If you use nslookup and you will also get a authoritative result as well. You can check whether the first value is from cache or not by checking if authoritative and first value are different.



DNS
512 bytes



Primary DNS server
Secondary DNS server

Both servers have identical values


Only TCP can handle ...

### registrars
registrars are under the control of ICANN

when you register a new domain name, you need to give the registrar
* ip address
* domain name

Dynamic Domain Name Service(DDNS) - automated

### attacks to DNS server



encryption or some sort of security is used to prevent attacks to DNS server


digest = f(m)


# Peer to peer (P2P) approach


when you join a group, then everyone around you knows ip address.

### Centralized server approach
Central server holds all the information. It holds
* each peers IP address
* list of files on each peer

when you join centralized server, you have to provide
* your ip address
* list of files on your server
* allow file to be downloaded by other peers

When you want a file from this centralized server approach, you contact the central server. The centralized server authenticates you. Then, the server sends you the IP address of the peer to contact and work with.

##### Problems with this approach
* centralized information holds too much info
  * lots of centralized memory
  * if server goes down, then all information is lost
* you have to contact centralized server and then also the peer
* if something changes in a peer, it must be communicated to centralized server
* all the traffic is going to server

can use 'ping' and 'pong' to determine whether server is alive.



### Distributed


### dht (distributed hash table)
figure out hash function that creates no collisions
* sha-3 is a hash function that dht can use
  * digest = function(message)
* md5


DIAGRAM - really pay attention to this. it is very confusing.



### Examples

* Client-Server
  * HTTP
  * FTP
  * SMTP
  * SSH
  * Telnet
* P2P
  * Centralized
    * napster
  * Distributed
    * gnutella

#### gnutella
gnutella is company that broadcasts message
* when you join group, you are assigned neighbors
* if you want to access a file, you ask neighbors. If neighbors don't have it, then broadcasts query to all nodes.
  * creates a tree of joiners. new joiners are put on bottom of tree (leaves). they cannot broadcast.
