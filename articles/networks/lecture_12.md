Lecture 12
===============

### Selective Repeat Protocol

m = number of bits used

max window size = 2^(m-1)

buffer size = 2^(m)

The buffer size and max window size is the same for the sender and receiver.

Sender side has variables Sf and Sn
* Sf = start of window
* Sn = end of window

Receiver side has variables Rn
* Rn = start of window
* The end of the receiver window is Rn + (max window size - 1)


**packets can be sent and acknowledged out of order**

**packets can be received out of order**

**timer starts after the last packet in the window is sent**. Once the timeout occurs, you start sending from the **leftmost side** again. You only send packets that have **not already been acknowledged**

##### Why do we have a max window size?
Note that if **m=2**, then **max window size = 2 ^(m-1) = 2** and **buffer size = 2^m = 4**

To illustrate that the window size cannot be greater, we will make max window size = 3 in this **INVALID** example.

![](lecture_12-images/e0bcc15ab8f2d0aee001f239a60febb1.png)

Key:
* The yellow means that that packet has been sent
* red outline reveals the window for the receiver
* green means that received packet that is in middle of window
* Sf and Sn specific window for sender
* Rn specifies window for receiver
* The initial window is only at 0 for sender
* The initial window is from 0-2 at receiver

Understand Flow of Image:

1. the sender sends packet 0
2. packet 0 is marked as being sent
3. once packet 0 arrives, window is slid over by 1
4. receiver sends acknowledgement for packet 0, but it fails to reach sender
5. repeat steps 1-4 for packet 1
6. repeat steps 1-4 for packet 2
7. Since the window size is only 3, we cannot send the next packet. So, we send the leftmost packet in our window - packet 0. (We are at 4th image at sender side)
8. *On the receiver end, packet 0 has technically already been received. But, we have no way of distinguishing it from the 0 in our current window*. **This is because our window size is too big**. This would not occur if the window size were smaller. It is marked green because the packet is in the middle of the window.

Thus, we should have a **max window size = 2 ^ (m-1)**


##### Properly working example of Selective Repeat Protocol
m=3

max window size = 2^(m-1) = 4

buffer size = 2^m = 8

IMAGE

Flow:
1. packet 0 is sent and marked as sent
2. packet 0 is received thus receiver window is slid by 1
3. acknowledgement for packet 0 is sent
4. acknowledgement for packet 1 is received and thus sender slides window by 1
5. packet 1 is sent and marked as sent
6. packet 1 is lost
7. since packet 1 never came to the receiver, the receiver does nothing related to it
8. packet 2 is sent and is marked as sent.
9. receiver realizes that packet 2 came out of order. It marks packet 2 has having been received and sends an acknowledgement for packet 2.
10. sender receives acknowledgement for packet 2 and marks it as acknowledged. It realizes that packet 2 was acknowledged out of order.
11. sender sends packet 3 and marks it as being sent.  
12. receiver marks 3 as received and sends an acknowledgement for packet 3.
13. the window is done. The sender was not able to move the window since the start of the window has not been acknowledged. So, the sender starts the timer.
14. the timer runs out.
15. the sender restarts sending from the leftmost side of the window. It sends only packets that have not been acknowledged.
16. Thus, the sender sends packet 1.
17. the receiver receives packet 1, marks it as having been received, and sends an acknowledgement for packet 1
18. the sender gets packet 1's acknowledgement and moves the window all the way up to 4. It does so because packets 2 and 3 have already been acknowledged as being sent correctly.
19. The sender would at this point send the next available packet - packet 4.
20. the receiver also moves its window all the way up to packet 4 because packets 2 and 3 have already been acknowledged as being received.
21. the process continues...



##### bidirectional
You can make selective repeat protocol bidirectional by simply making sender the receiver and vice versa.
  called piggybacking
  must add



3 protocols:
stop and wait
G...
selective repeat


### UDP
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


header : 20-60 bytes
  * 20 bytes normal, but can have optional extras (40 bytes of options)
  * first 20 bytes have
