package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type AdminOperationParams struct {
	AdminEmail    string `json:"admin_email"`
	AdminPassword string `json:"admin_pass"`
	UserEmail     string `json:"user_email"`
	UserPassword  string `json:"user_pass"`
}

type UserBanParams struct {
	AdminEmail    string `json:"admin_email"`
	AdminPassword string `json:"admin_pass"`
	UserEmail     string `json:"user_email"`
	UserPassword  string `json:"user_pass"`
	Reason        string `json:"reason"`
}
type BanLog struct {
	IsBan bool
	Who   string
	Why   string
	When  time.Time
}

func (u *UserService) Validate(w http.ResponseWriter, r *http.Request, email string, pass string) (bool, User) {

	passwordDigest := md5.New().Sum([]byte(pass))
	user, err := u.repository.Get(email)

	if err != nil {
		handleError(err, w)
		return false, User{}
	}
	if string(passwordDigest) != user.PasswordDigest {
		handleError(errors.New("invalid login params"), w)
		return false, User{}
	}

	return true, user
}

func isHierarchy(user string, admin string, w http.ResponseWriter) bool {

	if user == "admin" && admin == "admin" {
		handleAccessError(errors.New("admin cannot perform operation on another admin"), w)
		return false
	}
	if user == "superadmin" {
		handleAccessError(errors.New("nobody can perform operation on the superadmin"), w)
		return false
	}
	if admin == "user" {
		handleAccessError(errors.New("user cannot perform admin`s operations"), w)
		return false
	}
	return true
}

func handleAccessError(err error, w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Write([]byte(err.Error()))
}

func (u *UserService) AdminPromote(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	params := &AdminOperationParams{}
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

	isAdmin, admin := u.Validate(w, r, params.AdminEmail, params.AdminPassword)
	isUser, user := u.Validate(w, r, params.UserEmail, params.UserPassword)

	if !isAdmin || !isUser {
		return
	}

	if !isHierarchy(user.Role, admin.Role, w) {
		return
	}

	if user.IsBan {
		handleError(errors.New("user is banned"), w)
		return
	}

	user.Role = "admin"

	u.repository.Update(user.Email, user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user was succesfully promoted"))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/admin/promote").Observe(duration.Seconds())
}

func (u *UserService) AdminFire(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	params := &AdminOperationParams{}
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

	isAdmin, admin := u.Validate(w, r, params.AdminEmail, params.AdminPassword)
	isUser, user := u.Validate(w, r, params.UserEmail, params.UserPassword)

	if !isAdmin || !isUser {
		return
	}

	if admin.Role != "superadmin" {
		handleAccessError(errors.New("only superadmin can fire admins"), w)
		return
	}

	if user.Role != "admin" {
		handleAccessError(errors.New("only admins can be fired"), w)
		return
	}

	user.Role = "user"

	u.repository.Update(user.Email, user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("admin was successfully fired"))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/admin/fire").Observe(duration.Seconds())
}

func (u *UserService) UserBan(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	params := &UserBanParams{}
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

	isAdmin, admin := u.Validate(w, r, params.AdminEmail, params.AdminPassword)
	isUser, user := u.Validate(w, r, params.UserEmail, params.UserPassword)

	if !isAdmin || !isUser {
		return
	}

	if !isHierarchy(user.Role, admin.Role, w) {
		return
	}

	if user.Role != "user" {
		handleAccessError(errors.New("only user can be banned"), w)
		return
	}

	if user.IsBan {
		handleError(errors.New("user is already banned"), w)
		return
	}

	log := BanLog{
		IsBan: true,
		Who:   admin.Email,
		Why:   params.Reason,
		When:  time.Now(),
	}

	user.IsBan = true
	user.BanHistory = append(user.BanHistory, log)

	u.repository.Update(user.Email, user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user was succesfully banned"))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/admin/ban").Observe(duration.Seconds())
}

func (u *UserService) UserUnban(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	params := &AdminOperationParams{}
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		handleError(errors.New("could not read params"), w)
		return
	}

	isAdmin, admin := u.Validate(w, r, params.AdminEmail, params.AdminPassword)
	isUser, user := u.Validate(w, r, params.UserEmail, params.UserPassword)

	if !isAdmin || !isUser {
		return
	}

	if !isHierarchy(user.Role, admin.Role, w) {
		return
	}

	if user.Role != "user" {
		handleAccessError(errors.New("only user can be banned"), w)
		return
	}

	if !user.IsBan {
		handleError(errors.New("user is not banned"), w)
		return
	}

	log := BanLog{
		IsBan: false,
		Who:   admin.Email,
		Why:   "",
		When:  time.Now(),
	}

	user.IsBan = false
	user.BanHistory = append(user.BanHistory, log)

	u.repository.Update(user.Email, user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user was succesfully unbanned"))
	duration := time.Since(startTime)
	responseTimeHistogram.WithLabelValues("/admin/unban").Observe(duration.Seconds())
}

func (j *JWTService) inspect(users UserRepository, h ProtectedHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		adminToken := r.Header.Get("Admin")
		userToken := r.Header.Get("User")

		adminParsed, err := j.ParseJWT(adminToken)
		if err != nil {
			handleAccessError(errors.New("admin unauthorized"), rw)
			return
		}
		admin, err := users.Get(adminParsed.Email)

		if err != nil {
			handleAccessError(errors.New("admin unauthorized"), rw)
			return
		}

		if admin.Role == "user" {
			handleAccessError(errors.New("access forbidden"), rw)
			return
		}

		user, err := users.Get(userToken)
		if err != nil {
			handleAccessError(errors.New("no such user"), rw)
			return
		}

		h(rw, r, user)
	}
}
