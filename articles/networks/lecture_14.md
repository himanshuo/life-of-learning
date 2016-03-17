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

 
