Lecture 12
===============

selective repeat protocol

http://www.bing.com/videos/search?q=selective+repeat+protocol&&view=detail&mid=7AAACDE4AB71E6A5F2CA7AAACDE4AB71E6A5F2CA&rvsmid=58F2EBCF2375E339739C58F2EBCF2375E339739C&FORM=VDFSRV&fsscr=0

Example of selective repeat protocol


You can make selective repeat protocol bidirectional by simply making sender the receiver and vice versa.
  called piggybacking
  must add



3 protocols:
stop and wait
G...
selective repeat



What does UDP do?
* all packets are independent. thus you want to use stop and wait.
* udp does support checksum. you can enable or disable it. Default is disable.
* no need to establish a connection
* you have to add payload


UDP header format
* uses 8 bytes
udp does not support flow control, congestion control, or proper error control (checksum is sort of it but not really)

Applications that use UDP:
* SNMP (simple network management protocol)
* multimedia applications
* broadcasting


### TCP
* does all duplex(bidirectional) communications

3 phases:
  * connection establishment using 3 way handshaking
  * data transfer
  * connection termination

TCP uses
* multiple timers
* GBN + SR
  * SR for each particular packet
  * following packet number is expected from ACK
* cumulative and selective ACKs
* supports retransmission of lost, corrupted packets (combination of GBN and SR, not exactly the same)

TCP is reliable. How do you make it reliable?
* ACK packets
  * allows you to


Numbering system is interesting.
* sequence numbers
  * use byte number for sequence numbers.
  * 
* ACK numbers
