#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <API_base_URL>"
    exit 1
fi

api_base_url=$1

# Function to get token size
get_token_size() {
    local token_id="$1"
    local chainID="$2"
    local externalID="$3"

    local url="$api_base_url/tokens/$token_id?chainID=$chainID"
    if [ "$externalID" != "null" ]; then
        url="$url&externalID=$externalID"
    fi

    local response=$(curl -s "$url")
    local token_size=$(echo "$response" | jq -r '.size')

    echo "$token_size"
}

# Function to request tokens and create census
request_census() {
    local token_id="$1"
    local token_name="$2"
    local default_strategy="$3"
    local chainID="$4"
    local externalID="$5"

    # Get token size
    local token_size=$(get_token_size "$token_id" "$chainID" "$externalID")

    # Request census creation
    local response=$(curl -s -X POST -H "Content-Type: application/json" -d "{\"strategyId\": $default_strategy, \"anonymous\": true}" "$api_base_url/censuses")

    # Extract queueID
    local queueID=$(echo "$response" | jq -r '.queueID')

    # Get start time
    local start_time=$(date +%s)

    # Long polling
    local done=false
    echo "Waiting for census creation for token: $token_name, Size: $token_size"
    while [ "$done" != true ]; do
        response=$(curl -s "$api_base_url/censuses/queue/$queueID")
        done=$(echo "$response" | jq -r '.done')
        if [ "$done" = true ]; then
            # Get end time
            local end_time=$(date +%s)
            local elapsed_time=$((end_time - start_time))
            echo "Census created for token: $token_name, Size: $token_size, Time taken: $elapsed_time seconds"
            echo "$token_size,$elapsed_time" >>census_times.csv
        fi
        sleep 1 # Long polling interval
    done
}

# Request tokens
tokens_response=$(curl -s "$api_base_url/tokens?pageSize=-1")
tokens=$(echo "$tokens_response" | jq -r '.tokens[] | "\(.ID),\(.name),\(.defaultStrategy),\(.chainID),\(.externalID)"')

# Header for CSV
echo "Size,Time (s)" >census_times.csv

# Iterate over tokens
while IFS=',' read -r token_id token_name default_strategy chainID externalID; do
    echo "Processing token ID: $token_id, Name: $token_name"
    request_census "$token_id" "$token_name" "$default_strategy" "$chainID" "$externalID"
done <<< "$tokens"

echo "Script execution completed"
