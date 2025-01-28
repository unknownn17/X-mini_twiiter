package services

import (
	"context"
	"log"
	"notification/internal/api/handler"
	"notification/internal/kafka/producer"
	not17 "notification/internal/protos"
)

type Service struct {
	not17.UnimplementedNotificationServer
	W *handler.WebSocket
}

func (u *Service) Notification(ctx context.Context, req *not17.ProduceMessage) (*not17.EMailSendResponse, error) {
	if err := producer.Producer(req.Username, req.Message); err != nil {
		log.Println(err)
		return nil, err
	}
	return &not17.EMailSendResponse{Message: "succesfully"}, nil
}
