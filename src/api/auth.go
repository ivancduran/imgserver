package api

// import (
// 	"fmt"
// 	"time"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/kataras1/iris"
// 	"github.com/labstack/echo"
// 	"gopkg.in/mgo.v2/bson"

// 	"github.com/ivancduran/imgserver/src/db"
// 	"github.com/ivancduran/imgserver/src/libs"
// 	"github.com/ivancduran/imgserver/src/models"
// )

// type AuthAPI struct {
// 	c *echo.Context
// }

// // Register a new user in mongo, based on post method and return the user information
// func (this AuthAPI) Register(c *echo.Context) {

// 	usr := new(models.User)
// 	err := this.Bind(&usr)

// 	if err != nil {
// 		c.JSON(echo.StatusOK, models.Err("4"))
// 		panic(err.Error())
// 	}

// 	pass := libs.Password{}
// 	passGen := pass.Gen(string(usr.Pass))
// 	usr.Timestamp = time.Now()
// 	usr.Pass = passGen
// 	// usr.Token = pass.Token()

// 	Db := db.New()
// 	defer Db.Close()

// 	// Insert
// 	if err := Db.C("users").Insert(&usr); err != nil {
// 		// Is a duplicate key, but we don't know which one
// 		c.JSON(iris.StatusOK, models.Err("5"))
// 		if Db.IsDup(err) {
// 			c.JSON(iris.StatusOK, models.Err("6"))
// 		}
// 	} else {

// 		j := new(libs.JWT)
// 		hmac := j.Secret()
// 		token := j.Create(usr.Email, usr.Username, usr.Role)
// 		tokenString, _ := token.SignedString(hmac)

// 		c.Session().Set("login", "true")
// 		c.Session().Set("token", tokenString)
// 		c.Session().Set("email", usr.Email)

// 		libs.SendEmailComplex([]string{usr.Email})

// 		c.JSON(iris.StatusOK, iris.Map{"response": true,
// 			"token":    tokenString,
// 			"username": usr.Username,
// 			"email":    usr.Email,
// 			"role":     0,
// 			"status":   true})
// 	}

// }

// // Login for user and return the user information, based on post method
// func (this AuthAPI) Login(ctx *iris.Context) {

// 	reqData := models.User{}
// 	err := c.ReadJSON(&reqData)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	result := models.User{}

// 	_email := reqData.Email
// 	_pass := reqData.Pass

// 	// result := models.User{}
// 	// _email := string(c.FormValue("email"))
// 	// _pass := string(c.FormValue("pass"))

// 	Db := db.New()
// 	defer Db.Close()

// 	if err := Db.C("users").Find(bson.M{"email": _email}).One(&result); err != nil {
// 		c.JSON(iris.StatusOK, models.Err("9"))
// 		return
// 	}

// 	pass := libs.Password{}
// 	var cp = pass.Compare(result.Pass, _pass)

// 	if cp {
// 		// token := pass.Token()

// 		j := new(libs.JWT)
// 		hmac := j.Secret()
// 		token := j.Create(result.Email, result.Username, result.Role)
// 		tokenString, _ := token.SignedString(hmac)

// 		// Update
// 		colQuerier := bson.M{"_id": result.Uid}
// 		change := bson.M{"$set": bson.M{"token": tokenString}}

// 		err = Db.C("users").Update(colQuerier, change)
// 		if err != nil {
// 			panic(err)
// 		}

// 		c.Session().Set("login", "true")
// 		c.Session().Set("token", tokenString)
// 		c.Session().Set("email", result.Email)
// 		c.Session().Set("name", result.Username)

// 		// libs.SendEmail([]string{"ivan.cduran@gmail.com"}, "prueba de contenido")

// 		c.JSON(iris.StatusOK,
// 			iris.Map{"response": true,
// 				"token":    tokenString,
// 				"status":   true,
// 				"username": result.Username,
// 				"role":     result.Role})

// 	} else {
// 		c.JSON(iris.StatusOK, models.Err("7"))
// 	}

// }

// // Require JWT for check the token in mongodb
// func (this AuthAPI) Check(ctx *iris.Context) {

// 	user := c.Get("jwt").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	username := claims["username"].(string)

// 	fmt.Println(username)

// 	_pass := string(c.FormValue("pass"))
// 	token := c.Session().GetString("token")

// 	pass := libs.Password{}
// 	cp := pass.Compare(token, _pass)

// 	if cp {
// 		c.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
// 	} else {
// 		c.JSON(iris.StatusOK, models.Err("8"))
// 	}

// }

// // Require JWT
// func (this AuthAPI) Session(ctx *iris.Context) {

// 	login := c.Session().GetString("login")
// 	token := c.Session().GetString("token")

// 	if token != "" && login == "true" {
// 		c.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
// 	} else {
// 		c.JSON(iris.StatusOK, models.Err("8"))
// 	}

// }

// // Require JWT
// func (this AuthAPI) AuthCheck(ctx *iris.Context) {

// 	result := models.User{}

// 	_username := string(c.FormValue("username"))
// 	_token := string(c.FormValue("token"))

// 	if _username != "" && _token != "" {

// 		Db := db.New()
// 		defer Db.Close()

// 		if err := Db.C("users").Find(bson.M{"username": _username, "token": _token}).One(&result); err != nil {
// 			c.JSON(iris.StatusOK, models.Err("9"))
// 			return
// 		}

// 		c.JSON(iris.StatusOK, iris.Map{"response": true, "token": result.Token})
// 	} else {
// 		c.JSON(iris.StatusOK, models.Err("8"))
// 	}

// }

// // Require JWT
// func (this AuthAPI) UpdateChannel(ctx *iris.Context) {

// 	Db := db.New()
// 	defer Db.Close()

// 	req := models.User{}
// 	err := c.ReadJSON(&req)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// Update
// 	colQuerier := bson.M{"token": req.Token,
// 		"username": req.Username}
// 	change := bson.M{"$set": bson.M{"img": req.Img,
// 		"description": req.Description,
// 		"title":       req.Title}}

// 	err = Db.C("users").Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 		c.JSON(iris.StatusOK, models.Err("8"))
// 	} else {
// 		c.JSON(iris.StatusOK, iris.Map{"response": true})
// 	}

// }

// func (this AuthAPI) Suspend(ctx *iris.Context) {

// }

// func (this AuthAPI) Logout(ctx *iris.Context) {

// 	// token := c.Param("token")

// 	c.Session().Delete("token")
// 	c.Session().Delete("login")

// 	c.JSON(iris.StatusOK, iris.Map{"response": true})

// }
