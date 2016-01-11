#!/bin/bash

color1="Blue"
color2="dark red"
color3="light green"

# Incorrect since this creates array [Blue, dark, red, light, green]
for X in $color1 $color2 $color3
do
	echo "$X"
done

# Correct since this creates array [Blue, dark red, light green]
for X in "$color1" "$color2" "$color3"
do
	echo "$X"
done
