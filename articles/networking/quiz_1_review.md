Quiz 1 Review
===============

1. Local telephones use circuit-switched networks because you do not want a call to miss any packets else it will lead to lots of dropped messages within a given call.

2. A frame encapsulates a datagram.   
A frame goes with the data link layer.
The datagram goes with the network layer.   
The segment and user datagram go with the transport layer.

The frame adds a header to the datagram generated at the network layer. The act of taking something and adding onto it is called encapsulating. 


3. Packet-switched networks need support for both forwarding and storing capabilities because most packet switches use *store-and-forward transmission* in which the packet switch stores an incoming packet until all of its bits are stored properly on the switch and only then does it forward the packet to the next switch or end system.

4.
2 Key Principles for Protocol design
* opposite function occurs at each layer
  * there is logical and opposite connection between layers on opposite sides of the protocol. For example, if the left side is encrypting, then the right decrypts.
* symmetric
  * the layers to the left are in the same order as the layers to the right

5.

| Protocol | Layer |
| --- |---|
|SSH (secure shell)|application layer|
|DHCP (dynamic host configuration protocol)|application layer|
|SMTP (simple mail transfer protocol)|application|
|BGP (border gateway protocol)|application|
|DNS (domain name system protocol)|application|
|FTP (file transfer protocol)|application|
|HTTP (hypertext transfer protocol)|application|
|IMAP (internet message access protocol)|application|
|POP (post office protocol)|application|
|SMTP (simple mail transfer protocol)|application|
|telnet|application|
|MIME|presentation layer|
|TCP (transmission control protocol)|transport layer|
|UDP (use datagram protocol)|transport layer|
|SCTP (stream control transmission protocol)|transport layer|
|IP (internet protocol)|internet layer|
|MAC(media access control)|data link layer|


6. The presentation layer in the ISO/OSI model performs
* data compression
* data encryption

Congestion control is done in the transport layer.

Routing is done in the Internet layer
