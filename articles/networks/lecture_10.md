Lecture 10
==============

### Stop and Wait
features:
* window size = 1
* supports checksum
* supports sequence number
  * **only need 2 numbers, thus switch between 0 and 1 sequence numbers**
  * sequence numbers are used to detect duplicates
* has timer

##### Sender
Sender window only has one socket. Receiver window only has 1 socket.

Sender Steps:

1. sender sends packet.
2. Starts timer.
3. Waits until ACK comes back.
4. if timeout, then sends packet again
5. when ACK comes back, slides window forward


##### Receiver

Receiver Steps:

1. has window of packets it wants to receive
2. receives a packet
3. validates packet
4. determines sequence number for ACK packet
5. sends ACK with correct sequence number

##### handling corrupted packets
Both sender and receiver Handle corrupted packets by dropping them. Do not send ACK.



##### Sequence Numbers
Determines range of sequence numbers by:
* if packet arrives safe and sound -> receiver sends ACK message.
  * Packet number is x.
  * ACK number is x+1.
* packet is lost or is corrupted at receiver side
  * receiver thus never received proper x
  * once timer runs out, sender sends x back again.
  * once receiver gets correct packet, it sends back x+1  
* packet arrives at receiver, but ACK is corrupted or lost on the way back
  * sender just drops corrupted packet.
  * timer will go off
  * sender will resend the original packet with value x (duplicate)
  * receiver will send back ACK x+1

Note that there is only 0 and 1 in this sequence. Thus, x+1 means 0 if x=1.


##### States
  Save a copy at sender.

  If timer runs out when you are in blocking state, then resend packet in window. Then restart timer. Thus remain in blocking state.

  If corrupted ACK arrives, discard the ACK packet. Still in blocking state.

  Slide send window forward. Start=start+1


How receiver handles
