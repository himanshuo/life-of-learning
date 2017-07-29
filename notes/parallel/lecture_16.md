Lecture 16
==========


k = number of nodes

lmax = network diameter

LV = mean number of links crossed by each message = sum(l=1 -> lmax, l* phi(l))

phi(l) =  probability of message traveling exactly l hops

v_cl = link visit ratio = LV / num_links

V_peS_pe = V_clS_cl when balanced

### Bus
l_max = 1

phi(1) = 1

LV = sum(l=1 -> l=1, l * phi(l)) = sum(1 -> 1, 1*1) = 1

V_cl = LV /1 = 1/1 = 1

V_pe = 1/k (visit ratio)

S_pe / S_cl = V_cl / V_pe =  1 / (1/k) = k where k= number of nodes

Thus, as you add more nodes, you need to add more nodes.  

### ring (uniform)
nodes are arranged in a circular ring.

assume k=6. This means 6 links.

assume uniform routing

l_max= round_down(k/2)

phi(l) =  2/k

LV = sum(l=1->l=l_max, l * phi(l)) = sum(., l * 2/k) = 2/k * sum(., l) = 2/k * (k/2 * (k+1)/2) / 2 = k/4 + 1/2 ~ k/4 (don't care about constants)  

HINT: sum(l=0 -> l=N, i) = n * (n=1) / 2

V_cl = LV / num links = k/4 / k = 1/4
  * for every
  * not likely to scale because we have to go to more and more of the links


S_pe / S_cl = (1/4)/(1/k) = k/4
* as you increase ? , must increase granularity

Even though ring has a bunch more routing than bus, it doesn't make much of a difference. It is basically just a constant difference

S_pe = avg amount of time that it takes to ????? on a single communication link
  * compute / communication is the key thing we want. We use S_pe / S_cl to find this.


### sphere (sphere size =1)
with probability 1, you are only communicating with nearest neighbors.

sphere size 1 means that you only communicate with nearest neighbors

l_max = 1

phi(1) = 1

LV = sum(1->1, 1 * 1) = 1

V_cl = LV / num links = 1 / k

S_pe / S_cl = (1/k) / (1/k)  = 1

##### caveat



hw is nearest neighbor sphere size 1 problem

### torus
D-dimensional lattice of width w

each of the w^D nodes is connected in a ring of size w in D directions


num_links = D * w^D links
  * 4 wires, but each is connected to another

addresses are D digit, base w numbers (this is useful for routing purposes)

l_max = network diameter = D * round_down(w/2)
  * each ring has network diameter k/2
  * each ring is a 1 dimension.
  * in any given dimension, I can only go ???

assume uniform routing

we will computer phi(l) one dimension at a time

LV_one_dimension = w / 4   

LV = D * LV_one_dimension = D * w/4

V_cl = LV / num links = (D * w/4)  / (D * w*D) = (w/4) / (w^D) = 1/4 * w/w^D = 1/4 * (1/W^(D-1))

S_pe / S_cl = V_cl / V_pe = ( (1/4) * 1/W^(D-1)) / (1/W^D) = (1/w^(D-1)) / 4   * w^D = w/4


overall goal: given application topology, ????

if computation function is function of number of nodes, then you cannot scale.

w/4 is almost a constant so its pretty fast.

hypercubes were built entirely around this idea. hypercubes are toruses with w=2.
  * this is why we had binary addressing for hypercube
