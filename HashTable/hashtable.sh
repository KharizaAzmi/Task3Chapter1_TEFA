#!/bin/bash

declare -A hashtable

function _hash() {
  local key=$1
  local hash=0
  local i=0

  for ((i = 0; i < ${#key}; i++)); do
    hash=$((hash + $(printf '%d' "'${key:i:1}") * i))
  done

  echo $((hash % 100))
}

function set() {
  local key=$1
  local value=$2
  local address=$(_hash "$key")

  if [[ -z "${hashtable[$address]}" ]]; then
    hashtable[$address]="$key:$value"
  else
    hashtable[$address]="${hashtable[$address]}|$key:$value"
  fi
}

function get() {
  local key=$1
  local address=$(_hash "$key")

  if [[ -z "${hashtable[$address]}" ]]; then
    echo ""
    return
  fi

  local pairs=(${hashtable[$address]})

  for pair in "${pairs[@]}"; do
    local k=$(echo "$pair" | cut -d':' -f1)
    local v=$(echo "$pair" | cut -d':' -f2)

    if [[ "$k" == "$key" ]]; then
      echo "$v"
      return
    fi
  done

  echo ""
}

set "082124606606" "Rony Setyawan"
echo $(get "082124606606")

