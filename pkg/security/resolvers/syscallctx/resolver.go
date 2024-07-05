// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux

// Package syscallctx holds syscall context related files
package syscallctx

import (
	"encoding/binary"
	"fmt"

	lib "github.com/cilium/ebpf"

	manager "github.com/DataDog/ebpf-manager"

	"github.com/DataDog/datadog-agent/pkg/security/probe/managerhelper"
	"github.com/DataDog/datadog-agent/pkg/security/secl/model"
)

const (
	maxEntries = 8192 // see kernel definition
	argMaxSize = 128  // see kernel definition

	// types // see kernel definition
	strArg = 1
	intArg = 2
)

// KernelSyscallCtxEntryStruct maps the `syscall_ctx_entry_t` kernel structure
type KernelSyscallCtxEntryStruct struct {
	SyscallNr uint64
	ID        uint32
	Types     uint8
	Arg1      [argMaxSize]byte
	Arg2      [argMaxSize]byte
	Arg3      [argMaxSize]byte
}

// UnmarshalBinary unmarshalls a binary representation of itself
func (k *KernelSyscallCtxEntryStruct) UnmarshalBinary(data []byte) error {
	if len(data) < 8+4+1+argMaxSize*3 {
		return fmt.Errorf("invalid data size")
	}
	offset := 0
	k.SyscallNr = binary.LittleEndian.Uint64(data[offset:])
	offset += 8
	k.ID = binary.LittleEndian.Uint32(data[offset:])
	offset += 4
	k.Types = data[offset]
	offset++
	copy(k.Arg1[:], data[offset:offset+argMaxSize])
	offset += argMaxSize
	copy(k.Arg2[:], data[offset:offset+argMaxSize])
	offset += argMaxSize
	copy(k.Arg3[:], data[offset:offset+argMaxSize])
	return nil
}

// Resolver resolves syscall context
type Resolver struct {
	ctxMap *lib.Map
}

// Resolve resolves the syscall context
func (sr *Resolver) Resolve(ctxID uint32, ctx *model.SyscallContext) error {
	key := ctxID % maxEntries

	var ks KernelSyscallCtxEntryStruct
	if err := sr.ctxMap.Lookup(key, &ks); err != nil {
		return fmt.Errorf("unable to resolve the syscall context for `%d`: %w", ctxID, err)
	}

	if ctxID != ks.ID {
		return fmt.Errorf("incorrect id `%d` vs `%d`", ctxID, ks.ID)
	}

	isStrArg := func(pos int) bool {
		return (ks.Types>>(pos*2))&strArg > 0
	}

	isIntArg := func(pos int) bool {
		return (ks.Types>>(pos*2))&intArg > 0
	}

	if isStrArg(0) {
		arg, err := model.UnmarshalString(ks.Arg1[:], argMaxSize)
		if err != nil {
			return fmt.Errorf("unable to resolve the syscall context for `%d`: %w", ctxID, err)
		}
		ctx.StrArg1 = arg
	} else if isIntArg(0) {
		ctx.IntArg1 = int64(binary.NativeEndian.Uint64(ks.Arg1[:]))
	}

	if isStrArg(1) {
		arg, err := model.UnmarshalString(ks.Arg2[:], argMaxSize)
		if err != nil {
			return fmt.Errorf("unable to resolve the syscall context for `%d`: %w", ctxID, err)
		}
		ctx.StrArg2 = arg
	} else if isIntArg(1) {
		ctx.IntArg2 = int64(binary.NativeEndian.Uint64(ks.Arg2[:]))
	}

	if isStrArg(2) {
		arg, err := model.UnmarshalString(ks.Arg3[:], argMaxSize)
		if err != nil {
			return fmt.Errorf("unable to resolve the syscall context for `%d`: %w", ctxID, err)
		}
		ctx.StrArg3 = arg
	} else if isIntArg(2) {
		ctx.IntArg3 = int64(binary.NativeEndian.Uint64(ks.Arg3[:]))
	}

	ctx.SyscallNr = ks.SyscallNr

	return nil
}

// Start the syscall context resolver
func (sr *Resolver) Start(manager *manager.Manager) error {
	syscallCtx, err := managerhelper.Map(manager, "syscall_ctx")
	if err != nil {
		return err
	}
	sr.ctxMap = syscallCtx

	return nil
}

// Close the resolver
func (sr *Resolver) Close() error {
	return nil
}

// NewResolver returns a new syscall context
func NewResolver() *Resolver {
	return &Resolver{}
}
