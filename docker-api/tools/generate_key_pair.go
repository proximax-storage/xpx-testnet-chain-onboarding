package main

import (
    "fmt"
    crypto "github.com/proximax-storage/go-xpx-crypto"
)

func main() {
    KeyPair, _ := crypto.NewRandomKeyPair()
    
	fmt.Printf("BOOT_KEY_PRIVATE_KEY:\t%x\n",KeyPair.PrivateKey.Raw)
    fmt.Printf("BOOT_KEY_PUBLIC_KEY:\t%x\n",KeyPair.PublicKey.Raw)
	
	RestKeyPair, _ := crypto.NewRandomKeyPair()
	fmt.Printf("REST_PRIVATE_KEY:\t%x\n",RestKeyPair.PrivateKey.Raw)
}