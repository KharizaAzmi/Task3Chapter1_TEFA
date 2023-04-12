#!/bin/bash

address='DKI Jakarta'
addresses=("$address" "$address" "$address" "$address" "$address" "$address" "$address" "$address" "$address" "$address")

function findAddress {
    tx=$(date +%s%N | cut -b1-13)
    echo "The Default Address is ${addresses[0]}"
    ty=$(date +%s%N | cut -b1-13)
    echo "The Performance is $(expr $ty - $tx) ms"
}

findAddress "${addresses[@]}"