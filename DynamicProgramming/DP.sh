#!/bin/bash

counterPlainRecursive=0
counterDynamicProgramming=0

function fibonacciPlainRecursive() {
    ((counterPlainRecursive++))
    if [ $1 -lt 2 ]; then
        echo $1
    else
        echo $(( $(fibonacciPlainRecursive $(( $1 - 1 ))) + $(fibonacciPlainRecursive $(( $1 - 2 ))) ))
    fi
}

function fibonacciDynamicProgramming() {
    ((counterDynamicProgramming++))
    if [ $1 -lt 2 ]; then
        echo $1
    else
        cache=( 0 1 )
        for (( i=2; i<=$1; i++ )); do
            cache[$i]=$(( ${cache[$(( $i - 1 ))]} + ${cache[$(( $i - 2 ))]} ))
        done
        echo ${cache[$1]}
    fi
}

fibonacciPlainRecursive 20
fibonacciDynamicProgramming 20

echo "we did $counterPlainRecursive calculation for Plain Recursive"
echo "we did $counterDynamicProgramming calculation for Dynamic Programming"
