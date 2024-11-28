package matchmaking

import (
	"errors"

	"github.com/google/uuid"
)

var users map[int64]string = make(map[int64]string)
var links map[string]int64 = make(map[string]int64)

func GenerateLink(user_id int64) (string, error) {
	uuid := uuid.New()

	suuid := uuid.String()

	users[user_id] = suuid
	links[suuid] = user_id

	return suuid, nil
}

var ErrLinkNotFound = errors.New("Link not found")

func JoinLink(link string, user_id int64) error {
	opp_id, ok := links[link]
	if !ok {
		return ErrLinkNotFound
	}

	delete(links, link)
	delete(users, opp_id)

	return createGame(user_id, opp_id)
}
