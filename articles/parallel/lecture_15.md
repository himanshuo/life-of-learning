Lecture 15
===========


packet switched or network switched networks
  * packet switch - send one by one to each node
    * a node may be too full and may drop your packet
      * a node will have a buffer or a queue and that is what is gets fulled
    * you may get tree saturation
  * circuit switch - packet locks each path. slow initial connection to establish path. Then, actual packet is sent without much delay
    * if intial connection thing goes to a node and it is being used, then you have options
      * wait for node to open
        * if you ONLY hold, then you could deadlock the network
        * the way you avoid deadlock is by routing in a particular order
            * this means that you may block even though there is a valid path
      * find another node
        * ???
      * just stop and let initial sender know


### symmetrical d-dimensional lattice
any partition to this d-dimensional hypercube


d-dimensional hypercube will be named by d-digit binary number


GOAL: understand how to determine whether a system can scale
  * how to just increase size of machine and still allow it to scale
  * communication network is not a constant -> boils down to granularity



Performance bounds
  * some device I will saturate first
  * V_iS_i (visit ratio * amount of server)- average amount of service request at device i by each message
    * when this is 1, this device is saturated
    * prob that each will land on it is 1
  * x0 = rate at which we process messages
  * x0 <= 1/ V_bS_b  (max ratio of service)
    * something on the denominator will eventually be saturated.
    * V_bS_b = max_i(V_iS_i)

granularity
* scales
* K= number of nodes
  * all nodes are ???
  * thus V_pe =1/k (visit ratio)
    * have k nodes. What is probability that you will go to a particular node?
      * uniform routing: 1/k
      * ? routing: 1/k because you will pick among only k things. Nodes are uniformmaly distributed
        * assumoption: network looks same no matter where you are
* V_peS_pe = S_pe / k
* assertion: "system is balanced when V_peS_pe = V_clS_cl"
* S_pe / S_cl = V_cl / V_pe
  * service time on process / service time on communication link
  * this is similar to granularity ratio
* S_pe / S_cl = V_cl / V_pe  = K*V_cl
  * computational ratio (or granularity) must be ? if ? is not going to dominate


### some computations
l_max = network diameter
LV = mean number of links crossed
phi(l) = probability of message crossing exactly l links
V_cl = link message intensity
LV = sum(l=1 -> l=l_max, l*phi(l))
V_cl = LV / num links
