Lecture 5
===============


### FTP (file transfer protocol)
copies remote file from client to a remote server

format: ftp:>get

methods:
* get
* put
* mget - get multiple files at once


steps:  
1. client makes request to server
2. server transfers file to you
3. connection is closed


It is more efficient to use HTTP than FTP to transfer files.

Can use regular expression such as * to select multiple files. For example, you can do "mget file*" which means get all files that start with "file"

FTP Client: DIAGRAM
user interface - control
control process -
data transfer protocol handles


Default ports:
* 20 - data transfer
* 21 - control process



Request Commands (made by client):  

|  Command  | Arguments  |  Description  |
|  ---  |  ---  |  ---  |
|  USER  |  user id  |  send user name  |
| PASS  | password  | send password |
| PORT  | port number  | client chooses a port number |
| TYPE | A,E,or I | ascii, EBCDIC, image |
| RETR  | filename(s)  | files to receive |
| STOR | filename(s) | store files |
| STRU | F,R,or P | file, record, or page |
| MODE | S, B, C  | string, block,  |


A given block DIAGRAM


Response Messages are 3-digit codes.
Examples:

|  Code  | Description |
|  ---  |  ---  |
|125|data connection already open |
|150|file status ok|
|200|command ok|
|220|???|
|230|login ok|
|331|user name is ok|
|425|cannot open data connection|
|500|syntax error (unrecognized command)|

If you get "1xy:posittve preliminary reply"  
If you get "2xy:positive completion reply"   
If you get "3xy:positive intermediate reply"   
If you get "4xy:transient negative completion reply"  
If you get "5xy:permanent negative reply"

ftp>get /usr/user/johnsmith/reports/file1
DIAGRAM

Example: code=125, description=data connection already open


### Email
SMTP (outgoing protocol for email)  
POP3 (incoming protocol for email)  
IMAP4 (incoming protocol for email - most popular)

DIAGRAM about mail user agent and the flow of emails
a key idea is that mails are pushed to your server until you pull them

Another key idea is that the messages can go either direction. A server can be a client, and a client can be server in a mail system.

2 types of MUA (mail user agent)
* command driven - Examples: elm, pine
* gui based - Examples: Outlook, Eudora(mac and windows), Exchange(windows)


MTA and MUA end up being combined together. This means the same client can both send and receive emails.

email address format: local@domain.com  


email sender/reciever DIAGRAM


Request Commands:

|  Command  | Arguments  |  Description  |
|  ---  |  ---  |  ---  |
|  HELO  |  senders hostname (ip address)  |  identify host name  |
| MAIL-FROM  | sender of message(email address)  | identify senders email address |
| RCPT-TO  | recipient email  | recipient email |

Response Codes:

|Code| Description |
|...|...|
|211|system status |
|220|service is ready|
|354|start mail input|

Phases to send email
1. establish connection (via transport layer)
2. data transfer
3. connection termination

### MIME (multi purpose mail extensions)
used for sending images, videos, and non-text files

Has its own headers:
* MIME-version: 1.1
* Content-Type: type/subtype
* Content-transfer-encoding:
* Content-id:
* Content-description: explanation of all information given

type of encoding:
* base64 - groups 6 bits at a time and turns it into an ascii character
* quoted-printable

### Default ports
* ssh : 22
* smtp: 25
* http: 80
* https:443
* ftp
  * data transfer : 20
  * control process : 21
