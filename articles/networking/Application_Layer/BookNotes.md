#Credits
These notes are derived from class lectures and 
Computer Networking: A Top-Down Approach, Ed.6, by James F. Kurose and Keith W. Ross
Efforts have been taken to paraphrase, but if there is a complaint due to similarity of phrasing please contact us and it will be remedied.

#Network Application Architectures
##Client Server Paradigm
The client-server paradigm demands that there is an always-on host that called the server.
The clients are hosts that don't have to always be connected, and requests from the client are serviced by the server.

Clients **don't** communicate with one another, only with the servers.

Data centers are basically large groups of servers grouped together to be able to handle many requests at once.
    Big companies generally use these, e.g. Google

##P2P Paradigm
Minimal/no reliance upon dedicated servers.
Exploits communication between pairs of hosts called peers (which are generally home/office computers).
Examples include BitTorrent, Skype.

P2P systems are **self-scaling**, meaning that each additional computer adds load, but also adds load-handling capacity.
Also, more **cost effective**.

###Major Challenges to P2P
1. ISP Friendliness: residential ISPs designed for more download than upload, but P2P doesn't have this same focus
2. Security: Challenge to secure because they're very distributed, open to many people
3. Incentives: Depends on people being willing to share burden of bandwidth, storage, computation resources

#Socket Interface
Sockets provide the link between the application layer and the transport layer. 
They provide ways by which data can be off-handed from one host.
They do not do anything about how the data is handled after the initial offload from the host, i.e.
they assume there is something existing outside of the host to handle movement of the data.

Sockets provide APIs to connect applications to networks.

#HTTP
HyperText Transfer Protocol = HTTP
Basically, HTTP defines how clients and servers need to structure the messages they send back and forth
HTTP is a **stateless** protocol, because it does not have a memory. This is why when you double link a download link, two files will
be downloaded. 

The transport layer protocol that HTTP interfaces to is TCP.

##Non-Persistent Connections
In essence, each request and response between a client and server will occur over a brand new TCP connection

##Persistent Connections
In essence, each request and response between a client and server will occur over the same TCP connection.

#Cookies
HTTP has no memory, which can be inconvenient for developers at times, although is great for simplifying server design to allow for
many TCP connections to be handled in parallel.

In order to make user-specific web applications, HTTP uses cookies, which store data about users.

Cookies consist of 4 components (**direct from book**):
1. header line in HTTP resp. msg
2. header line in HTTP request msg. 
3. cookie file kept on the user's end system and managed by the user's browser
4. back-end database at the website

The browser is the one that handles the cookie file, and it looks like there is a single cookie file per browser based on what's written
in the book.

#Proxy Servers
Think proxy wars. Proxy servers handle requests for other servers. In proxy wars, the frontline fighting entities are not the only ones
fighting, there is a "hidden figure" also involved.

What happens (**Phrasing from book**):
1. Browser est. TCP conn. to Web cache and sends HTTP req. for obj. to cache
2. Cache checks to see if copy of obj. stored locally. If yes, cache ret/s obj. within HTTP resp. msg to client browser
3.                                                     If no, cache opens TCP conn. to origin and gets obj from origin server
4. When cache gets obj, it stores copy in local storage and sends copy in HTTP resp. msg to client browser

Advantages of Web Caching
* Reduces amount of time client has to wait for response from server
* High speed conn/s often exist between clients and caches, again affecting speed
* Reducing traffic on server end by keeping local copy itself, can thereby reduce costs and groups can delay equipment upgrade

