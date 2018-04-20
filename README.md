#  Go-web3 Client

[![Build Status](https://travis-ci.org/regcostajr/go-web3.svg?branch=master)](https://travis-ci.org/regcostajr/go-web3)

This is a Ethereum compatible Go Client
which implements the 
[Eth JSON RPC Module](https://github.com/ethereum/wiki/wiki/JSON-RPC) and
[NET JSON RPC Module](https://github.com/paritytech/parity/wiki/JSONRPC-net-module#net_version).

## Usage

#### Deploying a contract

```go

bytecode := ... #contract bytecode
abi := ... #contract abi

var connection = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))
contract, err := connection.Eth.NewContract(abi)

transaction := new(dto.TransactionParameters)
coinbase, err := connection.Eth.GetCoinbase()
transaction.From = coinbase
transaction.Gas = big.NewInt(4000000)

hash, err := contract.Deploy(transaction, bytecode, nil)

fmt.Println(hash)
	
```

#### Using contract public functions

```go

result, err = contract.Call(transaction, "balanceOf", coinbase)
if result != nil && err == nil {
	balance, _ := result.ToComplexIntResponse()
	fmt.Println(balance.ToBigInt())
}
	
```

#### Using contract payable functions

```go

hash, err = contract.Send(transaction, "approve", coinbase, 10)
	
```

#### Using RPC commands

GetBalance

```go

balance, err := connection.Eth.GetBalance(coinbase, block.LATEST)

```

SendTransaction

```go

transaction := new(dto.TransactionParameters)
transaction.From = coinbase
transaction.To = coinbase
transaction.Value = big.NewInt(10)
transaction.Gas = big.NewInt(40000)
transaction.Data = types.ComplexString("p2p transaction")

txID, err := connection.Eth.SendTransaction(transaction)

```


## Installation

### go get

```bash
go get -u github.com/regcostajr/go-web3
```

### glide

```bash
glide get github.com/regcostajr/go-web3
```

### Requirements

* go ^1.8.3
* golang.org/x/net

## Testing

Node running in dev mode:

```bash
geth --dev --ws --wsorigins="*" --rpc --rpcapi eth,web3,personal,ssh,net --mine
```

Full test:

```bash
go test -v test/modulename/*.go
```

Individual test:
```bash
go test -v test/modulename/filename.go
```

## License

Package go-web3 is licensed under the [GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html) License.

## Office website
– CirclePlus, [@Circleplus](https://www.circlechain.io)