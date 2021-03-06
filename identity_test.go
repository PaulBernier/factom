package factom_test

import (
	"encoding/json"

	ed "github.com/FactomProject/ed25519"

	. "github.com/FactomProject/factom"

	"testing"
)

func TestGetIdentityChainID(t *testing.T) {
	name := []string{"John", "Jacob", "Jingleheimer-Schmidt"}
	observedChainID := GetIdentityChainID(name)
	expectedChainID := "e0cf1713b492e09e783d5d9f4fc6e2c71b5bdc9af4806a7937a5e935819717e9"
	if observedChainID != expectedChainID {
		t.Errorf("got: %s but expected: %s", observedChainID, expectedChainID)
	}
}

func TestNewIdentityChain(t *testing.T) {
	name := []string{"John", "Jacob", "Jingleheimer-Schmidt"}
	secretKeys := []string{
		"idsec2rChEHLz3SPQQx3syQtB11pHAmxyGjux5FntnS7xqTCieHxxTc",
		"idsec1xuUyeCCrJhsojf2wLAZqRxPzPFR8Gidd9DRRid1yGy8ncAJG3",
		"idsec2J3nNoqdiyboCBKDGauqN9Jb33dyFSqaJKZqTs6i5FmztsTn5f",
		"idsec1jztZ7dypqtwtPPWxybZFNpvvpUh6g8oog6Mnk2gGCm1pNBTgE",
	}
	var keys []string
	for _, v := range secretKeys {
		k, err := GetIdentityKey(v)
		if err != nil {
			t.Error(err)
		}
		keys = append(keys, k.PubString())
	}

	newChain, err := NewIdentityChain(name, keys)
	if err != nil {
		t.Errorf("Failed to compose identity chain struct %s", err.Error())
	}
	expectedChainID := "44abb806a2029ed77dca63770e2e4ac4b2fedd2e1847339ac59b180ee223eb84"
	t.Run("ChainID", func(t *testing.T) {
		if newChain.ChainID != expectedChainID {
			t.Errorf("expected:%s\nrecieved:%s", expectedChainID, newChain.ChainID)
		}
	})
	t.Run("Keys accessible from Content", func(t *testing.T) {
		var contentMap map[string]interface{}
		content := newChain.FirstEntry.Content
		if err := json.Unmarshal(content, &contentMap); err != nil {
			t.Errorf("Failed to unmarshal content")
		}
		for i, v := range contentMap["keys"].([]interface{}) {
			if keys[i] != v.(string) {
				t.Errorf("Keys not properly formatted")
			}
		}
	})
}

func TestNewIdentityKeyReplacementEntry(t *testing.T) {
	chainID := "44abb806a2029ed77dca63770e2e4ac4b2fedd2e1847339ac59b180ee223eb84"
	oldKey, err := GetIdentityKey("idsec1jztZ7dypqtwtPPWxybZFNpvvpUh6g8oog6Mnk2gGCm1pNBTgE")
	if err != nil {
		t.Error(err)
	}
	newKey, err := GetIdentityKey("idsec2J3nNoqdiyboCBKDGauqN9Jb33dyFSqaJKZqTs6i5FmztsTn5f")
	if err != nil {
		t.Error(err)
	}
	signerKey, err := GetIdentityKey("idsec2wH72BNR9QZhTMGDbxwLWGrghZQexZvLTros2wCekkc62N9h7s")
	if err != nil {
		t.Error(err)
	}

	observedEntry, err := NewIdentityKeyReplacementEntry(chainID, oldKey.PubString(), newKey.PubString(), signerKey)
	if err != nil {
		t.Errorf("Failed to compose key replacement entry struct %s", err.Error())
	}

	t.Run("ChainID", func(t *testing.T) {
		if observedEntry.ChainID != chainID {
			t.Fail()
		}
	})
	t.Run("ExtIDs", func(t *testing.T) {
		if len(observedEntry.ExtIDs) != 5 {
			t.Errorf("len(ExtIDs) != 5")
		}
		if string(observedEntry.ExtIDs[0]) != "ReplaceKey" {
			t.Errorf("ReplaceKey is not first ExtID")
		}
		if string(observedEntry.ExtIDs[1]) != oldKey.String() ||
			string(observedEntry.ExtIDs[2]) != newKey.String() ||
			string(observedEntry.ExtIDs[4]) != signerKey.String() {
			t.Errorf("Keys not formatted properly")
		}
	})
	t.Run("Signature", func(t *testing.T) {
		observedSignature := new([64]byte)
		copy(observedSignature[:], observedEntry.ExtIDs[3])
		message := []byte(chainID + oldKey.String() + newKey.String())
		if !ed.Verify(signerKey.Pub, message, observedSignature) {
			t.Fail()
		}
	})
}

func TestNewIdentityAttributeEntry(t *testing.T) {
	receiverChainID := "5ef81cd345fd497a376ca5e5670ef10826d96e73c9f797b33ea46552a47834a3"
	destinationChainID := "5a402200c5cf278e47905ce52d7d64529a0291829a7bd230072c5468be709069"
	signerChainID := "44abb806a2029ed77dca63770e2e4ac4b2fedd2e1847339ac59b180ee223eb84"
	signerKey, err := GetIdentityKey("idsec2J3nNoqdiyboCBKDGauqN9Jb33dyFSqaJKZqTs6i5FmztsTn5f")
	if err != nil {
		t.Errorf("Failed to get identity key")
	}
	attributesJSON := `[{"key":"email","value":"abc@def.ghi"}]`

	observedEntry := NewIdentityAttributeEntry(receiverChainID, destinationChainID, attributesJSON, signerKey, signerChainID)

	if !IsValidAttribute(observedEntry) {
		t.Errorf("Improperly formatted attribute")
	}

	var attributes []IdentityAttribute
	if err = json.Unmarshal(observedEntry.Content, &attributes); err != nil {
		t.Errorf("Failed to unmarshal content: %v", err)
	}
	if attributes[0].Key != "email" {
		t.Errorf("Incorrect key")
	}
	if attributes[0].Value != "abc@def.ghi" {
		t.Errorf("Incorrect value")
	}
}

func TestNewIdentityAttributeEndorsementEntry(t *testing.T) {
	destinationChainID := "5a402200c5cf278e47905ce52d7d64529a0291829a7bd230072c5468be709069"
	signerChainID := "44abb806a2029ed77dca63770e2e4ac4b2fedd2e1847339ac59b180ee223eb84"
	signerKey, _ := GetIdentityKey("idsec2J3nNoqdiyboCBKDGauqN9Jb33dyFSqaJKZqTs6i5FmztsTn5f")
	entryHash := "52385948ea3ab6fd67b07664ac6a30ae5f6afa94427a547c142517beaa9054d0"

	observedEntry := NewIdentityAttributeEndorsementEntry(destinationChainID, entryHash, signerKey, signerChainID)

	if !IsValidEndorsement(observedEntry) {
		t.Errorf("Improperly formatted attribute endorsement")
	}
}
