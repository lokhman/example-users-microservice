package common

import (
	"encoding/json"

	"github.com/nsqio/go-nsq"
)

// Publishes data to NSQ by specified topic.
func NSQPublish(p *nsq.Producer, topic string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// may use async method (depending on the requirements)
	return p.Publish(topic, body)
}
