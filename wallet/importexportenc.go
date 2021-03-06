// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package wallet

import (
	"fmt"
	"os"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/factoid"
)

// ImportEncryptedWalletFromMnemonic creates a new wallet with a provided Mnemonic seed
// defined in bip-0039.
func ImportEncryptedWalletFromMnemonic(mnemonic, path, password string) (*Wallet, error) {
	mnemonic, err := factom.ParseMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	// check if the file exists
	_, err = os.Stat(path)
	if err == nil {
		return nil, fmt.Errorf("%s: file already exists", path)
	}

	db, err := NewEncryptedBoltDB(path, password)
	if err != nil {
		return nil, err
	}

	seed := new(DBSeed)
	seed.MnemonicSeed = mnemonic
	if err := db.InsertDBSeed(seed); err != nil {
		return nil, err
	}

	w := new(Wallet)
	w.transactions = make(map[string]*factoid.Transaction)
	w.WalletDatabaseOverlay = db

	return w, nil
}

// ExportEncryptedWallet writes all the secret/publilc key pairs from a wallet and the
// wallet seed in a pritable format.
func ExportEncryptedWallet(path, password string) (string, []*factom.FactoidAddress, []*factom.ECAddress, error) {
	// check if the file exists
	_, err := os.Stat(path)
	if err != nil {
		return "", nil, nil, err
	}

	w, err := NewEncryptedBoltDBWallet(path, password)
	if err != nil {
		return "", nil, nil, err
	}

	m, err := w.GetSeed()
	if err != nil {
		return "", nil, nil, err
	}
	fs, es, err := w.GetAllAddresses()
	if err != nil {
		return "", nil, nil, err
	}
	return m, fs, es, nil
}
