# Census3 Service documentation

## Index

1. [SQLite Database](#sqlite-database)
2. [Scanner](#scanner)
    1. [Holders providers](#holders-providers)
3. [API](#api)
    1. [Toolbox](#toolbox)

---

### SQLite Database
Contains all the service data. It contains six tables:
* *Tokens*: Contains the addresses, chainIDs and other attributes of the registered tokens in the service.
* *TokenTypes*: The relationship between the token type ID's and their names. It does not define the supported types, this is done by the scanner.
* *Strategies*: The list of strategies to create counts using the registered tokens.
* *Strategy Tokens*: The details of the tokens used in each strategy (such as the minimum balance required).
* *Censuses*: The list of censuses created.
* *Token Holders*: The list of addresses involved in the transfer of registered tokens.

Check out [`./db`](./db/) folder.

### Scanner
This component iterates until the service stops, keeping the holders and balances of each registered token updated. To do this, it performs the following steps:
1. Retrieves the token information from the database, trying to scan the new and smaller tokens first.
2. Depending on the type of token, it selects its owner provider and obtains the latest updates.
1. If the creation block of the token contract has not yet been calculated, the scanner will calculate it before updating its token holders.
3. Updates the token and its holder information in the database.
4. Loop forever.

Check out [`./scanner`](./scanner) folder.

#### Holders Providers
These components are defined by an interface that supports the creation of different types of holder providers, such as Web3 based (ERC20, ERC721, ERC777) or external service based (POAP or Gitcoin Passport).

This interface defines all the methods the scanner needs to work, and retains the logic of how the holder and balance updates are calculated.

See some examples in the [`./scanner/providers`](./scanner/providers) folder.

### API
The API has two goals:
* Provide information: It exposes the token information, the results of the scanner and the built strategies and censuses.
* Create resources: It allows to the user to register new tokens, create new strategies or build new censuses.

Check out [`./api`](./api) folder and the [API specification](./api/README.md).

#### Toolbox
It includes the following packages:
* **[`internal/lexer`](./internal/lexer/)**: The lexer package helps to analyse the predicate of strategies, allowing to define a syntax to combine the token holders of a group of tokens.
* **[`internal/queue`](./internal/queue/)**: The queue package allows to the API and other components of the Census3 service to run processes in the background, providing a way to know the status of a job or retrieve its results.
* **[`internal/roundedcensus`](./internal/roundedcensus/)**: The roundedcensus package provides an algorithm to anonymize participant balances in a token based census while maintaining a certain level of accuracy.
* **[`internal/strategyoperators`](./internal/strategyoperators/)**: This package defines the differents operators to be used in the strategies, and how they combine the holders of different tokens.


## Schema
![Components](https://i.postimg.cc/7LYXtDcF/census3-schema.png)