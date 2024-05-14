package main

import (
	"fmt"
	"sync"
)

// ----------------------------------------------------------------------------------------------------- Mutex for Sync
var mutex = &sync.Mutex{}

// -----------------------------------------------------------------------------------------------------   SystemConfig
type SystemConfig struct {
	bootMessage string
	step        int
}

// ---------------------------------------------------------------------------------------------------- getSystemConfig
var systemConfig *SystemConfig
var stepper int

// singleton - lazy init: only init when needed, no eager creation here
func getSystemConfig() *SystemConfig {
	if systemConfig == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if systemConfig == nil {
			fmt.Println("creating a new instance of SystemConfig")
			stepper = stepper + 1 // should never go beyond 1
			systemConfig = &SystemConfig{bootMessage: "Booted", step: stepper}
		} else {
			fmt.Println("systemConfig exists #1")
		}
	} else {
		fmt.Println("systemConfig exists #2")
	}

	return systemConfig
}
