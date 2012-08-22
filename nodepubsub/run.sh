#!/bin/bash

for i in {1..300} #number of channels
    do 
        node simple.js $i &
        sleep 1
    done 
