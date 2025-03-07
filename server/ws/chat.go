package ws

import (
	"errors"
	"log/slog"
	"vezgammon/server/types"

	"github.com/gorilla/websocket"
)

var ErrWritingMessage = errors.New("Error writing message")

type Database interface {
	GetCurrentGame(int64) (*types.ReturnGame, error)
	GetGame(int64) (*types.Game, error)
}

var db Database

func Init(database Database) {
	db = database
}

func chat(conn *websocket.Conn, userID int64) {
	defer conn.Close()
	var m Message
	for {
		err := conn.ReadJSON(&m)

		if err != nil {
			slog.With("err", err).Error("Error reading message")

			if websocket.IsCloseError(err, websocket.CloseGoingAway) ||
				websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				slog.Debug("Normal closure")
				value, ok := disconnect.Load(userID)
				if !ok {
					break
				}
				f := value.(func(int64) error)
				slog.Debug("Removing player from matchmaking queue")

				err := f(userID)
				slog.With("err", err).Error("Removing player from matchmaking queue")
			}
			break
		}

		err = chatRespondeToMessage(userID, m)
		if err != nil {
			slog.With("err", err).Error("Error responding to message")
			break
		}
	}
}

func chatRespondeToMessage(userID int64, m Message) error {
	slog.With("message", m).Debug("Received message")

	rg, err := db.GetCurrentGame(userID)
	if err != nil {
		slog.With("err", err).Error("Error getting current game in chat")
		return nil
	}

	g, err := db.GetGame(rg.ID)
	if err != nil {
		slog.With("err", err).Error("Error getting game in chat")
		return nil
	}

	var oppoonentId int64
	if g.Player1 == userID {
		oppoonentId = g.Player2
	} else {
		oppoonentId = g.Player1
	}

	err = SendMessage(oppoonentId, m)

	if err != nil {
		return ErrWritingMessage
	}
	return nil
}
