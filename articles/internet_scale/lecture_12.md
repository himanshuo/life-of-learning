Lecture 12 - Queuing
==========

### MS research paper
Bing created a bunch of features that were meant to affect search results.

Then, they measured how much each feature had an impact on user's happiness of search results.

They ended up removing ~90% of the features they created because only a few really mattered to users.

### overview
queues are ordered (usually persistent) lists of messages

when made accessible between application they enable
  * asynchronous processing
  * decoupled systems
  * data flow oriented APIs between systems

### Asynchronous Processing
Asynchronous processing just means that the function does not block.

A few examples of things you want to make asynchronous are
  * slow operations
  * things that could fail
  * image resizing, virus scanning, format conversion
  * logging and tracking
  * search index updates
  * making calls to 3rd party APIs

You want to make all these things asynchronous so that
  * user does not have to wait for these things
  * server thread does not have to wait


##### asynchronous processing with queuing
new_listing creates a new list item asynchronously.

    def new_listing(request):
      # create item ...
      ...
      public_msg({
        "type":"new_item",
        "title":item.title,
        "id":item.id
        })

update_search just listens for all incoming requests. If one of the requests come in, then the update_search function will get it via get_msg and run the code. update_search is the queuing aspect of this.


    def update_search():
      while(True):
        i = get_msg("new_item")
        search_index(i.title, i.id)



#### Decoupled systems examples
there are numerous different systems that allow you to create a new item by the user. However, each of these must call new item trigger.

The problem that can occur is that humans have to make sure they follow each of the steps as per the 'contract' that the trigger function has.

For example, create item might require you to do a bunch of things, before you call the createitem function, but the coder might not do so.

Another issue that can occur is that the order of the tasks may matter internally but not logically to the coder. Thus the coder might call the things out of order.


### Queuing Architectures
![](lecture_12-images/a31ffd836c7e7d3ee7a9626dc4dc4188.png)

* the left side is the user actions
* the right side is the functions to call
* the things in the middle are queues
  * the queue is represented by the squares
  * the cyclinder means that you write each request that comes in into the disk. This way if the queue fails, you can still run the events in some future time.
  * there are multiple ItemClicks to reveal different shards. Having multiple shards allows for load balancing of the queues
  * there are copies of each event in all shards in order to increase resiliency
  * you can make it so that something does not get processed in x amount of time, kill it. But then for more important things, you kill it if it does not run in x+y time. Thus importance of things can be handled by queuing systems.

### Summary of Queuing systems
message passing is another form of API. message passing is asynchronous.

push based application design. publisher triggers events that cause subscribers to take action.
  * publisher doesn't know who created the trigger. It just handles it.
  * this is another form of modularity/isolation

supports fan-out scaling
  * n publishers, m consumers.

persistence supports
  * failure recovery

scheduling when to run consumers (maybe at night when load otherwise low)


### message queueing variations
* publish/subcribe systems
  * publishers trigger event and subscribers receive it asynchronously
* exactly once delivery
  * every subscriber receives every message exactly once
 * usually implemented via broker/intermediary between publisher and subscribers
 * persistent or non-persistent queues

### implementations
Steps that the publisher takes:
  1. user publishs message M
  2. find shard that M should go to
  3. write M to all replicas in shard at end of a message log on all replicas
  4.acknolwedge receipt to publisher

Steps that the consumer/subscriber takes:
  1. subscriber x requests next message
  2. pick shard S to read message from
  3. lookup next message location L for subscriber X on shard S
  4. pick a replica R in shard S to read from
  5. lookup message at location L in replica R
  6. return message to subscriber X
  7. increment next message location for X on shard S
    * message location is stored in client/consumer so that

Key point: nothing gets deleted from the queue. It is just stored in the database. This allows each different type of consumer to have its own position in the queue to handle events. 
