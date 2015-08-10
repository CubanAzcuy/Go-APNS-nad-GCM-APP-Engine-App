package hello

import (
	"fmt"
	"net/http"
	"strconv"

	"appengine"
	"appengine/datastore"
	//	"appengine/user"
)

type User struct {
	PushId   string
	Platform int
	Device   string
	Id       string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/adduser", adduser)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello wordl")

}

func adduser(w http.ResponseWriter, r *http.Request) {
	i, _ := strconv.Atoi(r.FormValue("Platform"))

	u := User{
		PushId:   r.FormValue("PushId"),
		Platform: i,
		Device:   r.FormValue("Device"),
		Id:       r.FormValue("Id"),
	}
	fmt.Fprint(w, u)

	c := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(c, "User", nil)
	datastore.Put(c, key, &u)

}
