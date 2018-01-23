package main

import (
    "bytes"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
 //   "github.com/ethereum/go-ethereum/crypto"

    "github.com/howeyc/gopass"
    "log"
)

func main() {

    chainId := big.NewInt(1)

    fmt.Printf("Password: ")
    pass, _ := gopass.GetPasswd()
    strPass := fmt.Sprintf("%s", pass)
    senderPrivKey, err := keystore.DecryptKey([]byte(""), strPass)

    if err != nil {
        log.Panic(err)
    }

    recipientAddr := common.HexToAddress("")

    nonce := uint64(7)
    amount := big.NewInt(10000)
    gasLimit := uint64(100000)
    gasPrice := big.NewInt(20000000000) // 20 gwei

    tx := types.NewTransaction(nonce, recipientAddr, amount, gasLimit, gasPrice, nil)

    signer := types.NewEIP155Signer(chainId)
    signedTx, _ := types.SignTx(tx, signer, senderPrivKey.PrivateKey)
    fmt.Println(signedTx)

    var buff bytes.Buffer
    signedTx.EncodeRLP(&buff)
    fmt.Printf("0x%x\n", buff.Bytes())
}

