package controllers

import (
	"net/http"

	"github.com/semicolon27/api-e-voting/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/api/admin/login", middlewares.SetMiddlewareJSON(s.LoginAdmin, "LoginAdmin")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/participant/login", middlewares.SetMiddlewareJSON(s.LoginParticipant, "LoginParticipant")).Methods("POST", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/admins", middlewares.SetMiddlewareAdminAuthentication(s.GetAdmins, "GetAdmins")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.GetAdmin, "GetAdmin")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/admin", middlewares.SetMiddlewareAdminAuthentication(s.CreateAdmin, "CreateAdmin")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateAdmin, "UpdateAdmin")).Methods("PUT", http.MethodOptions)
	s.Router.HandleFunc("/api/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteAdmin, "DeleteAdmin")).Methods("DELETE", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/candidates", middlewares.SetMiddlewareJSON(s.GetCandidates, "GetCandidates")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/candidate/{id}", middlewares.SetMiddlewareJSON(s.GetCandidate, "GetCandidate")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/candidate", middlewares.SetMiddlewareAdminAuthentication(s.CreateCandidate, "CreateCandidate")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/candidate/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateCandidate, "UpdateCandidate")).Methods("PUT", http.MethodOptions)
	s.Router.HandleFunc("/api/candidate/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteCandidate, "DeleteCandidate")).Methods("DELETE", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/missions", middlewares.SetMiddlewareJSON(s.GetMissions, "GetMissions")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/mission/{id}", middlewares.SetMiddlewareJSON(s.GetMission, "GetMission")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/mission", middlewares.SetMiddlewareAdminAuthentication(s.CreateMission, "CreateMission")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/mission/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateMission, "UpdateMission")).Methods("PUT", http.MethodOptions)
	s.Router.HandleFunc("/api/mission/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteMission, "DeleteMission")).Methods("DELETE", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/participants", middlewares.SetMiddlewareJSON(s.GetParticipants, "GetParticipants")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/participant/{id}", middlewares.SetMiddlewareJSON(s.GetParticipant, "GetParticipant")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/participant", middlewares.SetMiddlewareAdminAuthentication(s.CreateParticipant, "CreateParticipant")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/participant/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateParticipant, "UpdateParticipant")).Methods("PUT", http.MethodOptions)
	s.Router.HandleFunc("/api/participant/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteParticipant, "DeleteParticipant")).Methods("DELETE", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/visions", middlewares.SetMiddlewareJSON(s.GetVisions, "GetVisions")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/vision/{id}", middlewares.SetMiddlewareJSON(s.GetVision, "GetVision")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/vision", middlewares.SetMiddlewareAdminAuthentication(s.CreateVision, "CreateVision")).Methods("POST", http.MethodOptions)
	s.Router.HandleFunc("/api/vision/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateVision, "UpdateVision")).Methods("PUT", http.MethodOptions)
	s.Router.HandleFunc("/api/vision/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteVision, "DeleteVision")).Methods("DELETE", http.MethodOptions)

	//Users routes
	s.Router.HandleFunc("/api/votes", middlewares.SetMiddlewareJSON(s.GetVotes, "GetVotes")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/votes/count", middlewares.SetMiddlewareJSON(s.GetCountVotes, "GetCountVotes")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/vote/{id}", middlewares.SetMiddlewareJSON(s.GetVote, "GetVote")).Methods("GET", http.MethodOptions)
	s.Router.HandleFunc("/api/vote", middlewares.SetMiddlewareAuthentication(s.CreateVote, "CreateVote")).Methods("POST", http.MethodOptions)

	// //Users routes
	// s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Posts routes
	// s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
