package handlers

import (
	"auth-ms/data"
	"auth-ms/utils"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"net/http"
	"time"
)

type AuthProvider struct {
	db       *utils.Repository
	validate *validator.Validate
	log      *log.Logger
}

func NewAuthProvider(repo *utils.Repository, logger *log.Logger) *AuthProvider {
	return &AuthProvider{repo, validator.New(), logger}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Status bool                `json:"status"`
	Msg    string              `json:"msg"`
	Data   data.Tokens         `json:"data,omitempty"`
	Errors utils.ValidateError `json:"errors,omitempty"`
}

func (req *LoginResponse) toJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(req)
	return err
}

func (req *LoginRequest) fromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(&req)
	return err
}

func handleLoginError(err string, w io.Writer) {
	response := LoginResponse{
		Status: false,
		Msg:    err,
	}
	response.toJSON(w)
}

func (auth *AuthProvider) SaveSession(uid, seed string) error {

	res := auth.db.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"seed", "started_at"}),
	}).Create(&data.Session{
		Uid:  uid,
		Seed: seed,
	})
	return res.Error
}

func GenerateTokens(auth AuthProvider, user data.User) (data.Tokens, error) {
	seed := utils.GenerateUID(16)
	var tokens data.Tokens
	accessExpireTime := 5 * time.Minute
	refreshExpireTime := 30 * 24 * 60 * time.Minute

	accessClaims := data.GetClaims(user.Uid, seed, "access", accessExpireTime)
	refreshClaims := data.GetClaims(user.Uid, seed, "refresh", refreshExpireTime)

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessStr, _ := access.SignedString([]byte(viper.GetString("JWT_SECRET")))
	refreshStr, _ := refresh.SignedString([]byte(viper.GetString("JWT_SECRET")))
	tokens = data.Tokens{
		accessStr, refreshStr, user.Uid,
	}
	err := auth.SaveSession(user.Uid, seed)
	return tokens, err
}
func (auth *AuthProvider) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	var response LoginResponse
	err := req.fromJSON(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handleLoginError("Invalid data! Unable to parse the payload", w)
		return
	}

	err = auth.validate.Struct(req)
	errors := utils.NewValidationError()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ValidateErrorFormat
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			*errors = append(*errors, &el)
		}
		response = LoginResponse{
			Status: false,
			Msg:    "Validation Error",
			Errors: *errors,
		}
		response.toJSON(w)
		return
	}

	var dbUser data.User
	err = auth.db.DB.Where("email = ?", req.Email).First(&data.User{}).Scan(&dbUser).Error
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		handleLoginError("Wrong username or password", w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		handleLoginError("Wrong username or password", w)
		return
	}
	tokens, err := GenerateTokens(*auth, dbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handleLoginError("Unexpected server error", w)
		return
	}
	response = LoginResponse{
		Status: true,
		Msg:    "success",
		Data:   tokens,
	}
	response.toJSON(w)
}
