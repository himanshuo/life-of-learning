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

Can you regular expression such as * to select multiple files. For example, you can do "mget file*" which means get all files that start with "file"

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
| RETR  | filename(s)  | files to receive |
| STOR | filename(s) | store files |
| STRU | F,R,or P | file, record, or page |
| TYPE | A,E,or I | ascii, EBCDIC, image |
| MODE | S, B, C  | string, block,  |


A given block DIAGRAM
