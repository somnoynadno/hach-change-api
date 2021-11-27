package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"hack-change-api/db"
	"hack-change-api/models/entities"
	u "hack-change-api/muxutil"
	"net/http"
	"strconv"
)

var UserCreate = func(w http.ResponseWriter, r *http.Request) {
	User := &entities.User{}
	err := json.NewDecoder(r.Body).Decode(User)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(User).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(User)
		u.RespondJSON(w, res)
	}
}

var UserRetrieve = func(w http.ResponseWriter, r *http.Request) {
	User := &entities.User{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BlogPosts").Preload("Comments").
		//Preload("Publishers").Preload("Subscribers").
		First(&User, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(User)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var UserUpdate = func(w http.ResponseWriter, r *http.Request) {
	User := &entities.User{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&User, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newUser := &entities.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&User).Updates(newUser).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var UserDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.User{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var UserQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.User
	var count string

	order := r.FormValue("_order")
	sort := r.FormValue("_sort")
	end, err1 := strconv.Atoi(r.FormValue("_end"))
	start, err2 := strconv.Atoi(r.FormValue("_start"))

	if err1 != nil || err2 != nil {
		u.HandleBadRequest(w, errors.New("bad _start or _end parameter value"))
		return
	}
	u.CheckOrderAndSortParams(&order, &sort)

	db := db.GetDB()
	err := db.Preload("BlogPosts").Preload("Comments").
		//Preload("Publishers").Preload("Subscribers").
		Order(fmt.Sprintf("%s %s", sort, order)).
		Offset(start).Limit(end - start).Find(&agroModels).Error

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	res, err := json.Marshal(agroModels)

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		db.Model(&entities.User{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
