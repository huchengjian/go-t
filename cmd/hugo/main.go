package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"github.com/ethereum/go-ethereum/cmd/hugo"
)

func main(){
	test()
}

func test() {
	fmt.Print("haha succ\n")

	rawurl := "ws://192.168.1.213:8546"


	rawClient, _ := ethclient.Dial(rawurl)
	//client, _ := geth.NewEthereumClient(rawurl)
	ctx := context.Background()
	var a *big.Int
	a = big.NewInt(100)
	block, _:= rawClient.BlockByNumber(ctx, a)
	fmt.Print(*block.Difficulty())

	filter := ethereum.FilterQuery{}
	filter.Addresses = make([]common.Address, 0)
	filter.Addresses = append(filter.Addresses, common.HexToAddress("0x..."))
	filter.FromBlock = big.NewInt(1000000)

	addr := "7C564Ad23683172086A6D88266A42F250B547Ba2"
	priv := "96c8b4289e8c6f30d0f943ac0b16779027cacae043cd54eca61a5cf09035a713"

	key, _ := crypto.HexToECDSA(priv)
	auth := bind.NewKeyedTransactor(key)
	log.Println(auth.From.Hex())


	contractAddress := common.HexToAddress(addr)
	tradeClient, _ := trade.NewTrade(contractAddress, rawClient)

	order_id := big.NewInt(19998)
	num := big.NewInt(3)

	tx, err := tradeClient.CreateNewTradeOrder(auth, order_id, contractAddress, num)
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1431798),
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	sub, err := rawClient.SubscribeFilterLogs(ctx, query, ch)


	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			fmt.Println("Log:", log)
		}
	}
}

func call()  {

	rawurl := "ws://192.168.1.213:8546"
	rawClient, _ := ethclient.Dial(rawurl)
	ctx := context.Background()

	var a *big.Int
	a = big.NewInt(100)
	block, _:= rawClient.BlockByNumber(ctx, a)
	fmt.Print(*block.Difficulty())

	addr := "c4a4576cbb766b8d5c55121862adaf5a9aa0e5c1" //greet
	priv := "96c8b4289e8c6f30d0f943ac0b16779027cacae043cd54eca61a5cf09035a713"

	key, _ := crypto.HexToECDSA(priv)
	auth := bind.NewKeyedTransactor(key)
	log.Println(auth.From.Hex())


	contractAddress := common.HexToAddress(addr)
	greeter, _ := trade.NewGreeter(contractAddress, rawClient)

	tx, err := greeter.Greet(auth, "aaabbbcccdddeeefffggghhhiiijjjkkklll")
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1431798),
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	sub, err := rawClient.SubscribeFilterLogs(ctx, query, ch)


	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			fmt.Println("Log:", log.Address.Hex(), log.Data, log.Topics[0])
		}
	}
}

