#!/bin/bash

function calculateBigO() {
    l=5
    p=8
    k=$((l * p))

    for i in "${input[@]}"; do
        addInput
        calculateNumber
    done

    echo "Output: true"
}