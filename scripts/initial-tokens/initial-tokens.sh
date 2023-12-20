#!/bin/bash

# Function to send POST requests
send_post_request() {
    local endpoint="$1"
    local data="$2"
    curl -X POST -H "Content-Type: application/json" -d "$data" "$endpoint/tokens"
}

# Check if curl is installed
if ! command -v curl &> /dev/null; then
    echo "curl is required but it's not installed. Please install curl."
    exit 1
fi

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "jq is required but it's not installed. Please install jq."
    exit 1
fi

# Parsing command line arguments
while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in
        --endpoint)
        endpoint="$2"
        shift
        shift
        ;;
        --tokens)
        tokens="$2"
        shift
        shift
        ;;
        *)
        shift
        ;;
    esac
done

# Validation for required arguments
if [ -z "$endpoint" ] || [ -z "$tokens" ]; then
    echo "Usage: $0 --endpoint <Census3 endpoint> --tokens <JSON file>"
    exit 1
fi

# Reading and processing the JSON file
token_data=$(jq -c '.tokens[]' "$tokens")
while read -r line; do
    echo "Sending POST request with body: $line"
    send_post_request "$endpoint" "$line"
done <<< "$token_data"
