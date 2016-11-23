package api

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"

	"github.com/ivancduran/imgserver/backend/db"
	"github.com/ivancduran/imgserver/backend/libs"
	"github.com/ivancduran/imgserver/backend/models"
)

type AuthAPI struct {
	*iris.Context
}

// Register a new user in mongo, based on post method and return the user information
func (this AuthAPI) Register(ctx *iris.Context) {

	usr := models.User{}
	err := ctx.ReadJSON(&usr)

	if err != nil {
		ctx.JSON(iris.StatusOK, models.Err("4"))
		panic(err.Error())
	}

	pass := libs.Password{}
	passGen := pass.Gen(string(usr.Pass))
	usr.Timestamp = time.Now()
	usr.Pass = passGen
	// usr.Token = pass.Token()

	Db := db.New()
	defer Db.Close()

	// Insert
	if err := Db.C("users").Insert(&usr); err != nil {
		// Is a duplicate key, but we don't know which one
		ctx.JSON(iris.StatusOK, models.Err("5"))
		if Db.IsDup(err) {
			ctx.JSON(iris.StatusOK, models.Err("6"))
		}
	} else {

		j := new(libs.JWT)
		hmac := j.Secret()
		token := j.Create(usr.Email, usr.Username, usr.Role)
		tokenString, _ := token.SignedString(hmac)

		ctx.Session().Set("login", "true")
		ctx.Session().Set("token", tokenString)
		ctx.Session().Set("email", usr.Email)

		libs.SendEmailComplex([]string{usr.Email})

		ctx.JSON(iris.StatusOK, iris.Map{"response": true,
			"token":    tokenString,
			"username": usr.Username,
			"email":    usr.Email,
			"role":     0,
			"status":   true})
	}

}

// Login for user and return the user information, based on post method
func (this AuthAPI) Login(ctx *iris.Context) {

	reqData := models.User{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		panic(err.Error())
	}

	result := models.User{}

	_email := reqData.Email
	_pass := reqData.Pass

	// result := models.User{}
	// _email := string(ctx.FormValue("email"))
	// _pass := string(ctx.FormValue("pass"))

	Db := db.New()
	defer Db.Close()

	if err := Db.C("users").Find(bson.M{"email": _email}).One(&result); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("9"))
		return
	}

	pass := libs.Password{}
	var cp = pass.Compare(result.Pass, _pass)

	if cp {
		// token := pass.Token()

		j := new(libs.JWT)
		hmac := j.Secret()
		token := j.Create(result.Email, result.Username, result.Role)
		tokenString, _ := token.SignedString(hmac)

		// Update
		colQuerier := bson.M{"_id": result.Uid}
		change := bson.M{"$set": bson.M{"token": tokenString}}

		err = Db.C("users").Update(colQuerier, change)
		if err != nil {
			panic(err)
		}

		ctx.Session().Set("login", "true")
		ctx.Session().Set("token", tokenString)
		ctx.Session().Set("email", result.Email)
		ctx.Session().Set("name", result.Username)

		// libs.SendEmail([]string{"ivan.cduran@gmail.com"}, "prueba de contenido")

		ctx.JSON(iris.StatusOK,
			iris.Map{"response": true,
				"token":    tokenString,
				"status":   true,
				"username": result.Username,
				"role":     result.Role})

	} else {
		ctx.JSON(iris.StatusOK, models.Err("7"))
	}

}

// Require JWT for check the token in mongodb
func (this AuthAPI) Check(ctx *iris.Context) {

	user := ctx.Get("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	fmt.Println(username)

	_pass := string(ctx.FormValue("pass"))
	token := ctx.Session().GetString("token")

	pass := libs.Password{}
	cp := pass.Compare(token, _pass)

	if cp {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("8"))
	}

}

// Require JWT
func (this AuthAPI) Session(ctx *iris.Context) {

	login := ctx.Session().GetString("login")
	token := ctx.Session().GetString("token")

	if token != "" && login == "true" {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("8"))
	}

}

// Require JWT
func (this AuthAPI) AuthCheck(ctx *iris.Context) {

	result := models.User{}

	_username := string(ctx.FormValue("username"))
	_token := string(ctx.FormValue("token"))

	if _username != "" && _token != "" {

		Db := db.New()
		defer Db.Close()

		if err := Db.C("users").Find(bson.M{"username": _username, "token": _token}).One(&result); err != nil {
			ctx.JSON(iris.StatusOK, models.Err("9"))
			return
		}

		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": result.Token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("8"))
	}

}

// Require JWT
func (this AuthAPI) UpdateChannel(ctx *iris.Context) {

	Db := db.New()
	defer Db.Close()

	req := models.User{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		panic(err.Error())
	}

	// Update
	colQuerier := bson.M{"token": req.Token,
		"username": req.Username}
	change := bson.M{"$set": bson.M{"img": req.Img,
		"description": req.Description,
		"title":       req.Title}}

	err = Db.C("users").Update(colQuerier, change)
	if err != nil {
		panic(err)
		ctx.JSON(iris.StatusOK, models.Err("8"))
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true})
	}

}

func (this AuthAPI) Suspend(ctx *iris.Context) {

}

func (this AuthAPI) Logout(ctx *iris.Context) {

	// token := ctx.Param("token")

	ctx.Session().Delete("token")
	ctx.Session().Delete("login")

	ctx.JSON(iris.StatusOK, iris.Map{"response": true})

}
