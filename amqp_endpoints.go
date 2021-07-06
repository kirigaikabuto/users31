package users31

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type UsersAmqpEndpoints struct {
	usersStore UsersStore
}

func NewUsersAmqpEndpoints(u UsersStore) UsersAmqpEndpoints {
	return UsersAmqpEndpoints{usersStore: u}
}

func (u *UsersAmqpEndpoints) CreateUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		user := &User{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, user)
		if err != nil {
			panic(err)
		}
		newUser, err := u.usersStore.Create(user)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(newUser)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}

func (u *UsersAmqpEndpoints) GetUserByIdAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		user := &User{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, user)
		if err != nil {
			panic(err)
		}
		currentUser, err := u.usersStore.Get(user.Id)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(currentUser)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}

func (u *UsersAmqpEndpoints) GetByUsernameAndPasswordAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		user := &User{}
		jsonData := message.Body
		err := json.Unmarshal(jsonData, user)
		if err != nil {
			panic(err)
		}
		currentUser, err := u.usersStore.GetByUsernameAndPassword(user.Username, user.Password)
		if err != nil {
			panic(err)
		}
		response, err := json.Marshal(currentUser)
		if err != nil {
			panic(err)
		}
		return &amqp.Message{Body: response}
	}
}
