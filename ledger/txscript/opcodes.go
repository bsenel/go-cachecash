package txscript

import (
	"bytes"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

type opcode struct {
	code   byte
	length int // The number of bytes taken by the instruction, including the opcode itself and any immediate.
	name   string
	fn     func(*VirtualMachine, *instruction) error
}

const (
	OP_0           uint8 = 0x00
	OP_DATA_20     uint8 = 0x14
	OP_DUP         uint8 = 0x76
	OP_HASH160     uint8 = 0xa9 // XXX: Do we want to change this to something SHA256-related?
	OP_EQUALVERIFY uint8 = 0x88
	OP_CHECKSIG    uint8 = 0xac
	// OP_CHECKSIGVERIFY            = 0xad
	// OP_CHECKMULTISIG             = 0xae
	// OP_CHECKMULTISIGVERIFY       = 0xaf
)

var OPCODES = map[uint8]*opcode{
	OP_0:           {OP_0, 1, "OP_0", opConst},
	OP_DATA_20:     {OP_DATA_20, 21, "OP_DATA_20", opPushData},
	OP_DUP:         {OP_DUP, 1, "OP_DUP", opDup},
	OP_HASH160:     {OP_HASH160, 1, "OP_HASH160", opHash160},
	OP_EQUALVERIFY: {OP_EQUALVERIFY, 1, "OP_EQUALVERIFY", opEqual},
	OP_CHECKSIG:    {OP_CHECKSIG, 1, "OP_CHECKSIG", opCheckSig},
	// OP_CHECKSIGVERIFY:      {OP_CHECKSIGVERIFY, 1, "OP_CHECKSIGVERIFY", opCheckSig},
	// OP_CHECKMULTISIG:       {OP_CHECKMULTISIG, 1, "OP_CHECKMULTISIG", opCheckSig},
	// OP_CHECKMULTISIGVERIFY: {OP_CHECKMULTISIGVERIFY, 1, "OP_CHECKMULTISIGVERIFY", opCheckSig},
}

var BadOpcode = errors.New("unexpected opcode for handler")

func opConst(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_0:
		vm.stack.PushInt(0)
		return nil
	default:
		return BadOpcode
	}
}

func opPushData(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_DATA_20:
		vm.stack.PushBytes(ins.immediates[0])
		return nil
	default:
		return BadOpcode
	}
}

func opDup(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_DUP:
		v, err := vm.stack.PeekBytes(vm.stack.Size() - 1)
		if err != nil {
			return err
		}
		vm.stack.PushBytes(v)
		return nil
	default:
		return BadOpcode
	}
}

func opHash160(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_HASH160:
		v, err := vm.stack.PopBytes()
		if err != nil {
			return err
		}

		vm.stack.PushBytes(Hash160Sum(v))
		return nil
	default:
		return BadOpcode
	}
}

func opEqual(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_EQUALVERIFY:
		// OP_EQUAL
		v0, err := vm.stack.PopBytes()
		if err != nil {
			return err
		}
		v1, err := vm.stack.PopBytes()
		if err != nil {
			return err
		}
		vm.stack.PushBool(bytes.Equal(v0, v1))

		// OP_VERIFY
		v, err := vm.stack.PopBool()
		if err != nil {
			return err
		}
		if !v {
			return errors.New("OP_EQUALVERIFY failed; top stack elements are not equal")
		}

		// Done!
		return nil
	default:
		return BadOpcode
	}
}

func opCheckSig(vm *VirtualMachine, ins *instruction) error {
	switch ins.opcode.code {
	case OP_CHECKSIG:
		vPubKey, err := vm.stack.PopNBytes(ed25519.PublicKeySize)
		if err != nil {
			vm.stack.PushBool(false)
			return err
		}
		pubKey := ed25519.PublicKey(vPubKey)

		sig, err := vm.stack.PopNBytes(ed25519.SignatureSize)
		if err != nil {
			vm.stack.PushBool(false)
			return err
		}

		hash, err := vm.tx.SigHash(vm.script, vm.txIdx, vm.inputAmount)
		if err != nil {
			vm.stack.PushBool(false)
			return err
		}

		valid := ed25519.Verify(pubKey, hash, sig)
		vm.stack.PushBool(valid)

		return nil
	default:
		return BadOpcode
	}
}
