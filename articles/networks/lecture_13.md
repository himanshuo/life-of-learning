Lecture 13
===========
### TCP headers

![](lecture_12-images/3c31fb718f91bb91450e93040df1b0a7.png)

Features
* header, as a whole, is 20-60 bytes.
* source port number (16 bits)
* destination port number (16 bits)
* sequence number (32 bits)
* ACK number (32 bits)
  * the sequence number of the *next* packet I am expecting = last byte number of current packet + 1
* HLEN (header length) (4 bits)
  * each number from 0-15 represents a word(4 bytes)
  * thus (2^4) * 4 = 64 different values
  * This header can be 20-60 bytes. Thus, 64 different values is more than enough to represent length of entire header
  * NOTE: total length is stored in the IP datagram in network layer for TCP
* reserved (6 bits)
  *
* URG (1 bit)
  * urgent flag
  * if set to 1, then urgent pointer is valid
* ACK (1 bit)
  * if 1, then this is an ACK packet
* PSH (1 bit)
  * if you have a large window, you can end up accumulating packets. You pump packets out once the current window is done. This is slow.
  * if 1, then as soon as you get a packet, you pump it out. This is good for multimedia applications. This can lead to discontinuity of packets.
* RST (1 bit)
  * reset
  * abort connection
* SYN (1 bit)
  * for handshaking
* FIN (1 bit)
  * for termination phase
* window size (16 bit)
* checksum (16 bit)
  * mandatory in TCP
  * calculated in the same
* urgent pointer (16 bit)
  * tells you what part of packet is urgent
  * urgent pointer is just offset in packet that is urgent
  * what 'urgent' means is dependent on specific use case
* options (up to 40 bytes)


### phases of TCP

1. connection established
  * 3-way handshaking
2. data transfer
3. connection termination


### connection 
