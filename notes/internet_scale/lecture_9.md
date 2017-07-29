Lecture 9 - Multiscreen and multichannels
===========================================

### What is multi-screen / multi-channel?
User interact through many channels (eg. desktop, mobile...)

You have to make sure that they all coordinate with each other and have consistent data. This means things like making sure that buying a product on desktop looks the same as buying it on mobile.


### Mobile App limitations
You move all the logic of the app to the server. This way you only have to get the data from the server and serve it.

There can be
* bandwidth and latency issues -> customers hate slow apps
* app logic doesnt change much -> slow to innovate
* different experiences for iOS and android


### moving things to server side
You have to think about how to combine service calls, logic behind the app, and tracking the app.


There are layers of services that go into server side. One of the key services is a service to organize the system.

### web view
a web view is basically a view in a mobile app that just renders the online web page

you generally dont want to do this.


### Why still build website?

* no install required
* click from google search results straight to page
* advantages with web views


### How to adopt to screen sizes
80% of traffic comes from about 15 different screen sizes.

You should use response design(blocks that change with screen size) to design page.

### Challenges with email
* limited layout/CSS/javascript
* email clients layout html different from each other
* email is read on phones so there are lots of different sizes
* data changes between when email sent and when it is open (item is on sale when email is sent, but sold out by the time that email is opened)


You should use open time rendering - send email with responsive body that autofills itself when opened using latest data from server side. This is still powered by all the services.
