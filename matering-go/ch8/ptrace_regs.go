package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	var r syscall.PtraceRegs
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err := cmd.Start()
	if err != nil {
		fmt.Println("Start:", err)
		return
	}

	err = cmd.Wait()
	fmt.Printf("State:%v\n", err)
	wpid := cmd.Process.Pid

	err = syscall.PtraceGetRegs(wpid, &r)
	if err != nil {
		fmt.Println("PtraceGetRegs:", err)
		return
	}
	fmt.Printf("Registers: %#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

	time.Sleep(2 * time.Second)
	//$ go run ptrace_regs.go echo "Mastering Go!"
	//State:stop signal: trace/breakpoint trap
	//Registers: syscall.PtraceRegs{R15:0x0, R14:0x0, R13:0x0, R12:0x0, Rbp:0x0, Rbx:0x0, R11:0x0, R10:0x0, R9:0x0, R8:0x0, Rax:0x0, Rcx:0x0, Rdx:0x0, Rsi:0x0, Rdi:0x0, Orig_rax:0x3b, Rip:0x7f2f97c09090, Cs:0x33, Eflags:0x200, Rsp:0x7ffe524be070, Ss:0x2b, Fs_base:0x0, Gs_base:0x0, Ds:0x0, Es:0x0, Fs:0x0, Gs:0x0}
	//R15=0, Gs=0
	//Mastering Go!
}
