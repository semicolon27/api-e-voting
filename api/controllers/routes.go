package controllers

import "github.com/semicolon27/api-e-voting/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/admin/login", middlewares.SetMiddlewareJSON(s.LoginAdmin, "LoginAdmin")).Methods("POST")
	s.Router.HandleFunc("/participant/login", middlewares.SetMiddlewareJSON(s.LoginParticipant, "LoginParticipant")).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/admins", middlewares.SetMiddlewareAdminAuthentication(s.GetAdmins, "GetAdmins")).Methods("GET")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.GetAdmin, "GetAdmin")).Methods("GET")
	s.Router.HandleFunc("/admin", middlewares.SetMiddlewareAdminAuthentication(s.CreateAdmin, "CreateAdmin")).Methods("POST")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateAdmin, "UpdateAdmin")).Methods("PUT")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteAdmin, "DeleteAdmin")).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/candidates", middlewares.SetMiddlewareJSON(s.GetCandidates, "GetCandidates")).Methods("GET")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareJSON(s.GetCandidate, "GetCandidate")).Methods("GET")
	s.Router.HandleFunc("/candidate", middlewares.SetMiddlewareAdminAuthentication(s.CreateCandidate, "CreateCandidate")).Methods("POST")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateCandidate, "UpdateCandidate")).Methods("PUT")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteCandidate, "DeleteCandidate")).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/missions", middlewares.SetMiddlewareJSON(s.GetMissions, "GetMissions")).Methods("GET")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareJSON(s.GetMission, "GetMission")).Methods("GET")
	s.Router.HandleFunc("/mission", middlewares.SetMiddlewareAdminAuthentication(s.CreateMission, "CreateMission")).Methods("POST")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateMission, "UpdateMission")).Methods("PUT")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteMission, "DeleteMission")).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/participants", middlewares.SetMiddlewareJSON(s.GetParticipants, "GetParticipants")).Methods("GET")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareJSON(s.GetParticipant, "GetParticipant")).Methods("GET")
	s.Router.HandleFunc("/participant", middlewares.SetMiddlewareAdminAuthentication(s.CreateParticipant, "CreateParticipant")).Methods("POST")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateParticipant, "UpdateParticipant")).Methods("PUT")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteParticipant, "DeleteParticipant")).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/visions", middlewares.SetMiddlewareJSON(s.GetVisions, "GetVisions")).Methods("GET")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareJSON(s.GetVision, "GetVision")).Methods("GET")
	s.Router.HandleFunc("/vision", middlewares.SetMiddlewareAdminAuthentication(s.CreateVision, "CreateVision")).Methods("POST")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareAdminAuthentication(s.UpdateVision, "UpdateVision")).Methods("PUT")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareAdminAuthentication(s.DeleteVision, "DeleteVision")).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.GetVotes, "GetVotes")).Methods("GET")
	s.Router.HandleFunc("/votes/count", middlewares.SetMiddlewareJSON(s.GetCountVotes, "GetCountVotes")).Methods("GET")
	s.Router.HandleFunc("/vote/{id}", middlewares.SetMiddlewareJSON(s.GetVote, "GetVote")).Methods("GET")
	s.Router.HandleFunc("/vote", middlewares.SetMiddlewareAuthentication(s.CreateVote, "CreateVote")).Methods("POST")

	// //Users routes
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
