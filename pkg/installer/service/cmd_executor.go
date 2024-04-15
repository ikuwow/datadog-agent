// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !windows

// Package service provides a way to interact with os services
package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/DataDog/datadog-agent/pkg/config/setup"
)

var updaterHelper = filepath.Join(setup.InstallPath, "bin", "installer", "helper")

const execTimeout = 30 * time.Second

// ChownDDAgent changes the owner of the given path to the dd-agent user.
func ChownDDAgent(path string) error {
	return executeHelperCommand(`{"command":"chown dd-agent","path":"` + path + `"}`)
}

// RemoveAll removes all files under a given path under /opt/datadog-packages regardless of their owner.
func RemoveAll(path string) error {
	return executeHelperCommand(`{"command":"rm","path":"` + path + `"}`)
}

func createAgentSymlink() error {
	return executeHelperCommand(`{"command":"agent-symlink"}`)
}

func rmAgentSymlink() error {
	return executeHelperCommand(`{"command":"rm-agent-symlink"}`)
}

// SetCapHelper sets cap setuid on the newly installed helper
func SetCapHelper(path string) error {
	return executeHelperCommand(`{"command":"setcap cap_setuid+ep", "path":"` + path + `"}`)
}

func executeHelperCommand(command string) error {
	cancelctx, cancelfunc := context.WithTimeout(context.Background(), execTimeout)
	defer cancelfunc()
	cmd := exec.CommandContext(cancelctx, updaterHelper, command)
	cmd.Stdout = os.Stdout
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	stderrOutput, err := io.ReadAll(stderr)
	if err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return errors.New(string(stderrOutput))
	}
	return nil
}

// BuildHelperForTests builds the helper binary for test
func BuildHelperForTests(pkgDir, binPath string, skipUIDCheck bool) error {
	updaterHelper = filepath.Join(binPath, "/helper")
	localPath, _ := filepath.Abs(".")
	targetDir := "datadog-agent/pkg"
	index := strings.Index(localPath, targetDir)
	pkgPath := localPath[:index+len(targetDir)]
	helperPath := filepath.Join(pkgPath, "installer", "service", "helper", "main.go")
	cmd := exec.Command("go", "build", fmt.Sprintf(`-ldflags=-X main.pkgDir=%s -X main.testSkipUID=%v`, pkgDir, skipUIDCheck), "-o", updaterHelper, helperPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
