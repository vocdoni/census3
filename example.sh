#!/usr/bin/env bash

set -euo pipefail

readonly CONTRACT_ADDRESS="0xFE67A4450907459c3e1FFf623aA927dD4e28c67a"
readonly CONTRACT_TYPE="erc20"
readonly API_ENDPOINT="127.0.0.1:7788/api"

create_token() {
    curl -X POST \
        --json "{\"ID\": \"$CONTRACT_ADDRESS\",\"type\": \"$CONTRACT_TYPE\",\"chainID\": 1}" \
        http://$API_ENDPOINT/tokens
}

get_token() {
    curl -X GET \
        http://$API_ENDPOINT/tokens/$CONTRACT_ADDRESS
}

create_census() {
    curl -X POST \
        --json "{\"strategyID\": $1,\"anonymous\": true}" \
        http://$API_ENDPOINT/censuses
}

get_census() {
    curl -X GET \
        http://$API_ENDPOINT/censuses/queue/$1
}

main() {
    # create the token
    echo "-> creating token..."
    create_token
    # wait to be synced (can be checked getting token info)
    echo "-> created, waiting 2m to token scan"
    sleep 120
    # get token info after some scan and store the 'defaultStrategy' id.
    echo "-> getting token info..."
    get_token
    # create the census with the token 'defaultStrategy' id and store the 
    # resulting 'queueId'.
    echo "-> enter the strategyID:"
    read strategyID
    echo "-> creating census..."
    create_census $strategyID
    echo "-> waiting 30s to census publication"
    sleep 30
    # get the enqueue census creation process with the 'queueId'
    echo "-> enter the enqueue census ID:"
    read queueId
    get_census $queueId
}

main "$@"
