package main

import (
	"log"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

func muteThemAll() {
	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	tw := anaconda.NewTwitterApi(OAuthToken, OAuthSecret)
	defer tw.Close()
	var v url.Values
	followers := tw.GetFriendsListAll(v)
	for {
		t := time.After(time.Second * 3)
		select {
		case uids, ok := <-followers:
			if !ok {
				log.Printf("Follwing chan closed.")
				return
			}
			if uids.Error != nil {
				log.Printf("Error: %s", uids.Error)
				continue
			}
			log.Printf("-- NEW PAGE --")
			for _, uid := range uids.Friends {
				for {
					u, err := tw.MuteUserId(uid.Id, v)
					if err == nil {
						log.Printf("Mute: \"%s\"(%s, %d)", u.Name, u.ScreenName, u.Id)
						break
					}
					log.Printf("Error while muting: %v", err)
					<-time.After(time.Minute)
				}
			}
		case <-t:
			log.Printf("Not received. Repeat...")
		}
	}
}

func main() {
	muteThemAll()
	log.Printf("Done.")
}
