package db

import (
	"activities-api/model"
	"database/sql"
)

type DB interface {
	GetActivities() ([]*model.Activity, error)
}

type PostgresDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

func (d PostgresDB) GetActivities() ([]*model.Activity, error) {
	rows, err := d.db.Query("select name, action from activities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var activities []*model.Activity
	for rows.Next() {
		act := new(model.Activity)
		err = rows.Scan(&act.Name, &act.Action)
		if err != nil {
			return nil, err
		}
		activities = append(activities, act)
	}
	return activities, nil
}
