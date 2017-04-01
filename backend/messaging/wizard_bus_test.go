package messaging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBus(t *testing.T) {
	b := &WizardBus{}
	b.Start()

	err := b.Publish("topic", []byte("message1"))
	assert.NoError(t, err)
	// should be able to publish with no subscribers

	c1 := make(chan []byte, 3)
	c2 := make(chan []byte, 3)
	assert.NoError(t, b.Subscribe("topic", "consumer1", c1))

	assert.NoError(t, b.Subscribe("topic", "consumer2", c2))

	err = b.Publish("topic", []byte("message2"))
	assert.NoError(t, err)
	err = b.Publish("topic", []byte("message3"))
	assert.NoError(t, err)

	b.Stop()

	received := 0
	messages := []string{}
	for m := range c1 {
		received++
		messages = append(messages, string(m))
	}

	for m := range c2 {
		received++
		messages = append(messages, string(m))
	}
	assert.Equal(t, "message2", messages[0])
	assert.Equal(t, "message3", messages[3])
	assert.Equal(t, 4, len(messages))
	assert.Equal(t, 4, received)
}
