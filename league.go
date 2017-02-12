package main

import (
	"fmt"
	"net/http"
)

func isUserInGame() bool {
	url := "https://na.api.pvp.net/observer-mode/rest/consumer/getSpectatorGameInfo/NA1/"+ SummonerId +"?api_key=" + ApiKey

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	if res.StatusCode != 200 {
		return false
	}

	return true

}
