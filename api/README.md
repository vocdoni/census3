# API endpoints

Endpoints:
 - [API info](#api-info)
 - [Tokens](#tokens)
 - [Strategies](#strategies)
 - [Censuses](#censuses)

## API Info

### GET `/info`

Show information about the API service.

- 📥 response:

```json
{
    "supportedChains": [
        {
            "chainID": 5,
            "shortName": "gor",
            "name": "Goerli"
        },
        {
            "chainID": 137,
            "shortName": "matic",
            "name": "Polygon Mainnet"
        },
        {
            "chainID": 80001,
            "shortName": "maticmum",
            "name": "Mumbai"
        },
        {
            "chainID": 1,
            "shortName": "eth",
            "name": "Ethereum Mainnet"
        }
    ]
}
```

- ⚠️ errors:

| HTTP Status | Message | Internal error |
|:---:|:---|:---:|
| 500 | `error encoding API info` | 5023 | 

## Tokens

### GET `/tokens`
List of already added tokens.

- 📥 response:

```json
{
    "tokens": [
        {
            "ID": "0x1234",
            "name": "Wrapped Aragon Network Token",
            "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
            "startBlock": 123456,
            "symbol": "wANT",
            "tags": "testTag1,testTag2",
            "chainID": 1
        }
    ]
}
```

> If `tags` is empty, it will be ommited.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `no tokens found` | 4007 |
| 500 | `error getting tokens information` | 5005 | 
| 500 | `error encoding tokens` | 5011 | 

### GET `/tokens/types`
List the supported token types.

- 📥 response:

```json
{
    "supportedTypes": [
        "erc20", "erc721", "erc777", 
        "erc1155", "nation3", "wANT"
    ]
}
```

- ⚠️ errors:    

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 500 | `error encoding supported tokens types` | 5012 | 

### POST `/tokens`
Triggers a new scan for the provided token, starting from the defined block.

- 📤 request:

```json
{
    "ID": "0x1234",
    "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
    "tags": "testTag1,testTag2",
    "chainID": 1
}
```

> `tags` attribute is *optional*.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed token information` | 4000 | 
| 409 | `token already created` | 4009 | 
| 400 | `chain ID provided not supported` | 4013 | 
| 500 | `the token cannot be created` | 5000 | 
| 500 | `error getting token information` | 5004 | 
| 500 | `error initialising web3 client` | 5019 | 

### GET `/tokens/{tokenID}`
Returns the information about the token referenced by the provided ID.

- 📥 response:

```json
{
    "ID": "0x1324",
    "type": "erc20",
    "decimals": 18,
    "startBlock": 123456,
    "symbol": "$",
    "totalSupply": "21323",
    "name": "Amazing token",
    "status": {
        "atBlock": 12345,
        "synced": true|false,
        "progress": 87
    },
    "defaultStrategy": 1,
    "tags": "testTag1,testTag2",
    "chainID": 1
}
```

> If `tags` is empty, it will be ommited.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 404 | `no token found` | 4003 |
| 500 | `error getting token information` | 5004 | 
| 500 | `error encoding token` | 5010 | 
| 500 | `chain ID provided not supported` | 5013 | 
| 500 | `error initialising web3 client` | 5019 | 
| 500 | `error getting number of token holders` | 5020 | 
| 500 | `error getting last block number from web3 endpoint` | 5021 | 

**MVP Warn**: If `defaultStrategy` is `0`, no strategy (neither the dummy strategy) is associated to the given token.

## Strategies

### GET `/strategies`
Returns the ID's list of the strategies registered.

- 📥 response:

```json
{
    "strategies": [ 1, 3 ]
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4008 |
| 500 | `error getting strategies information` | 5008 | 
| 500 | `error encoding strategies` | 5016 | 

### POST `/strategies`
Stores a new strategy based on the defined combination of tokens provided, these tokens must be registered previously.

- 📤 request:

```json
    {
        "alias": "test_strategy",
        "predicate": "(wANT OR ANT) AND USDC",
        "tokens": {
            "wANT": {
                "ID": "0x1324",
                "chainID": 1,
                "minBalance": "10000"
            },
            "ANT": {
                "ID": "0x1324",
                "chainID": 5,
            },
            "USDC": {
                "ID": "0x1324",
                "chainID": 1,
                "minBalance": "50"
            },
        }
    }
```

- 📥 response:

```json
{
    "strategyID": 1
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 404 | `no token found` | 4003 | 
| 400 | `malformed strategy provided` | 4014 |
| 400 | `the predicate provided is not valid` | 4015 | 
| 400 | `the predicate includes tokens that are not included in the request` | 4016 | 
| 500 | `error encoding strategy info` | 5015 | 
| 500 | `error creating strategy` | 5025 | 


### GET `/strategies/{strategyID}`
Returns the information of the strategy related to the provided ID.

- 📥 response:

```json
{
    "ID": 1,
    "alias": "strategy_alias",
    "predicate": "MON AND (ANT OR USDC)",
    "tokens": {
        "MON": {
            "ID": "0x1234",
            "chainID": 5
        },
        "ANT": {
            "ID": "0x1234",
            "chainID": 1,
            "minBalance": "1"
        },
        "USDC": {
            "ID": "0x1234",
            "chainID": 1,
        }
    }
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed strategy ID, it must be an integer` | 4002 | 
| 404 | `no strategy found with the ID provided` | 405 |
| 500 | `error getting tokens information` | 5005 | 
| 500 | `error getting strategy information` | 5007 | 
| 500 | `error encoding strategy info` | 5015 | 

### GET `/strategies/token/{tokenID}`
Returns ID's of the already created strategies including the `tokenAddress` provided.

- 📥 response:

```json
{
    "strategies": [ 2, 8 ]
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4008 |
| 500 | `error getting strategies information` | 5008 | 
| 500 | `error encoding strategies` | 5016 | 

### POST `/strategies/predicate/validate`
Returns if the provided strategy predicate is valid and well-formatted. If the predicate is valid the handler returns a parsed version of the predicate as a JSON.

- 📤 request:

```json
{
    "predicate": "DAI AND (ANT OR ETH)"
}
```

- 📥 response:

```json
{
    "result": {
        "childs": {
            "operator": "AND",
            "tokens": [
                {
                    "literal": "DAI"
                },
                {
                    "childs": {
                        "operator": "OR",
                        "tokens": [
                            {
                                "literal": "ANT"
                            },
                            {
                                "literal": "ETH"
                            }
                        ]
                    }
                }
            ]
        }
    }
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed strategy provided` | 4014 |
| 400 | `the predicate provided is not valid` | 4015 | 
| 500 | `error encoding validated strategy predicate` | 5024 | 

## Censuses

### POST `/censuses`
Request the creation of a new census with the strategy provided for the `blockNumber` provided and returns the census ID.
     
- 📤 request:

```json
{
    "strategyID": 1,
    "blockNumber": 123456,
    "anonymous": false
}
```

- 📥 response:

```json
{
    "queueID": "0123456789abcdef0123456789abcdef01234567"
}
```

- ⚠️ errors :

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed strategy ID, it must be an integer` | 4002 | 
| 500 | `error encoding census` | 5017 | 

### GET `/censuses/{censusID}`
Returns the information of the snapshots related to the provided ID.

- 📥 response:
```json
{ 
    "censusID": 2,
    "strategyID": 1,
    "merkleRoot": "e3cb8941e25dcdb36fc21acbe5f6c5a42e0d4f89839ae94952f0ebbd9acd04ac",
    "uri": "ipfs://Qma....",
    "size": 1000,
    "weight": "200000000000000000000",
    "anonymous": true
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed census ID, it must be a integer` | 4001 | 
| 404 | `census not found` | 4006 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding census` | 5017 | 

### GET `/census/queue/{queueID}`
Returns the information of the census that are in the creation queue.

- 📥 response:
```json
{
    "done": true,
    "error": {
        "code": 0,
        "err": "error message or null"
    },
    "census": { /* <same_get_census_response> */ }
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 404 | `census not found` | 4006 | 
| 400 | `malformed queue ID` | 4011 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding census queue item` | 5022 | 

- ⚠️ possible error values inside the body:

<small>The request could response `OK 200` and at the same time includes an error because it is an error of the enqueued process and not of the request processing).</small>

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 404 | `no token holders found` | 4004 |
| 404 | `no strategy found with the ID provided` | 4005 |
| 400 | `no tokens found for the strategy provided` | 4010 |
| 409 | `census already exists` | 4012 |
| 400 | `the predicate provided is not valid` | 4015 |
| 204 | `strategy has not registered holders` | 4017 |
| 500 | `error creating the census tree on the census database` | 5001 |
| 500 | `error evaluating strategy predicate` | 5026 |

### GET `/census/strategy/{strategyID}`
Returns a list of censusID for the strategy provided.

- 📥 response:

```json
{
    "censuses": [ 3, 5 ]
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4007 |
| 400 | `malformed census ID, it must be a integer` | 4001 | 
| 404 | `census not found` | 4006 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding censuses` | 5018 |