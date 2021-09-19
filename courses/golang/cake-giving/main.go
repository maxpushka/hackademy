package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func getCakeHandler(w http.ResponseWriter, r *http.Request, u User) {
	startTime := time.Now()

	if u.IsBan {
		handleAccessError(errors.New("you were banned"), w)
		return
	}
	w.Write([]byte(u.FavoriteCake))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/cake").Observe(duration.Seconds())

	// getCakeTime.Observe(float64())
	numberOfCakesGiven.Inc()
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request, u User) {
	startTime := time.Now()

	var history string
	history += "User " + u.Email + " history\n"
	for i := 0; i < len(u.BanHistory); i++ {
		log := u.BanHistory[i]
		if log.IsBan {
			history += "BANNED\n"
		} else {
			history += "UNBANNED\n"
		}
		history += "	by " + log.Who + "\n"
		history += "	at " + log.When.Format("2006-01-02 15:04:05") + "\n"
		if log.IsBan {
			history += "	reason: " + log.Why + "\n"
		}
	}
	w.Write([]byte(history))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/admin/inspect").Observe(duration.Seconds())
}

func main() {
	initEnv()

	go sender()
	go metrics()

	flag.Parse()
	hub := newHub()
	go hub.run()
	r := mux.NewRouter()
	users := NewInMemoryUserStorage()
	userService := UserService{repository: users}

	jwtService, err := NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/cake", logRequest(jwtService.jwtAuth(users, getCakeHandler))).Methods(http.MethodGet)
	r.HandleFunc("/user/register", logRequest(userService.Register)).Methods(http.MethodPost)
	r.HandleFunc("/user/jwt", logRequest(wrapJwt(jwtService, userService.JWT))).Methods(http.MethodPost)

	r.HandleFunc("/user/me", logRequest(userService.ShowMyCake)).Methods(http.MethodGet)
	r.HandleFunc("/user/favorite_cake", logRequest(userService.ChangeCake)).Methods(http.MethodPost)
	r.HandleFunc("/user/email", logRequest(userService.ChangeEmail)).Methods(http.MethodPost)
	r.HandleFunc("/user/password", logRequest(userService.ChangePassword)).Methods(http.MethodPost)

	r.HandleFunc("/admin/promote", logRequest(userService.AdminPromote)).Methods(http.MethodGet)
	r.HandleFunc("/admin/fire", logRequest(userService.AdminFire)).Methods(http.MethodGet)
	r.HandleFunc("/admin/ban", logRequest(userService.UserBan)).Methods(http.MethodGet)
	r.HandleFunc("/admin/unban", logRequest(userService.UserUnban)).Methods(http.MethodGet)
	r.HandleFunc("/admin/inspect", logRequest(jwtService.inspect(users, getHistoryHandler))).Methods(http.MethodGet)

	//JWT=$(curl -X POST localhost:8000/user/jwt --data '{"email":"admin@mail.com","password":"admin1111"}')
	//wscat -c ws://localhost:8000/ws -H Authorization:$JWT
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r, jwtService, *users)
	})

	srv := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {

		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	log.Println("Server started, hit Ctrl+C to stop")
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}
	log.Println("Good bye :)")
}

func wrapJwt(jwt *JWTService, f func(http.ResponseWriter, *http.Request, *JWTService)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}

func initEnv() {
	os.Setenv("CAKE_ADMIN_EMAIL", "admin@mail.com")
	os.Setenv("CAKE_ADMIN_PASSWORD", "admin1111")
}
