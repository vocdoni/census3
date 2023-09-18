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
    "chainIDs": [1, 5]
}
```

- ⚠️ errors:

| HTTP Status | Message | Internal error |
|:---:|:---|:---:|
| 500 | `error encoding API info` | 5023 | 

## Tokens

### GET `/token`
List of already added tokens.

- 📥 response:

```json
{
    "tokens": [
        {
            "id": "0x1234",
            "name": "Wrapped Aragon Network Token",
            "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
            "startBlock": 123456,
            "symbol": "wANT",
            "tag": "testTag1,testTag2"
        }
    ]
}
```

> If `tag` is empty, it will be ommited.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `no tokens found` | 4007 |
| 500 | `error getting tokens information` | 5005 | 
| 500 | `error encoding tokens` | 5011 | 

### GET `/token/types`
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

### POST `/token`
Triggers a new scan for the provided token, starting from the defined block.

- 📤 request:

```json
{
    "id": "0x1234",
    "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
    "tag": "testTag1,testTag2",
    "chainID": 1
}
```

> `tag` attribute is *optional*.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed token information` | 4000 | 
| 409 | `token already created` | 4009 | 
| 400 | `chain ID provided not supported` | 4013 | 
| 500 | `the token cannot be created` | 5000 | 
| 500 | `error getting token information` | 5004 | 
| 500 | `error initialising web3 client` | 5019 | 

### GET `/token/{tokenID}`
Returns the information about the token referenced by the provided ID.

- 📥 response:

```json
{
    "id": "0x1324",
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
    "tag": "testTag1,testTag2",
    "chainID": 1
}
```

> If `tag` is empty, it will be ommited.

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 404 | `no token found` | 4003 |
| 500 | `error getting token information` | 5004 | 
| 500 | `error initialising web3 client` | 5018 | 
| 500 | `error getting last block number from web3 endpoint` | 5021 | 
| 500 | `error encoding tokens` | 5011 | 

**MVP Warn**: If `defaultStrategy` is `0`, no strategy (neither the dummy strategy) is associated to the given token.

## Strategies

### POST `/strategy`
Stores a new strategy based on the defined combination of tokens provided, these tokens must be registered previously.

- 📤 request:

```json
    {
    "tokens": [
        {
            "id": "0x1324",
            "name": "wANT",
            "minBalance": "10000",
            "method": "0x8230"
        },
        {
            "id": "0x5678",
            "name": "USDC",
            "minBalance": "20000",
            "method": "0x3241" 
        },
        {
            "id": "0x9da2",
            "name": "ANT",
            "minBalance": "1",
            "method": "0x9db1"
        }
    ],
    "strategy": "(wANT OR ANT) AND USDC"
    }
```

- 📥 response:

```json
{
    "strategyId": 1
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4008 |
| 500 | `error getting strategies information` | 5008 | 
| 500 | `error encoding strategies` | 5016 | 

### GET `/strategy`
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

### GET `/strategy/{strategyId}`
Returns the information of the strategy related to the provided ID.

- 📥 response:

```json
{
    "id": 2,
    "tokens": [
        {
            "id": "0x1324",
            "name": "wANT",
            "minBalance": "10000",
            "method": "0x8230" 
        },
        {
            "id": "0x5678",
            "name": "USDC",
            "minBalance": "20000",
            "method": "0x3241" 
        },
        {
            "id": "0x9da2",
            "name": "ANT",
            "minBalance": "1",
            "method": "0x9db1" 
        }
    ],
    "strategy": "(wANT OR ANT) AND USDC"
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4002 |
| 404 | `no strategy found with the ID provided` | 4005 | 
| 500 | `error getting tokens information` | 5005 | 
| 500 | `error getting strategy information` | 5007 | 
| 500 | `error encoding strategy` | 5015 | 

### GET `/strategy/token/{tokenID}`
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

### GET `/census/strategy/{strategyId}`
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
| 404 | `census not found` | 4006 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding censuses` | 5018 |

## Censuses

### POST `/census`
Request the creation of a new census with the strategy provided for the `blockNumber` provided and returns the census ID.
     
- 📤 request:

```json
{
    "strategyId": 1,
    "blockNumber": 123456,
    "anonymous": false
}
```

- 📥 response:

```json
{
    "queueId": "0123456789abcdef0123456789abcdef01234567"
}
```

- ⚠️ errors :

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 400 | `malformed strategy ID, it must be an integer` | 4002 | 
| 500 | `error encoding strategy holders` | 5014 | 

### GET `/census/{censusId}`
Returns the information of the snapshots related to the provided ID.

- 📥 response:
```json
{ 
    "censusId": 2,
    "strategyId": 1,
    "merkleRoot": "e3cb8941e25dcdb36fc21acbe5f6c5a42e0d4f89839ae94952f0ebbd9acd04ac",
    "uri": "ipfs://Qma....",
    "size": 1000,
    "weight": "200000000000000000000",
    "chainId": 1,
    "anonymous": true
}
```

- ⚠️ errors:

| HTTP Status  | Message | Internal error |
|:---:|:---|:---:|
| 204 | `-` | 4007 |
| 400 | `malformed census ID, it must be a integer` | 4001 | 
| 404 | `census not found` | 4006 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding census` | 5017 | 

### GET `/census/queue/{queueId}`
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
| 400 | `malformed queue ID` | 4010 | 
| 404 | `census not found` | 4006 | 
| 500 | `error getting census information` | 5009 | 
| 500 | `error encoding census queue item` | 5022 | 
