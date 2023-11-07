import argparse
import json
import requests

parser = argparse.ArgumentParser(description='Set intial tokens into census3 endpoint.')
parser.add_argument('--endpoint', type=str, help='Census3 endpoint to set initial tokens into.')
parser.add_argument('--tokens', type=str, help='JSON file containing initial tokens to set into census3 endpoint. It must follow \'{"tokens": [ { "id": "0x1234", "type": "erc20", "chainID": 1 } ]}\'')
args = parser.parse_args()

def main():
    with open(args.tokens, "r") as f:
        initial_tokens = json.load(f)

        for token in initial_tokens['tokens']:
            requests.post(args.endpoint + "/tokens", json=token)

if __name__ == "__main__":
    main()