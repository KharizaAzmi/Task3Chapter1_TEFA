#!/bin/bash

telco=("Telkom" "Indosat" "XL" "Verizon" "AT&T" "Nippon" "Vodafone" "Orange" "KDDI" "Telefonica" "T-Mobile")

function findCompany() {
    for ((i=0; i<${#telco[@]}; i++)); do
        if [ $i -eq $1 ]; then
            echo "Company Found : ${telco[$1]}"
            break
        fi

        echo "Searching Company...$((i+1))"
    done
}

company=$((RANDOM % 11))
findCompany $company