// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wallet

import (
	"fmt"
	"sync"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/factoid"
)

// Wallet is a connection to a Factom Wallet Database
type Wallet struct {
	*WalletDatabaseOverlay
	DBPath       string
	txlock       sync.Mutex
	transactions map[string]*factoid.Transaction
	txdb         *TXDatabaseOverlay
}

func (w *Wallet) InitWallet() error {
	dbSeed, err := w.GetOrCreateDBSeed()
	if err != nil {
		return err
	}
	if dbSeed == nil {
		return fmt.Errorf("dbSeed not present in DB")
	}
	return nil
}

func NewOrOpenLevelDBWallet(path string) (*Wallet, error) {
	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	db, err := NewLevelDB(path)
	if err != nil {
		return nil, err
	}
	w.WalletDatabaseOverlay = db
	err = w.InitWallet()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewOrOpenBoltDBWallet(path string) (*Wallet, error) {
	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	db, err := NewBoltDB(path)
	if err != nil {
		return nil, err
	}
	w.WalletDatabaseOverlay = db
	err = w.InitWallet()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewMapDBWallet() (*Wallet, error) {
	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	db := NewMapDB()
	w.WalletDatabaseOverlay = db
	err := w.InitWallet()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewEncryptedBoltDBWallet(path, password string) (*Wallet, error) {
	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	db, err := NewEncryptedBoltDB(path, password)
	if err != nil {
		return nil, err
	}
	w.WalletDatabaseOverlay = db
	err = w.InitWallet()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewEncryptedBoltDBWalletAwaitingPassphrase(path string) (*Wallet, error) {
	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	w.DBPath = path
	return w, nil
}

// Close closes a Factom Wallet Database
func (w *Wallet) Close() error {
	return w.DBO.Close()
}

// AddTXDB allows the wallet api to read from a local transaction cashe.
func (w *Wallet) AddTXDB(t *TXDatabaseOverlay) {
	w.txdb = t
}

func (w *Wallet) TXDB() *TXDatabaseOverlay {
	return w.txdb
}

// GenerateECAddress creates and stores a new Entry Credit Address in the
// Wallet. The address can be reproduced in the future using the Wallet Seed.
func (w *Wallet) GenerateECAddress() (*factom.ECAddress, error) {
	return w.GetNextECAddress()
}

// GenerateFCTAddress creates and stores a new Factoid Address in the Wallet.
// The address can be reproduced in the future using the Wallet Seed.
func (w *Wallet) GenerateFCTAddress() (*factom.FactoidAddress, error) {
	return w.GetNextFCTAddress()
}

// GetAllAddresses retrieves all Entry Credit and Factoid Addresses from the
// Wallet Database.
func (w *Wallet) GetAllAddresses() ([]*factom.FactoidAddress, []*factom.ECAddress, error) {
	fcs, err := w.GetAllFCTAddresses()
	if err != nil {
		return nil, nil, err
	}
	ecs, err := w.GetAllECAddresses()
	if err != nil {
		return nil, nil, err
	}

	return fcs, ecs, nil
}

// GetSeed returns the string representaion of the Wallet Seed. The Wallet Seed
// can be used to regenerate the Factoid and Entry Credit Addresses previously
// generated by the wallet. Note that Addresses that are imported into the
// Wallet cannot be regenerated using the Wallet Seed.
func (w *Wallet) GetSeed() (string, error) {
	seed, err := w.GetDBSeed()
	if err != nil {
		return "", err
	}

	return seed.MnemonicSeed, nil
}

func (w *Wallet) GetVersion() string {
	return WalletVersion
}

func (w *Wallet) GetApiVersion() string {
	return ApiVersion
}
