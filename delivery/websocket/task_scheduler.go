package websocket

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/starrysilk/immortal/pkg/logger"
	"github.com/starrysilk/immortal/types"
)

const ExpirationTaskListName = "expiration_events"

func (s *Server) checkExpiration() { //nolint
	for range time.Tick(time.Minute) {
		tasks, err := s.redis.GetReadyTasks(ExpirationTaskListName)
		if err != nil {
			_, err := s.grpc.AddLog(context.Background(),
				"redis error while receiving ready tasks", err.Error())
			if err != nil {
				logger.Error("can't send log to manager", "err", err)
			}

			continue
		}

		failedTasks := make([]string, 0)

		if len(tasks) != 0 {
			for _, task := range tasks {
				data := strings.Split(task, ":")

				if len(data) != 2 {
					continue
				}

				kind, err := strconv.ParseUint(data[1], 10, 16)
				if err != nil {
					continue
				}

				// lint error is g115 gosec rule. this rule is broken.
				// see: https://github.com/securego/gosec/issues/1288
				// todo:: update gosec linter and remove //nolint comment.
				if err := s.handler.DeleteByID(data[0],
					types.Kind(kind)); err != nil { //nolint
					failedTasks = append(failedTasks, task)
				}
			}
		}

		if len(failedTasks) != 0 {
			for _, ft := range failedTasks {
				if err := s.redis.AddDelayedTask(ExpirationTaskListName,
					ft, time.Minute*10); err != nil {
					continue
				}
			}
		}
	}
}
