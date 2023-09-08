#!/usr/bin/env sh

set -euo pipefail

readonly CONTRACT_ADDRESS="0xFE67A4450907459c3e1FFf623aA927dD4e28c67a"
readonly CONTRACT_TYPE="erc20"
readonly API_ENDPOINT="127.0.0.1:7788/api"

create_token() {
    curl -X POST \
        --json "{\"id\": \"$CONTRACT_ADDRESS\",\"type\": \"$CONTRACT_TYPE\",\"chainID\": 1}" \
        http://$API_ENDPOINT/token
}

get_token() {
    curl -X GET \
        http://$API_ENDPOINT/token/$CONTRACT_ADDRESS
}

create_census() {
    curl -X POST \
        --json "{\"strategyId\": $1,\"anonymous\": true}" \
        http://$API_ENDPOINT/census
}

get_census() {
    curl -X GET \
        http://$API_ENDPOINT/census/queue/$1
}

main() {
    # create the token
    echo "-> creating token..."
    create_token
    # wait to be synced (can be checked getting token info)
    echo "-> created, waiting 4m to token scan"
    sleep 240
    # get token info after some scan and store the 'defaultStrategy' id.
    echo "-> getting token info..."
    get_token
    # create the census with the token 'defaultStrategy' id and store the 
    # resulting 'queueId'.
    echo "-> enter the strategyId:"
    read strategyId
    echo "-> creating census..."
    create_census $strategyId
    echo "-> waiting 1m to census publication"
    sleep 60
    # get the enqueue census creation process with the 'queueId'
    echo "-> enter the enqueue census:"
    read queueId
    get_census $queueId
}

main "$@"
