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

var LikeCommentCreate = func(w http.ResponseWriter, r *http.Request) {
	LikeComment := &entities.LikeComment{}
	err := json.NewDecoder(r.Body).Decode(LikeComment)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(LikeComment).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(LikeComment)
		u.RespondJSON(w, res)
	}
}

var LikeCommentRetrieve = func(w http.ResponseWriter, r *http.Request) {
	LikeComment := &entities.LikeComment{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("Comment").Preload("User").First(&LikeComment, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(LikeComment)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var LikeCommentUpdate = func(w http.ResponseWriter, r *http.Request) {
	LikeComment := &entities.LikeComment{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&LikeComment, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newLikeComment := &entities.LikeComment{}
	err = json.NewDecoder(r.Body).Decode(newLikeComment)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&LikeComment).Updates(newLikeComment).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeCommentDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.LikeComment{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var LikeCommentQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.LikeComment
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
	err := db.Preload("Comment").Preload("User").
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
		db.Model(&entities.LikeComment{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
