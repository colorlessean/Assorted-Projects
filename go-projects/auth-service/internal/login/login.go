package login

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/colorlessean/Assorted-Projects/go-projects/auth-service/internal/login/session"
	"github.com/gorilla/mux"
)

type Credentials struct {
    username string `json:"username"`
    password string `json:"password"`
}

func NewLoginServer() {
    router := mux.NewRouter()

    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/testlogin", testLoginHanlder).Methods("GET")
    router.HandleFunc("/logout", logoutHandler).Methods("GET")

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalf("server failed to start up: %v", err)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    credentials := &Credentials{}

    err := json.NewDecoder(r.Body).Decode(credentials)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    session, err := login(credentials)
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
func login(c *Credentials) (*session.Session, error) {
    session, err := session.NewSession()
    if err != nil {
        return nil, fmt.Errorf("session could not be created: %v", err)
    }
    
    return session, nil
}

func testLoginHanlder(w http.ResponseWriter, r *http.Request) {
    // yet to implement
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    // yet to implement
}

