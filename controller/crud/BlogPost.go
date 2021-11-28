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

var BlogPostCreate = func(w http.ResponseWriter, r *http.Request) {
	BlogPost := &entities.BlogPost{}
	err := json.NewDecoder(r.Body).Decode(BlogPost)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	instruments := BlogPost.Instruments
	BlogPost.Instruments = nil

	err = db.Create(BlogPost).Error
	for _, i := range instruments {
		db.Exec("insert into post_instruments (blog_post_id, financial_instrument_id) values (?, ?)", BlogPost.ID, i.ID)
	}

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(BlogPost)
		u.RespondJSON(w, res)
	}
}

var BlogPostRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BlogPost := &entities.BlogPost{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("Author").Preload("Instruments").Preload("Comments").
		Preload("Comments.ThreadComments").Preload("Likes").First(&BlogPost, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(BlogPost)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var BlogPostUpdate = func(w http.ResponseWriter, r *http.Request) {
	BlogPost := &entities.BlogPost{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BlogPost, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newBlogPost := &entities.BlogPost{}
	err = json.NewDecoder(r.Body).Decode(newBlogPost)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BlogPost).Updates(newBlogPost).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BlogPostDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.BlogPost{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BlogPostQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.BlogPost
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
	err := db.Preload("Author").Preload("Instruments").Preload("Comments").
		Preload("Comments.ThreadComments").Preload("Likes").Order(fmt.Sprintf("%s %s", sort, order)).
		Offset(start).Limit(end - start).Find(&agroModels).Error

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	res, err := json.Marshal(agroModels)

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		db.Model(&entities.BlogPost{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
