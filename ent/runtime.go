// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/artmisxyz/blockinspector/ent/event"
	"github.com/artmisxyz/blockinspector/ent/schema"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3collect"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3poolcreated"
	"github.com/artmisxyz/blockinspector/ent/uniswapv3transfer"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescName is the schema descriptor for name field.
	eventDescName := eventFields[0].Descriptor()
	// event.NameValidator is a validator for the "name" field. It is called by the builders before save.
	event.NameValidator = eventDescName.Validators[0].(func(string) error)
	// eventDescSignature is the schema descriptor for signature field.
	eventDescSignature := eventFields[1].Descriptor()
	// event.SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	event.SignatureValidator = eventDescSignature.Validators[0].(func(string) error)
	// eventDescAddress is the schema descriptor for address field.
	eventDescAddress := eventFields[2].Descriptor()
	// event.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	event.AddressValidator = eventDescAddress.Validators[0].(func(string) error)
	// eventDescTxHash is the schema descriptor for tx_hash field.
	eventDescTxHash := eventFields[4].Descriptor()
	// event.TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	event.TxHashValidator = eventDescTxHash.Validators[0].(func(string) error)
	// eventDescBlockHash is the schema descriptor for block_hash field.
	eventDescBlockHash := eventFields[6].Descriptor()
	// event.BlockHashValidator is a validator for the "block_hash" field. It is called by the builders before save.
	event.BlockHashValidator = eventDescBlockHash.Validators[0].(func(string) error)
	// eventDescHash is the schema descriptor for hash field.
	eventDescHash := eventFields[8].Descriptor()
	// event.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	event.HashValidator = eventDescHash.Validators[0].(func(string) error)
	uniswapv3collectFields := schema.UniswapV3Collect{}.Fields()
	_ = uniswapv3collectFields
	// uniswapv3collectDescRecipient is the schema descriptor for recipient field.
	uniswapv3collectDescRecipient := uniswapv3collectFields[1].Descriptor()
	// uniswapv3collect.RecipientValidator is a validator for the "recipient" field. It is called by the builders before save.
	uniswapv3collect.RecipientValidator = uniswapv3collectDescRecipient.Validators[0].(func(string) error)
	uniswapv3poolcreatedFields := schema.UniswapV3PoolCreated{}.Fields()
	_ = uniswapv3poolcreatedFields
	// uniswapv3poolcreatedDescToken0 is the schema descriptor for token0 field.
	uniswapv3poolcreatedDescToken0 := uniswapv3poolcreatedFields[0].Descriptor()
	// uniswapv3poolcreated.Token0Validator is a validator for the "token0" field. It is called by the builders before save.
	uniswapv3poolcreated.Token0Validator = uniswapv3poolcreatedDescToken0.Validators[0].(func(string) error)
	// uniswapv3poolcreatedDescToken1 is the schema descriptor for token1 field.
	uniswapv3poolcreatedDescToken1 := uniswapv3poolcreatedFields[1].Descriptor()
	// uniswapv3poolcreated.Token1Validator is a validator for the "token1" field. It is called by the builders before save.
	uniswapv3poolcreated.Token1Validator = uniswapv3poolcreatedDescToken1.Validators[0].(func(string) error)
	// uniswapv3poolcreatedDescPool is the schema descriptor for pool field.
	uniswapv3poolcreatedDescPool := uniswapv3poolcreatedFields[4].Descriptor()
	// uniswapv3poolcreated.PoolValidator is a validator for the "pool" field. It is called by the builders before save.
	uniswapv3poolcreated.PoolValidator = uniswapv3poolcreatedDescPool.Validators[0].(func(string) error)
	uniswapv3transferFields := schema.UniswapV3Transfer{}.Fields()
	_ = uniswapv3transferFields
	// uniswapv3transferDescFrom is the schema descriptor for from field.
	uniswapv3transferDescFrom := uniswapv3transferFields[1].Descriptor()
	// uniswapv3transfer.FromValidator is a validator for the "from" field. It is called by the builders before save.
	uniswapv3transfer.FromValidator = uniswapv3transferDescFrom.Validators[0].(func(string) error)
	// uniswapv3transferDescTo is the schema descriptor for to field.
	uniswapv3transferDescTo := uniswapv3transferFields[2].Descriptor()
	// uniswapv3transfer.ToValidator is a validator for the "to" field. It is called by the builders before save.
	uniswapv3transfer.ToValidator = uniswapv3transferDescTo.Validators[0].(func(string) error)
}