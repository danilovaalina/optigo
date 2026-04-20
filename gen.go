package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	names := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		names[i] = fmt.Sprintf("User%d", i+1)
	}

	data := map[string][]string{"names": names}
	file, _ := json.Marshal(data)
	os.WriteFile("data.json", file, 0644)
	fmt.Println("✅ data.json создан успешно")
}
