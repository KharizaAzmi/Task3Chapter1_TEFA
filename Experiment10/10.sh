#!/bin/bash

telco=("Telkom" "Indosat" "XL" "Verizon" "AT&T" "Nippon" "Vodafone" "Orange" "KDDI" "Telefonica" "T-Mobile")

function findCompany {
    input=$1
    for (( i=0; i<${#telco[@]}; i++ )); do
        if [ "$i" -eq "$input" ]; then
            echo "Company Found: ${telco[$input]}"
            break
        fi
        echo "Searching Company... $((i+1))"
    done
}

company=$((RANDOM % 11))
findCompany "$company"
