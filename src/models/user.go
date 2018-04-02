package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Uid         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name        string        `json:"name" bson:"name" form:"name"`
		Username    string        `json:"username" bson:"username" form:"username"`
		Email       string        `json:"email" bson:"email" form:"email"`
		Img         string        `json:"img" bson:"img" form:"img"`
		Pass        string        `json:"pass" bson:"pass" form:"pass"`
		Role        int8          `json:"role" bson:"role" form:"role"`
		Active      int8          `json:"active" bson:"active" form:"active"`
		Ip          string        `json:"ip" bson:"ip" form:"ip"`
		Token       string        `json:"token" bson:"token" form:"token"`
		Gender      bool          `json:"gender" bson:"gender" form:"gender"`
		Birth       string        `json:"birth" bson:"birth" form:"birth"`
		Description string        `json:"description" bson:"description" form:"description"`
		Title       string        `json:"title" bson:"title" form:"title"`
		Timestamp   time.Time     `json:"timestamp" bson:"timestamp,omitempty" form:"timestamp"`
	}
)
