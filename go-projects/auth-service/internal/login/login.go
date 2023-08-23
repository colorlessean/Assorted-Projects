package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/database"
	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/login/session"
	"github.com/gorilla/mux"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Server interface {
    Start() error
}

type LoginServer struct {
    db database.Database
}

func NewLoginServer(db database.Database) (Server, error) {
    return &LoginServer{
        db: db,
    }, nil
}

func (ls *LoginServer) Start() error {
    router := mux.NewRouter()

    router.HandleFunc("/login", ls.loginHandler).Methods("POST")
    router.HandleFunc("/testlogin", ls.testLoginHanlder).Methods("GET")
    router.HandleFunc("/logout", ls.logoutHandler).Methods("GET")

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalf("server failed to start up: %v", err)
        return err
    }

    return nil
}

func (ls *LoginServer) loginHandler(w http.ResponseWriter, r *http.Request) {
    credentials := &Credentials{}

    err := json.NewDecoder(r.Body).Decode(credentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    session, err := ls.login(credentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    sessionCookie := &http.Cookie{
    	Name:       "_session_id",
    	Value:      session.SessionID,
    	Path:       "/",
    	Expires:    session.Expiry, 
    	Secure:     false,
    	HttpOnly:   false,
    	SameSite:   0,
    }

    http.SetCookie(w, sessionCookie)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    message := "successfully processed login request"
    json.NewEncoder(w).Encode(map[string]string{"message":message})
}

// login takes in credentials and returns a sessionID and/or error
func (ls *LoginServer) login(c *Credentials) (*session.Session, error) {
    gotUser, err := ls.db.GetUser(c.Username)
    if err != nil {
        return nil, err
    }

    session, err := session.NewSession()
    if err != nil {
        return nil, fmt.Errorf("session could not be created: %v", err)
    }
    
    return session, nil
}

func (ls *LoginServer) testLoginHanlder(w http.ResponseWriter, r *http.Request) {
    // yet to implement
}

func (ls *LoginServer) logoutHandler(w http.ResponseWriter, r *http.Request) {
    // yet to implement
}

