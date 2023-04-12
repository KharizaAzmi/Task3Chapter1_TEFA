#!/bin/bash

COMPANY = "telkom"
LARGECOMPANYNAME[]
for i in $(COMPANY); do
    LARGECOMPANYNAME[]+=($i)
done

findCompany(ARRAY){
    for i in {1..${#array[@]}}
    do
       if ARRAY[i] == "telkom"
        echo "Company Found!" 
    done

    THETIME=$({ time "$@" ; } 1> /dev/null | awk '{ print }');
}

findCompany(LARGECOMPANYNAME)