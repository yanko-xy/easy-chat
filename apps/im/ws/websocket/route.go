/**
 * @author: Yanko/xiaoxiaoyang-sheep
 * @doc:
 **/

package websocket

type Route struct {
	Method  string
	Handler HandlerFunc
}

type HandlerFunc func(srv *Server, conn *Conn, msg *Message)
