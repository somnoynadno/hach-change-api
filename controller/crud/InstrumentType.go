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

var InstrumentTypeCreate = func(w http.ResponseWriter, r *http.Request) {
	InstrumentType := &entities.InstrumentType{}
	err := json.NewDecoder(r.Body).Decode(InstrumentType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(InstrumentType).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(InstrumentType)
		u.RespondJSON(w, res)
	}
}

var InstrumentTypeRetrieve = func(w http.ResponseWriter, r *http.Request) {
	InstrumentType := &entities.InstrumentType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("Instruments").First(&InstrumentType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(InstrumentType)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var InstrumentTypeUpdate = func(w http.ResponseWriter, r *http.Request) {
	InstrumentType := &entities.InstrumentType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&InstrumentType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newInstrumentType := &entities.InstrumentType{}
	err = json.NewDecoder(r.Body).Decode(newInstrumentType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&InstrumentType).Updates(newInstrumentType).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var InstrumentTypeDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.InstrumentType{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var InstrumentTypeQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.InstrumentType
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
	err := db.Preload("Instruments").Order(fmt.Sprintf("%s %s", sort, order)).
		Offset(start).Limit(end - start).Find(&agroModels).Error

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	res, err := json.Marshal(agroModels)

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		db.Model(&entities.InstrumentType{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
