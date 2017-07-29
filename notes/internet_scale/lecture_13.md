Lecture 13 - Users, Content, and Reputation
==========
### UGC
UGC = User Generated Content

For users, UGC is good because
  * perceived as less biased than editorial content
  * niche expertise
  * cheaper to produce than editorial content

For companies, UGC is good because
  * drives user engagement
  * unique asset
  * traffic from search engines and social sites


### 2-sided UGC sites
One-sided market place - one type of users.
  * target.com has only one type of user - shoppers

2-sided market place - 2 types of users
  * wikipedia.com has writers and readers

Most UGC sites have 2-sided split
  * usually have more readers than creators though
  * conceptually, 90-9-1 split between readers, typo fixers, true writers


### Why do people contribute to UGC sites
* financial
  * sellers on Ebay
  * spammers writing Amazon or Yelp
* demonstrate expertise
  * answering questions on StackOverflow to get job offers
* justice
  * writing bad review when you feel cheated
* accomplishment
  * write authoritative wikipedia page
* ideology, self-expression, identity

### challenges to UCG
* cold start - no one wants to write first review on site
  * fix: add UGC once you already have traffic
  * fix: hire initial contributors
  * fix: syndicate content from other sites
  * fix: copy
  * fix: figure out whether it is more important to have sellers or buyers in a 2-sided market
* defining cultural norms - what is and not acceptable
  * fix: define culture of what is and what is not acceptable
  * fix: explicit rules need to be written down
  * fix: implicit rules can be shown via example
    * there are lots of things that not explicitly written down as rules, but become norms for a website (ie. give funny pics to others in Habitica)
* weeding out bad actors
  * want to get rid of incompetent, commercially motivated, politically motivated people.
  * fix: manual review
    * staff or community
  * fix: reputation systems
    * every user is assigned reputation score
    * high reputation users get privileges
    * low reputation users get punishments
    * reputation also gives social status
    * this requires a system for assigning reputation
  * fix: automatic review
    * language models
      * topic modeling - is the topic of review a product topic or a spammy topic?
      * link blocking
      * look for repeated text or links across reviews
    * graph methods
      * PageRank-like flow algorithms for assigning reputation
        * my reputation flows to other people when I upvote their reviews
        * if I upvote X other reviews, their reputation is boosted by 1/X of my reputation
        * people gain reputation by having othe rhigh reputation usres endorse their work
      * looking for islands
      * looking for outliers
  * fix: product design to limit who can create content


### ranking  UGC
how do you determine which is the best response to question/best review for product/...?
  * voting
  * do not show flagged items
  * reputation / date
  * interestingness - number of comments, reputation, page views ...
