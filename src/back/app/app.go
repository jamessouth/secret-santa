package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Card decodes a card
type Card struct {
	Name      string  `json:"name"`
	Level     int     `json:"level,omitempty"`
	CardLevel int     `json:"cardLevel"`
	MaxLevel  int     `json:"maxLevel,omitempty"`
	Count     int     `json:"count"`
	IconUrls  IconURL `json:"iconUrls,omitempty"`
	URL       string  `json:"url"`
	Display   string  `json:"display"`
	Rarity    Rarity  `json:"rarity"`
}

// Init starts the server
func Init() {
	// fs := http.FileServer(http.Dir("./dist"))
	// http.Handle("/", fs)
	http.HandleFunc("/", handleConns)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}
	log.Println("server running on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConns(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// }

	// playerID := r.FormValue("id")
	// if strings.Count(playerID, "#") > 1 {
	// 	playerID = strings.ReplaceAll(playerID, "#", "")
	// }
	// if !strings.HasPrefix(playerID, "#") {
	// 	playerID = "#" + playerID
	// }
	// playerID = url.QueryEscape(playerID)

	// token, err := os.Open(".\\tok.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer token.Close()
	// tok := new(bytes.Buffer)
	// tok.ReadFrom(token)

	// client := &http.Client{}
	// req, err := http.NewRequest(http.MethodGet, "https://api.clashroyale.com/v1/players/"+playerID, nil)
	// req.Header.Add("Authorization", "Bearer "+tok.String())

	// resp := data{Data: }L0R2LJ8V

	// res, err := json.Marshal(resp)
	// resp, err := client.Do(req)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// }
	body, err := ioutil.ReadFile("resp.json")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var player PlayerData
	err = json.Unmarshal(body, &player)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	newCards := []Card{}
	for _, c := range player.Cards {
		level := 13 - c.MaxLevel + c.Level
		newCards = append(newCards, Card{
			Name:      c.Name,
			Count:     c.Count,
			CardLevel: level,
			URL:       c.IconUrls.Medium,
			Display:   "flex",
			Rarity:    rarityData[c.MaxLevel],
		})
	}
	finalPlayer := PlayerData{
		Name:  player.Name,
		Cards: newCards,
	}

	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4444")
	pl, err := json.Marshal(finalPlayer)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(pl)
}
