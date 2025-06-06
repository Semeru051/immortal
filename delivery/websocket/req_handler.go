package websocket

import (
	"fmt"
	"slices"

	"github.com/starrysilk/immortal/pkg/utils"
	"github.com/starrysilk/immortal/types"
	"github.com/starrysilk/immortal/types/message"
	"github.com/gorilla/websocket"
)

// handleReq handles new incoming REQ messages from client.
func (s *Server) handleReq(conn *websocket.Conn, m message.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer measureLatency(s.metrics.RequestLatency)()

	status := success
	defer func() {
		s.metrics.RequestsTotal.WithLabelValues(status).Inc()
	}()

	msg, ok := m.(*message.Req)
	if !ok {
		_ = conn.WriteMessage(1, message.MakeNotice("error: can't parse REQ message."))

		status = parseFail

		return
	}

	client, ok := s.conns[conn]
	if !ok {
		_ = conn.WriteMessage(1, message.MakeNotice(fmt.Sprintf("error: can't find connection %s",
			conn.RemoteAddr())))

		status = serverFail

		return
	}

	if (s.config.GetLimitation().AuthRequired ||
		slices.Contains(msg.Filter.Kinds, types.KindGiftWrap)) &&
		!*client.isKnown {
		client.challenge = utils.GenerateChallenge(10)
		authm := message.MakeAuth(client.challenge)

		closem := message.MakeClosed(
			msg.SubscriptionID,
			"auth-required: we can't serve notes to unauthenticated users",
		)

		_ = conn.WriteMessage(1, closem)

		_ = conn.WriteMessage(1, authm)
		status = authFail

		return
	}

	if len(msg.SubscriptionID) >= int(s.config.GetLimitation().MaxSubidLength) {
		_ = conn.WriteMessage(1, message.MakeNotice(fmt.Sprintf("error: max limit of sub id is: %d",
			s.config.GetLimitation().MaxSubidLength)))

		status = limitsFail

		return
	}

	if len(client.subs) >= int(s.config.GetLimitation().MaxSubscriptions) {
		_ = conn.WriteMessage(1, message.MakeNotice(fmt.Sprintf("error: max limit of subs is: %d",
			s.config.GetLimitation().MaxSubscriptions)))

		status = limitsFail

		return
	}

	res, err := s.handler.HandleReq(&msg.Filter, *client.pubkey)
	if err != nil {
		_ = conn.WriteMessage(1, message.MakeNotice(fmt.Sprintf("error: can't process REQ message: %s", err.Error())))
		status = databaseFail

		return
	}

	for _, e := range res {
		msg := message.MakeEvent(msg.SubscriptionID, &e)
		_ = conn.WriteMessage(1, msg)
	}

	_ = conn.WriteMessage(1, message.MakeEOSE(msg.SubscriptionID))

	client.Lock()
	s.metrics.Subscriptions.Inc()
	client.subs[msg.SubscriptionID] = msg.Filter
	client.Unlock()
}
