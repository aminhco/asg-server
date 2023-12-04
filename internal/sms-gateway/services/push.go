package services

import (
	"context"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type PushService struct {
	CredentialsJSON string

	client *messaging.Client
	mux    sync.Mutex
}

type PushServiceConfig struct {
	CredentialsJSON string
}

func NewPushService(config PushServiceConfig) *PushService {
	return &PushService{
		CredentialsJSON: config.CredentialsJSON,
	}
}

// init
func (s *PushService) init(ctx context.Context) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.client != nil {
		return
	}

	opt := option.WithCredentialsJSON([]byte(s.CredentialsJSON))

	var app *firebase.App
	app, err = firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return
	}

	s.client, err = app.Messaging(ctx)

	return
}

// send
func (s *PushService) Send(ctx context.Context, token string, data map[string]string) error {
	if err := s.init(ctx); err != nil {
		return err
	}

	_, err := s.client.Send(ctx, &messaging.Message{
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		Token: token,
	})

	return err
}
