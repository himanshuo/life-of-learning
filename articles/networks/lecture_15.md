Lecture 15
===========
### rules for generating acknowledgements
  * when A sends a data segment to B, A must include (piggyback) an ACK given the next sequence number it expects to receive
    * reduces traffic
  * when receiver has no data to send and it receives an in order segment and the previous segment has already been acknowledged, the receiver delays sending an ack segment until another segment arrives or until a period of time (eg. 500 ms) has passed
    * saves number of ack segments
  * when a segment arrives with a sequence number that is expected by the receiver, and the previous in-order segment has not been acknowledged the receive immediately sends an ACK segment
  * when a segment arrives with an out-of-order sequence number that is higher than expected, the receiver immediately sends an ACK segment announcing the sequence number of the next expected segment
  * when a missing segment arrives, the receiver sends an ACK segment to announce the next sequence number expected
  * if a duplicate segment arrives, then the receiver discards the segment but immediately sends an ACK segment indicating the next in-order segment expected


TCP congestion control
  * congestion cannot occur on both sides because of the dynamic window resizing
  * congestion can heppen in middle because of routers
  * if congested, routers drop packets
  * define 2 windows  
    * congestion window (cwnd)
    * rwnd
  * send window size = MIN(cwnd, rwnd)

  * determine if a congestion has occured
    * 2 cases;
      * timeout
        * say receiving ACK but didn;t receivine for a long time.   
      * 3 duplicate ACKs
    * first one is more dangerous. the 3 duplicate ACKs mean that atleast you received three 3 packets.

4 basic congestion policies
  * slow start with exponential increase  
    * make cwnd very small (MSS - maximum segment size).
    * initialize to 1
  * congestion avoidance
  * fast recovery
  * 
