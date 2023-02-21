#!/usr/bin/env sh

set -euo pipefail

readonly CONTRACT_ADDRESS="0xa117000000f279D81A1D3cc75430fAA017FA5A2e"
readonly CONTRACT_TYPE="erc20"
readonly BLOCK_NUMBER="11060110"
readonly API_ENDPOINT="127.0.0.1:7788/api"

add_contract() {
    curl \
        --silent \
        --fail \
        --show-error \
        --request POST \
        --header 'Content-Type: application/json' \
        --url "${API_ENDPOINT}/addContract/${CONTRACT_ADDRESS}/${CONTRACT_TYPE}/${BLOCK_NUMBER}" \
        --output /dev/null
}

get_contract() {
    curl \
        --silent \
        --fail \
        --show-error \
        --request GET \
        --header 'Content-Type: application/json' \
        --url "${API_ENDPOINT}/getContract/${CONTRACT_ADDRESS}"
}

get_balances() {
    curl \
        --silent \
        --fail \
        --show-error \
        --request GET \
        --header 'Content-Type: application/json' \
        --url "${API_ENDPOINT}/balances/${CONTRACT_ADDRESS}"
}

main() {
    add_contract
    sleep 10
    get_contract
    sleep 10
    get_balances
}

main "$@"
