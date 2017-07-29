Lecture 14
=============
### phases of TCP

1. connection established
  * 3-way handshaking
2. data transfer
3. connection termination
  * 3 way handshaking
  * 4 way handshaking (half-close)

### 4 way handshaking (half-close)

total of 4 windows:
* server receiving
* server sending
* client receiving
* client sending


sender will sent 0 window size to tell the reciever that it is no longer sending anything. there is no more space.



### silly window syndrome
* sender side
  * rlogin
  * 20 bytes for IP header + 20 bytes for TCP header + 1 = 41 bytes
  * we are only sending 1 byte worth of actual data, but because of headers, there are 41 bytes sent
  * solved via Nagels algorithm
* reciever side
  * this is syndome on recieveer side. let solve the problem different.
  * solved via a. clark's solution
  * solved also by delay ACK
    * creates another problem - timer will run out so will resend what it already sent


##### nagels algorithm
max segment size(MSS)
576 - 1500 bytes

wait until ACK comes back or until max segment size is reached


### Error Control for TCP
Problems that can occur:
  * corrupted segments
  * lost segments
  * out of order segments
  * duplicates

Tools:
  * checksum
  * timer
  * acknowledgement segments

ACK:
  * cumulative ack's (positive)
  * selective ack's (SACK's)

rules for generating acknowledgements
  * when A sends a data segment to B, A must include (piggyback) an ACK given the next sequence number it expects to receive
    * reduces traffic
  * when receiver has no data to send and it receives an in order segment and the previous segment has already been acknowledged, the receiver delays sending an ack segment until another segment arrives or until a period of time (eg. 500 ms) has passed
    * saves number of ack segments
  * 
