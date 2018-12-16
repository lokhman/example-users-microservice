package api

import (
	"github.com/jinzhu/gorm"
	"github.com/nsqio/go-nsq"
)

// Sort of a controller.
type API struct {
	// database connection instance
	DB *gorm.DB

	// NSQ producer instance
	NSQ *nsq.Producer
}
