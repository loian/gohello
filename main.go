package main

import (
"fmt"
	"math/rand"
	"net/http"
	"time"
)

//spreadTheWords selects and returns a message to deliver to the world
func spreadTheWords(quotes []string) string {
	idx := rand.Int() % len(quotes)
	return quotes[idx]
}

func main() {
	rand.Seed(time.Now().Unix())

	pillsOfVadersWisdom := []string{
		"Luke, I am your father.",
		"I am altering the deal. Pray I don’t alter it any further.",
		"I find your lack of faith disturbing.",
		"The circle is now complete. When I left you, I was but the learner. Now I am the master.",
		"Now, release your anger. Only your hatred can destroy me.",
		"Be Careful Not To Choke On Your Aspirations.",
		"Give yourself to the Dark Side.",
		"I have brought peace, freedom, justice, and security to my new empire.",
		"He will join us or die!",
		"You don’t know the power of the dark side.",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, spreadTheWords(pillsOfVadersWisdom))
	})

	http.ListenAndServe(":80", nil)
}

