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
  * 3 way handshaking
  * 4 way handshaking (half-close)


### 3 way handshaking for TCP
IMAGE

Key:
  * client and server both have a process running and a transport layer that handles the sending of communications between the two
  * each of the boxes that are being sent across are called **segments** in the transport layer.
Steps:

1. server goes through **passive open**.
  * this just means that the server is going to allow connections. The passive part means that the server is listening for connections rather than making them.
  * notably, the server process has to tell the transport process to start listening
2. **active open** for client
  * client makes a connection to server
  * the sequence number is random
  * no ack number
  * 's' means that it is a SYN packet. The 'syn' flag is 1.
3. server responds with ACK+SYN
  * seq is random
  * ack is 1+last byte of current packet thus 1+8000=8001
  * 's' SYN
  * 'A' ACK
4. the transport layer of the client sends a syn to finalize the connection
  * seq is 8001 because previous seq + 1 = 8000+1=8001
  * **ack = 15001 because it is acknowledging the packet that the server sent it. cur seq + 1 = 15000+1=15001**
  * 'A' ACK


You generally can refer to these packets as the SYN, SYN-ACK, ACK packets.

If the SYN packet is lost, then the client will get no response and try again. Thus the SYN packet is verified by the SYN-ACK packet.

If the SYN-ACK packet is lost, then client will get no response and will retry. At the same time, the server will never get a ACK packet and so will resend the SYN-ACK packet. Thus the SYN-ACK packet is verified by the ACK packet.

**The ACK packet is verified by the first data packet**. The first data packet contains extra information in it to allow the server to know whether the ACK packet is lost or delayed. If it is lost or delayed, the server/client simply stop caring about it.


### data transfer when ACK packet is lost   
IMAGE

Steps:

1. assume connection has been established via 3-way handshake, however, the ACK packet is lost.
2. client sends a request
  * sequence number is 1+last byte of previous packet = 1+8000. Note that it is the same as the failed ACK packet (in the previous image).
  * ACK number is 15001 to acknowledge the SYN-ACK packet from the server (in previous image).
  * 'A' for ACK
  * 'P' for PUSH (assuming this is a multimedia type transmission so push it out right away)
  * data part uses up bytes 8001 to 9000. **payload size = total length - header length**. **header length = IP header length + TCP header length**.
    * Note, that the TCP header already contains total length so you never really have to actually do the header length = IP header + TCP header calculation.    
3. server process *pulls* the data from the server transport.
  * it pulls 1000 bytes, representing only the data part of the packet
4. client sends another request
  * sequence number 9001
  * ACK 15001 because 15000 is the last packet sent from server   
  * 'A' ACK, 'P' PUSH
5. server responds
  * seq 15001 because previous seq that server sent was 15000.
  * ack 10001 because previous packet the the server recieved ended on 10000. 10000+1=10001.
  * 'A' ACK because acknowledging previous packet
  * 'P' PUSH because sending data
  * sending 2000 bytes of data in this case
6. client responds
  * seq 10001 because current packet starts on 10001
  * ack 17001 because previous packet server sent ended on 17000. 17000+1=17001.
  * 'A' acknowledgement. No 'P' because this is not sending any data.

Fun Hint: packets that are sent that have no data in them (SYN, SYN-ACK, ACK) have data size 0. Thus following packets will have previous_packet_seq + 1.

Fun Hint: packets that are sent after a data packet will have x data size. Thus the packet will have seq previous_packet_seq + x (this accounts for the fact that the previous_packet_seq number contributes to the packet size. Thus no +1 is needed.)

Piggybacking is shown in this picture whenever you have both 'A' and 'P' in the same packet.

### Connection Termination
There are 2 ways to close the connection
  * 3 way handshaking
  * 4 way handshaking (also called half-close)
    * this is when client wants to close the connection but server still has something to send

#### 3 way handshaking
IMAGE

Steps:

1. **active close** by client means that the client initiates closing of the connection.
2. the client transport sends the closing packet
  * 'F' for FIN means client is initiating the close
3. servers process sees the FIN packet by the client and allows the transport layer to start closing the connection. The server process telling the server transport to do this is called **passive close**.
4. the server sends the response
5. client sees the server response on the client transport and does a **connection close**
6. client sends a final packet.
7. upon seeing final packet in server transport, server does a connection close.

If the last packet fails, then we simply don't care, as per the specifications. The server just closes itself eventually.


FUN FACT: if you get confused about what the SEQ is, just remember that the previously recieved packet must have had a ACK value. That value is the same as the SEQ value you are currently sending.

##### window size
Window size for sender is entirely determined by receiver. The receiver tells the sender what window size it should have.

Why? Because only the receiver knows what an effective window size would be for it to receive packets. The criteria that the receiver uses is whatever the window size can be used to do congestion control and flow control effectively.

IMAGE

* when the left edge of the window moves right, the window closes
* when the right edge of the window moves right, the window opens
* *you do not want the distance between the two to be too small else you will have problems with how many packets you are acknowledging*
