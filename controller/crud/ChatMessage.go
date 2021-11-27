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

var ChatMessageCreate = func(w http.ResponseWriter, r *http.Request) {
	ChatMessage := &entities.ChatMessage{}
	err := json.NewDecoder(r.Body).Decode(ChatMessage)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(ChatMessage).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		res, _ := json.Marshal(ChatMessage)
		u.RespondJSON(w, res)
	}
}

var ChatMessageRetrieve = func(w http.ResponseWriter, r *http.Request) {
	ChatMessage := &entities.ChatMessage{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("From").Preload("To").First(&ChatMessage, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	res, err := json.Marshal(ChatMessage)
	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var ChatMessageUpdate = func(w http.ResponseWriter, r *http.Request) {
	ChatMessage := &entities.ChatMessage{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&ChatMessage, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleInternalError(w, err)
		}
		return
	}

	newChatMessage := &entities.ChatMessage{}
	err = json.NewDecoder(r.Body).Decode(newChatMessage)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db.Model(&ChatMessage).Update("seen", newChatMessage.Seen)
	err = db.Model(&ChatMessage).Updates(newChatMessage).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var ChatMessageDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&entities.ChatMessage{}, id).Error

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var ChatMessageQuery = func(w http.ResponseWriter, r *http.Request) {
	var agroModels []entities.ChatMessage
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

	from, err3 := strconv.Atoi(r.FormValue("from"))
	to, err4 := strconv.Atoi(r.FormValue("to"))

	db := db.GetDB()
	var err error

	if err3 == nil && err4 == nil {
		// query certain chat message history
		err = db.Preload("From").Preload("To").
			Where("from_id = ? and to_id = ?", from, to).
			Order(fmt.Sprintf("%s %s", sort, order)).
			Offset(start).Limit(end - start).Find(&agroModels).Error
	} else if err3 == nil {
		// query last messages for all chats from one user
		var toIDs []uint
		db.Table("chat_messages").Where("to_id = ", to).Select("to_id, max(id) as id").
			Group("to_id").Pluck("max(id) as id", &toIDs)
		err = db.Preload("From").Preload("To").
			Where("from_id = ?", from).Where("id IN (?)", toIDs).
			Order(fmt.Sprintf("%s %s", sort, order)).
			Offset(start).Limit(end - start).Find(&agroModels).Error
	} else if err4 == nil {
		// query last messages for all chats from one user
		var fromIDs []uint
		db.Table("chat_messages").Where("from_id = ", from).Select("from_id, max(id) as id").
			Group("from_id").Pluck("max(id) as id", &fromIDs)
		err = db.Preload("From").Preload("To").
			Where("to_id = ?", to).Where("id IN (?)", fromIDs).
			Order(fmt.Sprintf("%s %s", sort, order)).
			Offset(start).Limit(end - start).Find(&agroModels).Error
	} else {
		// query all messages in db (debug purpose)
		err = db.Preload("From").Preload("To").
			Order(fmt.Sprintf("%s %s", sort, order)).
			Offset(start).Limit(end - start).Find(&agroModels).Error
	}

	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	res, err := json.Marshal(agroModels)

	if err != nil {
		u.HandleInternalError(w, err)
	} else {
		db.Model(&entities.ChatMessage{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
