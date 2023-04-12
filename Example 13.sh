#!/bin/bash

generateData() {
    for ((i=1; i<=$1; i++)); do
        printf "082%09d\n" $i
    done
}

sendPromoDiscount() {
    for i in ${input1[@]}; do
        echo "Sending Promo to $i"
    done
    echo "Its Finished to send Promo to ${#input1[@]} Customers"

    for i in ${input2[@]}; do
        echo "Sending Discount to $i"
    done
    echo "Its Finished to send Discount to ${#input2[@]} Customers"
}

dataPromo=($(generateData 100000))
dataDiscount=($(generateData 1000))

input1=("${dataPromo[@]}")
input2=("${dataDiscount[@]}")

sendPromoDiscount