package controllers

import "github.com/semicolon27/api-e-voting/api/middlewares"

func (s *Server) initializeRoutes() {

	// // Home Route
	// s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/participant/login", middlewares.SetMiddlewareJSON(s.GetCandidate)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/candidates", middlewares.SetMiddlewareJSON(s.GetCandidates)).Methods("GET")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareJSON(s.GetCandidate)).Methods("GET")
	s.Router.HandleFunc("/candidate", middlewares.SetMiddlewareAuthentication(s.CreateCandidate)).Methods("POST")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateCandidate)).Methods("PUT")
	s.Router.HandleFunc("/candidate/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCandidate)).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/classes", middlewares.SetMiddlewareJSON(s.GetClasses)).Methods("GET")
	s.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareJSON(s.GetClass)).Methods("GET")
	s.Router.HandleFunc("/class", middlewares.SetMiddlewareAuthentication(s.CreateClass)).Methods("POST")
	s.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateClass)).Methods("PUT")
	s.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteClass)).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/missions", middlewares.SetMiddlewareJSON(s.GetMissions)).Methods("GET")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareJSON(s.GetMission)).Methods("GET")
	s.Router.HandleFunc("/mission", middlewares.SetMiddlewareAuthentication(s.CreateMission)).Methods("POST")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateMission)).Methods("PUT")
	s.Router.HandleFunc("/mission/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteMission)).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/participants", middlewares.SetMiddlewareJSON(s.GetParticipants)).Methods("GET")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareJSON(s.GetParticipant)).Methods("GET")
	s.Router.HandleFunc("/participant", middlewares.SetMiddlewareAuthentication(s.CreateParticipant)).Methods("POST")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateParticipant)).Methods("PUT")
	s.Router.HandleFunc("/participant/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteParticipant)).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/visions", middlewares.SetMiddlewareJSON(s.GetVisions)).Methods("GET")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareJSON(s.GetVision)).Methods("GET")
	s.Router.HandleFunc("/vision", middlewares.SetMiddlewareAuthentication(s.CreateVision)).Methods("POST")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateVision)).Methods("PUT")
	s.Router.HandleFunc("/vision/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteVision)).Methods("DELETE")

	//Users routes
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.GetVotes)).Methods("GET")
	s.Router.HandleFunc("/vote/{id}", middlewares.SetMiddlewareJSON(s.GetVote)).Methods("GET")
	s.Router.HandleFunc("/vote", middlewares.SetMiddlewareAuthentication(s.CreateVote)).Methods("POST")
	// s.Router.HandleFunc("/vote/count", middlewares.SetMiddlewareAuthentication(s.GetVoteCount)).Methods("GET")

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
