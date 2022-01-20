package main

import (
	"context"
	"golang.org/x/sync/singleflight"
)

type service struct {
	requestGroup singleflight.Group
}

func (s *service) handleRequest(ctx context.Context, request Request) (Response, error) {
	v, err, _ := requestGroup.Do(request.Hash(), func() (interface{}, error) {
		rows, err := // select * from tables
		if err != nil {
			return nil, err
		}
		return rows, nil
	})
	if err != nil {
		return nil, err
	}
	return Response{
		rows: rows,
	}, nil
}
