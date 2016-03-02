Lecture 4
============

To successfully complete a project, you need:
* design docs
  * user stories
  * data model
* services
  * It is good to have your services be modular.
  * Benefits of modularity:
    * independent FAILING 
    * independent Scaling
    * Abstracting implementation
    * delineating Team boundaries
* High level design considerations
  * consider building a service for each model in your system
  * strictly enforce only accessing data through services
* Implementation considerations
  * HTTP transport
    * leverages lots of existing infrastructure
    * alternative: ProtocolBuffers if you need low overhead
    * What are protocol buffers?
       * Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.
  * JSON response
    * JSON is efficient, especially when combined with gzip encoding.
    * xml is complicated and unneccessary nowadays
  * naming
    * /api/v1/users/43/tp
      * versioning - v1
      * unique - users/43
      * human parsable - tp
  * read versus update actions
    * GET for read, POST for update
  * APIs should NOT be deprecated so think about how you will come out with a v2, v3... using current architecture
  * call things concurrently as much as possible
    * future = Future(call(some_service))
  * handle failures when calling services
    * use timeouts on DNS lookup, on connect, and on read
    * retry request if sensible
    * add circuit breakers
      * after N failures/timeouts in S seconds, stop trying to make more requests to all related APIs
    * if non-mandatory services fail, then just do not show those portions of page instead of not loading the page at all
