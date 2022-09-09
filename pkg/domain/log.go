package audit

import (
	"errors"
	"time"
)

const (
	ENTITY_USER = "USER"
	ENTITY_POST = "POST"

	ACTION_CREATE   = "CREATE"
	ACTION_UPDATE   = "UPDATE"
	ACTION_GET      = "GET"
	ACTION_DELETE   = "DELETE"
	ACTION_LIST     = "LIST"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

var (
	entities = map[string]LogRequest_Entities{
		ENTITY_USER: LogRequest_USER,
		ENTITY_POST: LogRequest_POST,
	}

	actions = map[string]LogRequest_Actions{
		ACTION_CREATE:   LogRequest_CREATE,
		ACTION_UPDATE:   LogRequest_UPDATE,
		ACTION_GET:      LogRequest_GET,
		ACTION_DELETE:   LogRequest_DELETE,
		ACTION_LIST:     LogRequest_LIST,
		ACTION_REGISTER: LogRequest_REGISTER,
		ACTION_LOGIN:    LogRequest_LOGIN,
	}
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	UserID    int64     `bson:"user_id"`
	Timestamp time.Time `bson:"timestamp"`
}

func ToPbEntity(entity string) (LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, errors.New("invalid entity")
	}

	return val, nil
}

func ToPbAction(action string) (LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, errors.New("invalid action")
	}

	return val, nil
}
