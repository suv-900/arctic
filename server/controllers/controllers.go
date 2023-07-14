package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"

	//	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var jwtkey = []byte(os.Getenv("JWT-KEY"))

const bcryptCost = 15

var expiryTime = time.Now().Add(10 * time.Minute)

type Userdb struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type Userjson struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type GetUsername struct {
	Username string `json:"username"`
}
type GetUserPass struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
type jwtClaims struct {
	Username string
	jwt.RegisteredClaims
}
type GetUsernamedb struct {
	Username string `db:"username"`
}

/*
var schema = `
CREATE TABLE User(

	username varchar(30) primary key not null,
	password varchar(30) not null,
	email varchar(30) not null

);
`
*/
func ConnectAndMigrateDB() {
	db1, err := sqlx.Open("mysql", "root:Core@123@/netflix?")

	if err != nil {
		log.Fatal(err)
	}
	db1.Ping()
	//sdb1.MustExec(schema)
	db = db1
}

//create user
//login user
//logout user

// get input fetch req to search for similar usernames
func SearchForSimilarUsernames(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var obj GetUsername
	err = json.Unmarshal(body, &obj)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	}
	var a GetUsernamedb
	err = db.Get(&a, "SELECT username FROM users WHERE username=?", obj.Username)
	if err != nil {
		w.WriteHeader(200)
		return
	} else {
		w.WriteHeader(409)

	}

}

func CreatNewUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//unmarshal
	var user Userjson
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
	}
	//hash pass
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	}
	user.Password = string(hashedpass)

	//transaction
	tx := db.MustBegin()

	_, err = tx.NamedExec("INSERT INTO users (username,password,email) VALUES(:username,:password,:email)", &Userdb{Username: user.Username, Password: user.Password, Email: user.Email})
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()

	if err != nil {
		w.WriteHeader(501)
		log.Fatal(err)

	}
	//create n return token
	claims := jwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "loginToken",
		Value:   tokenString,
		Expires: expiryTime,
	})
	w.WriteHeader(201)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	reqbody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var user Userjson
	err = json.Unmarshal(reqbody, &user)
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
	}
	var findUser GetUserPass
	err = db.Get(&findUser, "SELECT username,password FROM users WHERE username=?", user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		} else {
			fmt.Print(err)
			return
		}
	}
	emptystring := ""
	if findUser.Username == emptystring {
		//fmt.Printf("%v", findUser.Username)
		w.WriteHeader(404)
		return
	}
	//hash and check pass
	//nil on success
	result := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password))
	if result != nil {
		w.WriteHeader(401)
		fmt.Fprintf(w, "invalid password")
		return
	}

	claims := jwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(500)
		fmt.Print(err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "loginToken",
		Value:   tokenString,
		Expires: expiryTime,
	})
	w.WriteHeader(202)
}

func LogOutUser(w http.ResponseWriter, r *http.Request) {
	//destroy token
	http.SetCookie(w, &http.Cookie{
		Name:    "loginToken",
		Expires: time.Now(),
	})
}
