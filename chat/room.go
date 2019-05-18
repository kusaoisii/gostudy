package main

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
	// joinはチャットルームに参加しているとしているクライアントのためのチャンネルでです
	join chan *client
	// leaveはチャットルームから退席しようとしているクライアントのためのチャネルです。
	leave chan *client
	// clientsには在籍している全てのクライアントが保持されます。
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true

		case client := <-r.leave:
			//退室
			delete(r.clients, client)
			close(client.send)
		case msg := <- r.forward:
			// 全てのクライアントにメッセエージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					//　メッセージ
				default:
					//　送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}

		}
	}
}