package main

import (
	"fmt"
	"io/ioutil"
	"log"

	wasm3 "github.com/tarcisiozf/go-wasm3"
)

const (
	wasmFilename = "examples/sum/sum.wasm"
	fnName       = "sum"
)

func main() {
	log.Print("Initializing WASM3")

	runtime := wasm3.NewRuntime(&wasm3.Config{
		Environment: wasm3.NewEnvironment(),
		StackSize:   64 * 1024,
		EnableWASI:  true,
	})
	log.Println("Runtime ok")

	wasmBytes, err := ioutil.ReadFile(wasmFilename)
	if err != nil {
		panic(err)
	}
	log.Printf("Read WASM module (%d bytes)\n", len(wasmBytes))

	module, err := runtime.ParseModule(wasmBytes)
	if err != nil {
		panic(err)
	}
	module, err = runtime.LoadModule(module)
	if err != nil {
		panic(err)
	}
	log.Print("Loaded module")

	fn, err := runtime.FindFunction(fnName)
	if err != nil {
		panic(err)
	}
	log.Printf("Found '%s' function (using runtime.FindFunction)", fnName)
	result, err := fn(1, 1)
	if err != nil {
		log.Printf("Error calling function: %s", err)
	} else {
		assert(result == 2, fmt.Sprintf("Expected sum(1, 1) to equal 2, got %v", result))
	}

	// Different call approach, retrieving functions from the module object:
	fn2, err := module.GetFunctionByName("sum")
	if err != nil {
		panic(err)
	}
	log.Printf("Found '%s' function (using module.GetFunctionByName)", fnName)
	result, _ = fn2.Call(2, 2)
	assert(result == 4, fmt.Sprintf("Expected sum(2, 2) to equal 4, got %v", result))
}

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
