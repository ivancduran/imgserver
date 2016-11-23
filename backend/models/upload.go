package models

import (
	"time"

	"github.com/ivancduran/imgserver/backend/db"
)

type Upload struct {
	Response  bool      `json:"response", bson:"response"`
	Code      string    `json:"code", bson:"code"`
	Url       string    `json:"url", bson:"url"`
	Format    string    `json:"format", bson:"format"`
	Extension string    `json:"extension", bson:"extension"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp,omitempty" form:"timestamp"`
}

func (t Upload) Save() {

	Db := db.New()
	defer Db.Close()

	t.Timestamp = time.Now()

	if err := Db.C("upload").Insert(&t); err != nil {
		// Is a duplicate key, but we don't know which one
	}

}
