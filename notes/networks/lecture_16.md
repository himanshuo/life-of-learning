Lecture 16
===============

2 types of windows: sent and receive

### 3 policies
* slow start with exponential increase
  * 1 MSS
  * aggresive
* congestion avoidance
  * additive increase
  * conservative
* fast recovery
  * used to take care of 3 duplicate ACK's

### Taho TCP
uses slow start with exponential increase + congestion avoidance polices

### Reno TCP
uses slow start + congestion avoidance + fast recovery policies


throughput = [ (max + min ) /2 ] / RTT = ([max + .5 max ]/2) / RTT = .75 max / RTT


min = 1/2 * max


MSS = 10kb, RTT = 100ms

using image. 10 is max segment size
w_max = (10+12+10+8+8) / 5 = 9.6 mss

throughput = 10.75 * 9.6 * 10 kb   / 100ms = 7.2 Mbps is average throughput using diagram



### 4 timers
* retransmission timer
* persistence timer
  * when server sends 0 to client, then client has to stop transmitting packets to it
  * client sends persistence timer
  * after some time, client sends probe to server asking for non-zero packet. If failed, then will double persistence timer and send again, .... Until threshold. After threshold has been reached, then stay there.
* keep alive timer
  * if client has crashed, and server has gotten latest request from client then server will start keep alive timer.
  * timer waits 2 hrs. Then 10 probes are done (each send atleast 75 seconds apart). Then, server can assume client is dead.  
* TIME-WAIT timer
 * maximum segment lifetime (MSL)
 * each segment has a lifetime that specifies TTL (time to live)
 *

 KEY POINT: read tranmission timeout is very important. 
