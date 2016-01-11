#!/bin/bash

X=1
Y=2

if [ $X -lt $Y ]
then
	echo "hi"
fi

if test $X -lt $Y
then
	echo "hi"
fi

