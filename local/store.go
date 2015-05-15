package local

import "gopkg.in/natefinch/lumberjack.v2"

type MessageStore struct {
	logger *lumberjack.Logger
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		logger: &lumberjack.Logger{
			Filename:   "/var/log/myapp/foo.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
		},
	}
}

// Consider the following for binary marshaling
// https://golang.org/pkg/encoding/gob/
// https://developers.google.com/protocol-buffers/

//Starting Place for On-disk format of a message from Kafka
//
//message length : 4 bytes (value: 1+4+n)
//"magic" value  : 1 byte    # version
//crc            : 4 bytes   #crc of payload
//payload        : n bytes

type Message struct {
}

func (m *MessageStore) Write(message *Message) {
	//TODO: serialize message to bytes
	data := []byte{}
	m.logger.Write(data)
}
