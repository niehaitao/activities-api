package web

import (
	"activities-api/db"
	"activities-api/model"
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	summaryHandler := app.GetSummary
	app.handlers["/summary"] = summaryHandler
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	activities, err := a.d.GetActivities()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	summary := model.Summary{}

	for _, activity := range activities {
		switch activity.Action {
		case "Create":
			summary.Create++

		case "Update":
			summary.Update++

		case "Delete":
			summary.Delete++
		}
	}
	err = json.NewEncoder(w).Encode(summary)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) GetActivities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	activities, err := a.d.GetActivities()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
