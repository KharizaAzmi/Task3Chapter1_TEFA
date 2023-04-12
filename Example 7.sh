#!/bin/bash

foods=('Sate' 'Bakso' 'Dimsum' 'Rames')
drinks=('Jeruk' 'Teh' 'Kelapa' 'Cendol')
counter=0

for i in "${foods[@]}"; do
    for j in "${drinks[@]}"; do
        ((counter++))
        echo "Menu $counter: $i dan $j"
    done
done