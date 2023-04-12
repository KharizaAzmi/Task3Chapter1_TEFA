#!/bin/bash

company="telkom"
largeCompanyName=()

for (( i=0; i<10; i++ )); do
  largeCompanyName[$i]=$company
done

findCompany() {
  tx=$(date +%s%N)
  for i in "${array[@]}"; do
    if [[ "$i" == "$company" ]]; then
      echo "Company Found!"
    fi
  done
  ty=$(date +%s%N)
  echo "The performance is $((($ty-$tx)/1000000)) ms"
}

findCompany "${largeCompanyName[@]}"

#In this code, we use a for loop to initialize the largeCompanyName array and the ${#array[@]} notation to get the length of the array.

#We also use the date command to measure the execution time instead of the performance object in JavaScript or the time package in Go. The ${varname} notation is used to access the value of a variable in Bash.

#Finally, we define the findCompany function to loop through the array and check for the company name. The $((expression)) notation is used to perform arithmetic operations in Bash.