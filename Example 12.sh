#!/bin/bash

generateData() {
    baseNumber="082"
    customers=()

    for ((i=1; i<=n; i++)); do
        mobileNumber="$baseNumber$(printf '%09d' $i)"
        customers+=( "$mobileNumber" )
    done

    echo "${customers[@]}"
}

sendPromoDiscount() {
    for customer in "$@"; do
        echo "Sending Promo to $customer"
        echo "Sending Discount to $customer"
    done

    echo "Finished sending Promo and Discount to ${#input[@]} Customers"
}

n=1000
input=($(generateData))
sendPromoDiscount "${input[@]}"