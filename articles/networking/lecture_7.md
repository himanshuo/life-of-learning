Lecture 7
============

### DNS (Domain Name System)

URL -> 128.96.13.5 (ip address)

FTP client -> DNS client -> query -> DNS server

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

examples include
* napster

when you join a group, then everyone around you knows ip address.

### Centralized server approach
when you join centralized server, you have to provide
* your ip address
* list of files on your server
* allow file to be downloaded by other peers


can use ping and pong to determine whether server is alive.


issue is that all the traffic is going to server

### distributed
gnuetella - company that broadcasts message

extra

new joiners are put on bottom of tree (leaves). they cannot broadcast.

### dht (distributed hash table)
figure out hash function that creates no collisions
* sha-3 is a hash function that dht can use
  * digest = function(message)
* md5


DIAGRAM - really pay attention to this. it is very confusing.
