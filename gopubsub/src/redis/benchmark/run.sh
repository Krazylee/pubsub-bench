#!/bin/bash

for i in {1..500} #number of channels
    do 
        ./benchmark $i &
        sleep 1
    done 
