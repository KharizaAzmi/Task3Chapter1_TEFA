#!/bin/bash

arrange() {
    array=("$@")
    # base case: if only one candidate, return it as a slice of a slice
    if [ ${#array[@]} -eq 1 ]; then
        echo "[[${array[0]}]]"
        return
    fi

    results=()

    # loop through all candidates and make them the first element of the permutation
    for (( i=0; i<${#array[@]}; i++ )); do
        arr=${array[$i]}

        # create a new array with the candidate removed
        current=("${array[@]:0:i}" "${array[@]:i+1}")

        # recursive call to generate all permutations of the remaining candidates
        memory=$(arrange "${current[@]}")

        # append the current candidate to the beginning of each sub-permutation
        while read -r subPerm; do
            result=("$arr" "${subPerm[@]}")
            results+=("${result[@]}")
        done <<< "$memory"
    done

    # print the resulting permutations as a 2D slice
    printf '%s\n' "${results[@]}"
}

candidates=("Baswedan" "Subianto" "Maharani" "Ganjar")
chairs=$(arrange "${candidates[@]}")

# Print the header row
printf "%-15s %-15s %-15s %-15s\n" "Chair 1" "Chair 2" "Chair 3" "Chair 4"
echo "------------------------------------"

# Print each row of the 2D slice
while read -r row; do
    printf "%-15s %-15s %-15s %-15s\n" "${row[@]}"
done <<< "$chairs"
