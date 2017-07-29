Lecture 10 - Security
=======================

Security - privacy; integrity; availability
Buffer overrun; SQL injection; Command injection
Cross Site Scripting (XSS) - steal cookie information; redirects the viewer to your script, which records their cookie to your log file, then redirects the viewer back to the unmodified search page so they don't know anything happened. Note that this injection will only work properly if you aren't actually modifying the page source on the server's end. Otherwise the unmodified page will actually be the modified page and you'll end up in an endless loop
Cookie stealing is when you insert a script into the page so that everyone that views the modified page inadvertently sends you their session cookie
Snooping - see network traffic; watering hole (attack sites where target will go to)
