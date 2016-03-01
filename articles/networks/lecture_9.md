Lecture 9 - Transport Layer
=========
### Multiplexing and Demultiplexing for Transport layer
This involves sending a process on one end-system to another.

##### Goal of Transport Layer
On the destination side, the transport layer receives segments from the network layer just below it. The transport layer then delivers the segment to the appropriate process in the applications layer.

##### Sockets for transport layer
Transport layer sockets require:
* sockets have *unique identifiers*
* each segment has *special fields* that indicate socket to which segment is to be delivered
  * source port number field
  * destination port number field

##### port numbers
TCP/IP uses 16 bit port numbers. Thus you have 65536 different port numbers possible. This means that you can have a total of 65536 processes running to send your segment to.

*well known port numbers*  
0 - 2023 port numbers are reserved.
* 0-1023: well defined (HTTP). These are system ports.
* 1024-49151: register region port numbers. These are user ports.

The last region of port numbers can be custom. These are called dynamic/private ports. The region is 49152-65535.

*common well known port numbers*  
* 161 - snmp - collect and organize info about devices on IP network
* 162 - snmp-trap - error handling

*How to determine port number for a given service.*  

    grep <process name> /etc/services

For example

    grep tftp /etc/services

/etc/services is a file that holds port number information about well known processes

##### Socket address
Socket address contains
* port number (from transport layer)
* IP address of host (from network layer)

##### Encapsulation
While going down protocol tree (application to transport), the various layers add more and more information onto the message. The fact that the transport layer takes the message from the application and wraps it up with the information that it needs to put on the message is called *encapsulation*.

##### Decapsulation
When package is received by reciever, then message loses contents at each layer. Each layer is decapsulating.


##### Demultiplexing in transportion layer
In the destination end, transportion layer needs to determine which process to send an incoming packet to. This is *one-to-many demultiplexing* because a single packet can go to numerous different processes.



### Three things that the transport layer has to handle
* flow control
* error control
* congestion control

### Flow Control
Control rate of incoming and outgoing packets.

Consider producer/consumer problem - producer can push something to consumer. Don't care if consumer is ready or not.

Flow control is more required in the producer to consumer region.


##### Push method
Flow control is needed to make sure that the consumer does not get overwhelmed.

One way to implement flow control in this scenario is to use a buffer pool. Producer can keep on pushing until buffer pool is not full. Once buffer pool is full, consumer can tell producer that it needs to stop.

##### Pull method
Flow control is **not** needed to make sure that the consumer does not get overwhelmed.

Consumer just asks producer to send it data when it is ready.

##### Buffer
Flow control is attained by using a buffer. The buffer stores information and notifies the sender when to stop sending. You collect data in the buffer and do not send until it is full.

The size of the buffer is called the *window size*. The window size can be *enlarged* when the buffer is *underloaded*. This means that the buffer is too small.


### Error Control
4 types of bad things that can happen
* packets can be lost
* corrupted packets
* out of order arrival of packets
* packet is sent in duplicate


##### Lost packets
This occurs when a sender sends a packet and the receiver does not receive it.

You handle this issue by setting up a timer that waits for x amount of time. It is waiting for an ACK. If ACK does not come, then resend packet.  

##### Corrupted Packets
Receiver received packet. It uses the checksum of the data to determine whether the packet was sent correctly. The checksum of the correct packet is sent as part of the packet.

You handle this issue by simply not sending ACK back to sender. The sender will then send you a duplicate copy of the packet.

##### Out of order packets
You know if the packets you are receiving are out of order if sequence number is off. You just add assign each packet a number in a sequence. This is done in binary so it's not too space-consuming.

The reciever You can handle this issue by waiting for the . You just wait until the duplicate copy comes in.


##### Duplicate Packet
This occurs if you send a packet but then do not receive an ACK. You will send another copy of the packet. This leads to a duplicate packet being sent the receiver.
