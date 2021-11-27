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

var FinancialInstrumentCreate = func(w http.ResponseWriter, r *http.Request) {
	FinancialInstrument := &entities.FinancialInstrument{}
	err := json.NewDecoder(r.Body).Decode(FinancialInstrument)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(FinancialInstrument).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(FinancialInstrument)
		u.RespondJSON(w, res)
	}
}

var FinancialInstrumentRetrieve = func(w http.ResponseWriter, r *http.Request) {
	FinancialInstrument := &entities.FinancialInstrument{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("InstrumentType").Preload("BlogPosts").First(&FinancialInstrument, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(FinancialInstrument)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var FinancialInstrumentUpdate = func(w http.ResponseWriter, r *http.Request) {
	FinancialInstrument := &entities.FinancialInstrument{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&FinancialInstrument, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newFinancialInstrument := &entities.FinancialInstrument{}
	err = json.NewDecoder(r.Body).Decode(newFinancialInstrument)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&FinancialInstrument).Updates(newFinancialInstrument).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var FinancialInstrumentDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.FinancialInstrument{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var FinancialInstrumentQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.FinancialInstrument
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
	err := db.Preload("InstrumentType").Preload("BlogPosts").
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
		db.Model(&entities.FinancialInstrument{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
