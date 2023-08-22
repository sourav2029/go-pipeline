package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	out := `{"error": true}`
	var err error
	if err != nil {
		fmt.Print("[VaultUtil]: Error occured: ", err.Error())
		panic(err)
	}
	output := string(out)
	var resp map[string]interface{}
	err = json.Unmarshal([]byte(output), &resp)
	if err != nil {
		fmt.Print("[VaultUtil]: Error occured: ", err.Error())
		panic(err)
	}
	if resp["error"].(bool) {
		fmt.Print("[VaultUtil]: Error occured: ", err.Error())
		panic(resp["message"].(string))
	}
}
