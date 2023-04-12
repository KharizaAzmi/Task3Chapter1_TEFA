#!/bin/bash

candidates=("Baswedan" "Subianto" "Maharani")

# define a recursive function to generate permutations
arrange() {
    local items=("$@")
    # base case: if only one candidate, return it as a bash array
    if (( ${#items[@]} == 1 )); then
        echo "${items[0]}"
        return
    fi
    local results=()
    # loop through all candidates and make them the first element of the permutation
    for i in "${!items[@]}"; do
        candidate="${items[$i]}"
        # create a new array with the candidate removed
        current=("${items[@]:0:$i}" "${items[@]:$((i+1))}")
        # recursive call to generate all permutations of the remaining candidates
        memory=($(arrange "${current[@]}"))
        # append the current candidate to the beginning of each sub-permutation
        for sub_perm in "${memory[@]}"; do
            results+=("$candidate $sub_perm")
        done
    done
    # return the results as a bash array
    echo "${results[@]}"
}

# call the permutation function with the candidates array
arrange=($(arrange "${candidates[@]}"))

# print all permutations
for perm in "${arrange[@]}"; do
    echo "$perm"
done
