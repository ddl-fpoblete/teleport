// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64 && bpf

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type commandDataT struct {
	Pid     uint64
	Ppid    uint64
	Command [16]int8
	Type    uint32
	Argv    [128]int8
	Retval  int32
	Cgroup  uint64
}

// loadCommand returns the embedded CollectionSpec for command.
func loadCommand() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_CommandBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load command: %w", err)
	}

	return spec, err
}

// loadCommandObjects loads command and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*commandObjects
//	*commandPrograms
//	*commandMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadCommandObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadCommand()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// commandSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type commandSpecs struct {
	commandProgramSpecs
	commandMapSpecs
}

// commandSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type commandProgramSpecs struct {
	TracepointSyscallsSysEnterExecve   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_execve"`
	TracepointSyscallsSysEnterExecveat *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_execveat"`
	TracepointSyscallsSysExitExecve    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_execve"`
	TracepointSyscallsSysExitExecveat  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_execveat"`
}

// commandMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type commandMapSpecs struct {
	ExecveEvents     *ebpf.MapSpec `ebpf:"execve_events"`
	LostCounter      *ebpf.MapSpec `ebpf:"lost_counter"`
	LostDoorbell     *ebpf.MapSpec `ebpf:"lost_doorbell"`
	MonitoredCgroups *ebpf.MapSpec `ebpf:"monitored_cgroups"`
}

// commandObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadCommandObjects or ebpf.CollectionSpec.LoadAndAssign.
type commandObjects struct {
	commandPrograms
	commandMaps
}

func (o *commandObjects) Close() error {
	return _CommandClose(
		&o.commandPrograms,
		&o.commandMaps,
	)
}

// commandMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadCommandObjects or ebpf.CollectionSpec.LoadAndAssign.
type commandMaps struct {
	ExecveEvents     *ebpf.Map `ebpf:"execve_events"`
	LostCounter      *ebpf.Map `ebpf:"lost_counter"`
	LostDoorbell     *ebpf.Map `ebpf:"lost_doorbell"`
	MonitoredCgroups *ebpf.Map `ebpf:"monitored_cgroups"`
}

func (m *commandMaps) Close() error {
	return _CommandClose(
		m.ExecveEvents,
		m.LostCounter,
		m.LostDoorbell,
		m.MonitoredCgroups,
	)
}

// commandPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadCommandObjects or ebpf.CollectionSpec.LoadAndAssign.
type commandPrograms struct {
	TracepointSyscallsSysEnterExecve   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_execve"`
	TracepointSyscallsSysEnterExecveat *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_execveat"`
	TracepointSyscallsSysExitExecve    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_execve"`
	TracepointSyscallsSysExitExecveat  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_execveat"`
}

func (p *commandPrograms) Close() error {
	return _CommandClose(
		p.TracepointSyscallsSysEnterExecve,
		p.TracepointSyscallsSysEnterExecveat,
		p.TracepointSyscallsSysExitExecve,
		p.TracepointSyscallsSysExitExecveat,
	)
}

func _CommandClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed command_bpfel_arm64.o
var _CommandBytes []byte
