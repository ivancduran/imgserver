package models

import (
	"time"

	"github.com/ivancduran/imgserver/backend/db"
)

type Transform struct {
	Response  bool   `json:"response", bson:"response"`
	Code      string `json:"code", bson:"code"`
	Url       string `json:"url", bson:"url"`
	Format    string `json:"format", bson:"format"`
	Extension string `json:"extension", bson:"extension"`
	Width     int    `json:"width", bson:"width"`
	Height    int    `json:"height", bson:"height"`
	Transform string `json:"transform", bson:"transform"`
	Face      int    `json:"face", bson:"face"`
	Bucket    string `json:"bucket", bson:"bucket"`
	// User      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp,omitempty" form:"timestamp"`
}

func (t Transform) Save() {

	Db := db.New()
	defer Db.Close()

	t.Timestamp = time.Now()

	if err := Db.C("transform").Insert(&t); err != nil {
		// Is a duplicate key, but we don't know which one
	}

}
