package mongo

import "log"

func (db *database) Disconnect() {
	defer db.Cancel()
	defer func() {
		if err := db.Client.Disconnect(db.Ctx); err != nil {
			panic(err)
		}
		log.Printf("Disconected from MongoDB")
	}()
}
