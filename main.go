package main

import (
	"fmt"

	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}
	mnemonic, _ := bip39.NewMnemonic(entropy)
	//mnemonic := ""

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		panic(err)
	}

	privKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		panic(err)
	}

	pubKeyHex, err := wallet.PublicKeyHex(account)
	if err != nil {
		panic(err)
	}


	address, err := wallet.Address(account)
	if err != nil {
		panic(err)
	}

	fmt.Println("mnemonic:", mnemonic)
	fmt.Println("address:", address)
	fmt.Println("privKeyHex:", privKeyHex)
	fmt.Println("pubKeyHex:", pubKeyHex)
}
