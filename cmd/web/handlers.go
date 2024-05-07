package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (app *Application) homepage(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	err := app.page.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}
	// w.Write([]byte("Hello from Nfactorial"))
}

func (app *Application) getMeetupsDB() (*sql.Rows, int, error) {
	cmd := `select id, name, date, address from meetups`
	rows, err := app.db.Query(cmd)
	if err != nil {
		return nil, 0, err
	}
	countcmd := `select count(id) from meetups`
	var count int
	err = app.db.QueryRow(countcmd).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func (app *Application) getMeetupInfoDB(id int) *sql.Row {
	cmd := fmt.Sprintf(`select * from meetups where id = %v`, id)
	row := app.db.QueryRow(cmd)
	return row
}

func (app *Application) getMeetupsApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, count, err := app.getMeetupsDB()
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(`["500 Internal Server Error"]`))
		return
	}
	responseList := make([]string, count)
	for i := 0; rows.Next(); i++ {
		var id, name, date, address string
		copyerr := rows.Scan(&id, &name, &date, &address)
		if copyerr != nil {
			log.Println(copyerr.Error())
			w.Write([]byte(`["500 Internal Server Error"]`))
			return
		}
		responseList[i] = fmt.Sprintf(`{"id":"%v","name":"%v","date":"%v","address":"%v"}`, id, name, date, address)
	}
	w.Write([]byte(fmt.Sprintf("[%v]", strings.Join(responseList, ","))))
}

func (app *Application) getMeetupInfoApi(w http.ResponseWriter, r *http.Request) {
	idr, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || idr < 0 {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var id, name, description, date, address string
	err = app.getMeetupInfoDB(idr).Scan(&id, &name, &description, &date, &address)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(`["500 Internal Server Error"]`))
		return
	}
	w.Write([]byte(fmt.Sprintf(
		`{"id":"%v","name":"%v","description":"%v","date":"%v","address":"%v"}`,
		id, name, description, date, address)))
	// w.Write([]byte(fmt.Sprintf(`{"id": %v}`, id)))
}
