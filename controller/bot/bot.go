package bot

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"hack-change-api/models/auxiliary"
	u "hack-change-api/muxutil"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var TalkToBalaboba = func(w http.ResponseWriter, r *http.Request) {
	seeds := []string {
		"Инвестор",
		"Инвестиции",
		"Покупка акций",
		"Ценные бумаги",
		"Фондовый рынок",
		"Инвестирование",
		"Брокерский счёт",
		"Покупка облигаций",
		"Дивиденды и купоны",
		"Покупка ценных бумаг",
		"Начните инвестировать",
		"Финансовые инструменты",
		"Инвестиционный портфель",
		"Торговля на фондовом рынке",
		"Инвестиционная деятельность",
	}

	seed := seeds[rand.Intn(len(seeds) - 1)]

	URL, _ := url.Parse("https://zeapi.yandex.net/lab/api/yalm/text3")
	body := ioutil.NopCloser(strings.NewReader(fmt.Sprintf(`{"query": "%s", "intro": %d, "filter": 0}`, seed, 0)))

	req := &http.Request{
		Method: "POST",
		URL:    URL,
		Body:   body,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"User-Agent":   {"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15"},
			"Origin":       {"https://yandex.ru"},
			"Referer":      {"https://yandex.ru/"},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		u.HandleInternalError(w, err)
		return
	}

	boba := &auxiliary.BalabobaOutput{}
	err = json.NewDecoder(res.Body).Decode(boba)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	log.Info(boba.Text)
	u.Respond(w, map[string]interface{}{"text": boba.Query + boba.Text})
}
