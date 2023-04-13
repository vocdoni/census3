# API endpoints

+ **GET** `/tokens`
List of already added tokens.
    - 📥 response:
    ```json
    {
        tokens: [
            {
                "id": "wANT"
                "address": "0x1234",
                "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
                "creationBlock": 123456
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
Adds a new token and triggers a new scan for it, starting from the creation block.
    - 📤 request:
    ```json
    {
        "id": "wANT"
        "address": "0x1234",
        "type": "erc20|erc721|erc777|erc1155|nation3|wANT",
        "creationBlock": 123456
    }
    ```

+ **GET** `/tokens/:id`
Returns the information about the token referenced by the provided id.
    - 📥 response:
    ```json
    {
        "address": "0x1324",
        "type": "erc20",
        "decimals": 18,
        "creationBlock": 123456,
        "name": "Amazing token",
        "totalSupply": 1233456,
        "symbol": "AT",
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
              "id": "wANT",
              "minBalance": "10000",
              "method": "balanceOfAt" 
          },
          {
              "id": "USDC",
              "minBalance": "20000",
              "method": "balanceOf" 
          },
          {
              "id": "ANT",
              "minBalance": "1",
              "method": "balanceOf" 
          }
       ],
       "predicate": "(wANT OR ANT) AND USDC"
     }
    ```
    - 📥 response:
    ```json
    {
        "strategyId": "0x12345"
    }
    ```

+ **GET** `/strategies`
Returns the ID's list of the strategies registered.
    - 📥 response:
    ```json
    {
        "strategies": [ "0x12345", "0x67890" ]
    }
    ```

+ **GET** `/strategies/{strategyId}`
Returns the information of the strategy related to the provided ID.
    - 📥 response:
    ```json
    {
        "tokens": [
          {
              "id": "wANT",
              "minBalance": "10000",
              "method": "balanceOfAt" 
          },
          {
              "id": "USDC",
              "minBalance": "20000",
              "method": "balanceOf" 
          },
          {
              "id": "ANT",
              "minBalance": "1",
              "method": "balanceOf" 
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
         "strategies": [ "0x12345", "0x67890" ]
     }
    ```

+ **POST** `/census`
Request the creation of a new census with the strategy provided for the `blockNumber` provided and returns the census ID.
     - 📤 request:
    ```json
    {
        "strategyId": "0x123",
        "blockNumber": "123445"
    }
    ```
    - 📥 response:
    ```json
      {
        "censusId": "0x123"
      }
    ```

+ **GET** `/census/{strategyId}/`
Returns a list of censusID for the strategy provided.
    - 📥 response:
    ```json
        {
            "censuses": [ "0x12345", "0x67890" ]
        }
    ```

+ **GET** `/census/{censusId}`
Returns the information of the snapshots related to the provided ID.
    - 📥 response: Returns 200 or 204
    ```json
    { 
      "root": "0x12345",
      "blockNumber": "12345",
      "uri": "ipfs://Qma...."
    }
    ```
