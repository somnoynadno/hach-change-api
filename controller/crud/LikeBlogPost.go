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

var LikeBlogPostCreate = func(w http.ResponseWriter, r *http.Request) {
	LikeBlogPost := &entities.LikeBlogPost{}
	err := json.NewDecoder(r.Body).Decode(LikeBlogPost)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(LikeBlogPost).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(LikeBlogPost)
		u.RespondJSON(w, res)
	}
}

var LikeBlogPostRetrieve = func(w http.ResponseWriter, r *http.Request) {
	LikeBlogPost := &entities.LikeBlogPost{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BlogPost").Preload("User").First(&LikeBlogPost, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(LikeBlogPost)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var LikeBlogPostUpdate = func(w http.ResponseWriter, r *http.Request) {
	LikeBlogPost := &entities.LikeBlogPost{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&LikeBlogPost, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newLikeBlogPost := &entities.LikeBlogPost{}
	err = json.NewDecoder(r.Body).Decode(newLikeBlogPost)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&LikeBlogPost).Updates(newLikeBlogPost).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeBlogPostDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.LikeBlogPost{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeBlogPostQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.LikeBlogPost
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
	err := db.Preload("BlogPost").Preload("User").
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
		db.Model(&entities.LikeBlogPost{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
