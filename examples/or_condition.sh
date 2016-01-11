#!/bin/bash

read X

if [ $X = "y" ] || [ $X = "Y" ]
then
    echo "YES"
else
    echo "NO"
fi