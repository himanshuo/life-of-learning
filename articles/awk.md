AWK
====

### Print 1st column from file

    awk '{print $1}' file

* inside the '{print ""}' you can tell awk to print stuff
* awk takes in stdin. If the stdin is a file, then its the contents of the file
* $1 refers to the 1st column of the input

### rename files to append a suffix

    ls folder | awk '{print "mv "$1" "$1".new"}' | sh

* you can run commands via printing that command and then piping it to sh
* $1 will here is a filename
* simply add a new ending to the filename
* the "" around $1 are there to handle spaces

### rename a file within the name

    ls folder | awk '{print "mv "$1" "$1}' | sed s/old/new/2 | sh

*


### syntax


- [ ] ok
- [ ] ok
- [ ] ok
- [ ] ok
