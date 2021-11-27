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

var LikeThreadCommentCreate = func(w http.ResponseWriter, r *http.Request) {
	LikeThreadComment := &entities.LikeThreadComment{}
	err := json.NewDecoder(r.Body).Decode(LikeThreadComment)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(LikeThreadComment).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(LikeThreadComment)
		u.RespondJSON(w, res)
	}
}

var LikeThreadCommentRetrieve = func(w http.ResponseWriter, r *http.Request) {
	LikeThreadComment := &entities.LikeThreadComment{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("ThreadComment").Preload("User").First(&LikeThreadComment, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(LikeThreadComment)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var LikeThreadCommentUpdate = func(w http.ResponseWriter, r *http.Request) {
	LikeThreadComment := &entities.LikeThreadComment{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&LikeThreadComment, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newLikeThreadComment := &entities.LikeThreadComment{}
	err = json.NewDecoder(r.Body).Decode(newLikeThreadComment)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&LikeThreadComment).Updates(newLikeThreadComment).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeThreadCommentDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.LikeThreadComment{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeThreadCommentQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.LikeThreadComment
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
	err := db.Preload("ThreadComment").Preload("User").
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
		db.Model(&entities.LikeThreadComment{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
