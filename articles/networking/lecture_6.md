Lecture 6
==========

### MIME (multi purpose mail extensions)
used for sending images, videos, and non-text files

Has its own headers:
* MIME-version: 1.1
  * current version is 1.1
* Content-Type: type/subtype
* Content-transfer-encoding:
* Content-id:
* Content-description:
  * explanation of all information given

type of encoding:
* base64 - groups 6 bits at a time and turns it into an ascii character
* quoted-printable


Type/Subtypes  
* text
  * subtypes are html, plain
* multipart
  * subtypes are mixed
    * mixed is when you have text followed by image followed by audio...
  * you have to
* image
* video


Content Transfer Encoding

* 7-bit NVT ASCII code
* 8-bit ascii extension
* binary
* base64 - primarily used for images, video
* quoted printable


2^6 = 64 = ?.63

1100110010000000011100


62 comes from 26 upper case letters, 26 lower case letters, 10 digits, +, and /.

is 61, + is 62

representing 3 bytes will take 8 bytes so there is 25% redundancy


*quoted-printable*   
decimal value range from 23-126 in ascii for printable items.
If numbers are in between 23-126, then don't do any conversion.  
0-23 and 126-255 are supposed to be replaced with =byte1byte2. Thus each byte is replaced with 3 bytes(1 byte for =, 2 more bytes for conversion )
Note that if you are converting

examples:   
* 00100110
    * in decimal this is 38.
    * 38 decimal is ascii &
* 01001100
    * 76 decimal is ascii
* 10011101
  * 157 decimal
  * 1001 is decimal 9
  * 1101 is decimal D
  * you convert the 9 into ascii 9
  * you convert the D into ascii D
  * thus the resulting ascii will be ????
* 00111001
  * 9
* 01001011
  * K


### Telnet
remote login
* enter login id and password

NVT (network virtual terminal) - converts local text


commands that telnet sends
* open - open a connection
* close - close a connection
* mode - change mode from line to character or character to line
* send - send data
* quit - quit session

### SSH (secure shell )
replaced telnet

2 versions - ssh 1 and ssh 2

forget about ssh 1. the two are not compatible

layers
* application
* ssh-conn
  * talking to correct server?
* ssh-auth
  * authorized
* ssh-tran
  * secure transmission
* tcp


SSH-Tran provides
* privacy and confidentiality
* data integrity
  * even if attacker sees data, attacker cannot change data
* server authentication
* compression

SSH-auth provides
* verifies users password/login id

SSH-conn provides
* makes connection between client and server

2 key things you can do once you ssh
* remote login
* data transfer

SSH Tunneling (port forwarding) - creates secure tunnel between server and client. FTP can be run on top of ssh which allows ftp to indirectly talk to another ftp server via a secure shell tunnel.
