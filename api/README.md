# API endpoints

+ **GET** `/tokens` (SDK OK, API OK*)
List of already added tokens.

    - 📥 response:
    
    ```json
    {
        tokens: [
            {
                "id": "0x1234",
                "name": "wANT",
                "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
                "startBlock": 123456
            }
        ]
    }
    ```
    
    - ⚠️ errors:
    
    | HTTP Status  | Message | Internal error |
    |:---:|:---|:---:|
    | 204 | `-` | 4007 |
    | 500 | `error getting tokens information` | 5005 | 
    | 500 | `error encoding tokens` | 5011 | 

+ **GET** `/tokens/types` (SDK OK, API OK)
List the supported token types.

    - 📥 response:
    
    ```json
    {
        supportedTokens: [
            "erc20", "erc721", "erc777", 
            "erc1155", "nation3", "wANT"
        ]
    }
    ```
    
    - ⚠️ errors:    
    
    | HTTP Status  | Message | Internal error |
    |:---:|:---|:---:|
    | 500 | `error encoding supported tokens types` | 5012 | 

+ **POST** `/tokens` (API OK)
Triggers a new scan for the provided token, starting from the defined block.

    - 📤 request:
    
    ```json
    {
        "id": "0x1234",
        "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
        "startBlock": 123456
    }
    ```
    
    - ⚠️ errors:
    
    | HTTP Status  | Message | Internal error |
    |:---:|:---|:---:|
    | 400 | `malformed token information` | 4000 | 
    | 500 | `the token cannot be created` | 5000 | 
    | 500 | `error getting token information` | 5004 | 
    | 500 | `error initialising web3 client` | 5018 | 

+ **GET** `/tokens/{tokenID}` (SDK OK, API OK*)
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
        "name": "Amazing token"
        "status": {
            "atBlock": 12345,
            "synced": true|false,
            "progress": 87
        },
        "defaultStrategy": 1,
    }
    ```
    
    - ⚠️ errors:
    
    | HTTP Status  | Message | Internal error |
    |:---:|:---|:---:|
    | 404 | `no token found` | 4003 |
    | 500 | `error getting token information` | 5004 | 
    | 500 | `error encoding tokens` | 5011 | 
    
    **MVP Warn**: If `defaultStrategy` is `0`, no strategy (neither the dummy strategy) is associated to the given token.

+ **POST** `/strategies`
Stores a new strategy based on the defined combination of tokens provided, these tokens must be registered previously.

    - 📤 request:
    
    ```json
     {
        "tokens": [
          {
              "id": "0x1324",
              "name": "wANT"
              "minBalance": "10000",
              "method": "0x8230" 
          },
          {
              "id": "0x5678"
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

+ **GET** `/strategies` (SDK OK, API OK)
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

+ **GET** `/strategies/{strategyId}` (SDK OK, API OK*)
Returns the information of the strategy related to the provided ID.

    - 📥 response:
    
    ```json
    {
        "id": 2
        "tokens": [
          {
              "id": "0x1324",
              "name": "wANT"
              "minBalance": "10000",
              "method": "0x8230" 
          },
          {
              "id": "0x5678"
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

+ **GET** `/strategies/token/{tokenID}` (SDK OK, API OK)
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

+ **POST** `/census` (API OK)
Request the creation of a new census with the strategy provided for the `blockNumber` provided and returns the census ID.
     
     - 📤 request:
    
    ```json
    {
        "strategyId": 1,
        "blockNumber": 123456
    }
    ```
    
    - 📥 response:
    
    ```json
      {
        "censusId": 12
      }
    ```
    
    - ⚠️ errors :
    
    | HTTP Status  | Message | Internal error |
    |:---:|:---|:---:|
    | 400 | `malformed strategy ID, it must be a integer` | 4002 | 
    | 404 | `no strategy found with the ID provided` | 4005 | 
    | 500 | `error creating the census tree on the census database` | 5001 | 
    | 500 | `error encoding strategy holders` | 5014 | 


+ **GET** `/census/strategy/{strategyId}` (API OK)
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
    | 500 | `error encoding cenuses` | 5018 |

+ **GET** `/census/{censusId}` (API OK)
Returns the information of the snapshots related to the provided ID.

    - 📥 response:
    ```json
    { 
      "id": 2,
      "strategyId": 1,
      "merkleRoot": "e3cb8941e25dcdb36fc21acbe5f6c5a42e0d4f89839ae94952f0ebbd9acd04ac"
      "uri": "ipfs://Qma...."
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
