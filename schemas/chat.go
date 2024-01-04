package schemas

import (
	"github.com/sashabaranov/go-openai"
	"gopkg.in/mgo.v2/bson"
)

type Chat struct {
	ID      bson.ObjectId                  `bson:"_id,omitempty"`
	User    User                           `bson:"user"`
	Name    string                         `bson:"name"`
	Content []openai.ChatCompletionMessage `bson:"content"`
}
