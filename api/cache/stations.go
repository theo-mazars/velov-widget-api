package cache

import (
	"encoding/json"
	"fmt"
	"os"

	"velov-widget-api/api/types"
)

var Stations []types.StationStruct

func init() {
	file, err := os.Open("./resources/stations.json")
	if err != nil {
		fmt.Println("❌ Error opening stations.json file")
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Stations)
	if err != nil {
		fmt.Println("❌ Error decoding stations.json file")
		panic(err)
	}

	fmt.Println("✅ Stations cache initialized")
}
