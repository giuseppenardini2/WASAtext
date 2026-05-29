package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.doLogin)
	rt.router.GET("/users", rt.getUsers)
	rt.router.GET("/users/:userid", rt.getMyConversations)
	rt.router.PUT("/users/:userid/username", rt.setMyUserName)
	rt.router.PUT("/users/:userid/picture", rt.setMyPhoto)
	rt.router.GET("/users/:userid/conversations", rt.getConversation)
	rt.router.POST("/users/:userid/conversations/:conversationid/messages", rt.sendMessage)
	rt.router.POST("/users/:userid/conversations/:conversationid/messages/:messageid", rt.forwardMessage)
	rt.router.DELETE("/users/:userid/conversations/:conversationid/messages/:messageid", rt.deleteMessage)
	rt.router.POST("/users/:userid/conversations/:conversationid/messages/:messageid/reactions", rt.commentMessage)
	rt.router.DELETE("/users/:userid/conversations/:conversationid/messages/:messageid/reactions/:reactionid", rt.uncommentMessage)

	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
