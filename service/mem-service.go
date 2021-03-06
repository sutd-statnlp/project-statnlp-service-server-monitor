package service

import (
	"github.com/shirou/gopsutil/mem"
)

// MemService .
type MemService struct {
}

// GetVirtualMemory .
func (memService MemService) GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

// GetSwapMemory .
func (memService MemService) GetSwapMemory() (*mem.SwapMemoryStat, error) {
	return mem.SwapMemory()
}
