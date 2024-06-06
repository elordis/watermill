package gochannel_test

import (
	"testing"

	"github.com/elordis/watermill/pubsub/gochannel"

	"github.com/elordis/watermill/pubsub/tests"

	"github.com/elordis/watermill"
	"github.com/elordis/watermill/message"
)

func BenchmarkSubscriber(b *testing.B) {
	tests.BenchSubscriber(b, func(n int) (message.Publisher, message.Subscriber) {
		pubSub := gochannel.NewGoChannel(
			gochannel.Config{OutputChannelBuffer: int64(n)}, watermill.NopLogger{},
		)
		return pubSub, pubSub
	})
}

func BenchmarkSubscriberPersistent(b *testing.B) {
	tests.BenchSubscriber(b, func(n int) (message.Publisher, message.Subscriber) {
		pubSub := gochannel.NewGoChannel(
			gochannel.Config{
				OutputChannelBuffer: int64(n),
				Persistent:          true,
			},
			watermill.NopLogger{},
		)
		return pubSub, pubSub
	})
}
