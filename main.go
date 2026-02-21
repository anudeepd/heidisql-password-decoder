package main

import (
	"strconv"
	"syscall/js"
)

func heidiDecode(hex string) any {
	if len(hex) < 2 {
		return js.ValueOf(map[string]any{
			"result": "",
			"error":  "Hex string too short",
		})
	}

	shift, err := strconv.Atoi(string(hex[len(hex)-1]))
	if err != nil {
		return js.ValueOf(map[string]any{
			"result": "",
			"error":  "Invalid shift value (last character must be 0-9)",
		})
	}

	hexBody := hex[:len(hex)-1]

	if len(hexBody)%2 != 0 {
		return js.ValueOf(map[string]any{
			"result": "",
			"error":  "Invalid hex length (must be an even number of characters)",
		})
	}

	var result []byte
	for i := 0; i < len(hexBody); i += 2 {
		val, err := strconv.ParseInt(hexBody[i:i+2], 16, 32)
		if err != nil {
			return js.ValueOf(map[string]any{
				"result": "",
				"error":  "Invalid hex string (contains non-hex characters)",
			})
		}
		result = append(result, byte(int(val)-shift))
	}

	return js.ValueOf(map[string]any{
		"result": string(result),
		"error":  "",
	})
}

func main() {
	js.Global().Set("heidiDecode", js.FuncOf(func(this js.Value, args []js.Value) any {
		defer func() {
			// Recover from panics to prevent WASM instance from breaking
			if r := recover(); r != nil {
			}
		}()
		return heidiDecode(args[0].String())
	}))

	<-make(chan bool)
}
