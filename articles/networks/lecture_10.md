Lecture 10
==============

### Stop and Wait
features:
* window size = 1
* supports checksum
* supports sequence number
* has timer

Steps:
1. sender sends packet.
2. Starts timer.
3. Waits until ACK comes back.

Determines range of sequence numbers by:
* if packet arrives safe and sound -> receiver sends ACK message.
  * Packet number is x.
  * ACK number is x+1.
* packet is lost or is corrupted at receiver side
  * receiver thus never received proper x
  * once timer runs out, sender sends x back again.
  * once receiver gets correct packe, it sends back x+1  
* packet arrives at receiver, but ACK is corrupted or lost on the way back
  * sender just drops corrupted packet.
  * timer will go off
  * sender will resend the original packet with value x (duplicate)
  * receiver will send back ACK x+1


Handle corrupted packets by dropping them. Do not send ACK.

Sender window only has one socket. Receiver window only has 1 socket.


**You only need 2 numbers for this, thus you only go between 0 and 1**


Save a copy at sender.

If timer runs out when you are in blocking state, then resend packet in window. Then restart timer. Thus remain in blocking state.

If corrupted ACK arrives, discard the ACK packet. Still in blocking state.

Slide send window forward. Start=start+1


How receiver handles 
