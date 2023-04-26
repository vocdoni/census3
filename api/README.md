# API endpoints

+ **GET** `/tokens`
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

+ **GET** `/tokens/types` 
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

+ **POST** `/tokens`
Triggers a new scan for the provided token, starting from the defined block.
    - 📤 request:
    ```json
    {
        "address": "0x1234",
        "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
        "startBlock": 123456
    }
    ```

+ **GET** `/tokens/:id` 
Returns the information about the token referenced by the provided ID.
    - 📥 response:
    ```json
    {
        "address": "0x1324",
        "type": "erc20",
        "decimals": 18,
        "startBlock": 123456,
        "name": "Amazing token"
        "status": {
            "atBlock": 12345,
            "synced": true|false,
            "progress": 87
        }
    }
    ```

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

+ **GET** `/strategies`
Returns the ID's list of the strategies registered.
    - 📥 response:
    ```json
    {
        "strategies": [ 1, 3 ]
    }
    ```

+ **GET** `/strategies/{strategyId}`
Returns the information of the strategy related to the provided ID.
    - 📥 response:
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

+ **GET** `/strategies/token/{tokenAddress}`
Returns ID's of the already created strategies including the `tokenAddress` provided.
    - 📥 response:
    ```json
     {
         "strategies": [ 2, 8 ]
     }
    ```

+ **POST** `/census`
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

+ **GET** `/census/{strategyId}/`
Returns a list of censusID for the strategy provided.
    - 📥 response:
    ```json
        {
            "censuses": [ 3, 5 ]
        }
    ```

+ **GET** `/census/{censusId}`
Returns the information of the snapshots related to the provided ID.
    - 📥 response: Returns 200 or 204
    ```json
    { 
      "id": 2,
      "strategyId": 1,
      "merkleRoot": "e3cb8941e25dcdb36fc21acbe5f6c5a42e0d4f89839ae94952f0ebbd9acd04ac"
      "uri": "ipfs://Qma...."
    }
    ```
