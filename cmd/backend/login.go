package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4"
	"github.com/sgielen/vierkantle/pkg/database"
	"github.com/sgielen/vierkantle/pkg/database/gendb"
	pb "github.com/sgielen/vierkantle/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const cookieName = "vkt-auth"

type contextKey string

var vktJwtKey = contextKey("jwt")

// TODO: make these flags or read from env
var jwtKeyId = "1"
var jwtSecret = []byte("746c2cbb43675102284ee4218e1492391ed222ab3743aba2f5002ea8d3c003f2")
var jwtExpiry = 10 * 365 * 24 * 3600 * time.Second /* 10 years */
var cookieMaxAge = 10 * 365 * 24 * 3600            /* 10 years */
var cookieSecure = false
var cookiePath = "/api"
var cookieSameSite = "Strict"

var ErrNotLoggedIn = status.Error(codes.Unauthenticated, "unauthenticated")
var ErrTokenExpired = status.Error(codes.Unauthenticated, "token has expired")
var ErrPermissionDenied = status.Error(codes.PermissionDenied, "permission denied")
var ErrNoSuchUser = status.Error(codes.Unauthenticated, "invalid username/password")

func (s *vierkantleService) Whoami(ctx context.Context, req *pb.WhoamiRequest) (*pb.WhoamiResponse, error) {
	ctx, err := grpcWebAuth(ctx)
	if err != nil {
		return nil, err
	}
	userId, username, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}
	if err := SetCookie(ctx, userId, username); err != nil {
		return nil, err
	}
	return &pb.WhoamiResponse{
		Username: username,
	}, err
}

func (s *vierkantleService) StartLogin(ctx context.Context, req *pb.StartLoginRequest) (*pb.StartLoginResponse, error) {
	// TODO: check if user or e-mail exists in the database
	// if so, take e-mail of that user
	// generate a valid JWT for this user, but do *not* set it as a cookie
	// send the JWT to the user via e-mail
	// implement a feature in the frontend such that it receives the JWT via the URL
	// then sends FinishLogin so that the user is logged in
	return nil, fmt.Errorf("unimplemented")
}

func (s *vierkantleService) FinishLogin(ctx context.Context, req *pb.FinishLoginRequest) (*pb.FinishLoginResponse, error) {
	ctx, err := authenticateWithJwt(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	userId, username, err := AuthUser(ctx)
	if err != nil {
		return nil, err
	}
	if err := SetCookie(ctx, userId, username); err != nil {
		return nil, err
	}
	var res pb.FinishLoginResponse
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		username, err := q.LoginUser(ctx, userId)
		res.Username = username
		return err
	}); err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *vierkantleService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var res pb.RegisterResponse
	var email sql.NullString
	if req.Email != "" {
		email = sql.NullString{Valid: true, String: req.Email}
	}
	if err := database.RunRWTransaction(ctx, pgx.RepeatableRead, func(q *gendb.Queries) error {
		res.Reset()
		userId, err := q.RegisterUser(ctx, gendb.RegisterUserParams{
			Username: req.Username,
			Email:    email,
		})
		if err != nil {
			if database.ToSQLState(err) == "23505" /* Duplicate key value */ {
				return fmt.Errorf("that username or e-mail address is already in use")
			} else {
				return err
			}
		}

		// TODO: if we have an e-mail address, should generate a JWT here, and *e-mail* it to this user
		// then, they can use FinishLogin
		// for now, just immediately log them in here
		if _, err := q.LoginUser(ctx, userId); err != nil {
			return err
		}
		if err := SetCookie(ctx, userId, req.Username); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &res, nil
}

func jwtFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	cookies, ok := md["cookie"]
	if !ok {
		return ""
	}

	header := http.Header{}
	for _, cookie := range cookies {
		header.Add("Cookie", cookie)
	}
	request := http.Request{Header: header}
	for _, cookie := range request.Cookies() {
		if cookie.Name == cookieName {
			return cookie.Value
		}
	}

	return ""
}

func grpcWebAuth(ctx context.Context) (context.Context, error) {
	tokenString := jwtFromContext(ctx)
	if tokenString == "" {
		return ctx, ErrNotLoggedIn
	}
	return authenticateWithJwt(ctx, tokenString)
}

func authenticateWithJwt(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
		}

		if kid, ok := token.Header["kid"]; !ok {
			return nil, status.Error(codes.Unauthenticated, "no key id")
		} else if kidstr, ok := kid.(string); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected key id: %v", kid)
		} else if kidstr != jwtKeyId {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected key id: %s", kidstr)
		}

		return jwtSecret, nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, ErrTokenExpired
	}
	if err != nil {
		log.Printf("jwt token error: %v", err)
		return nil, ErrNotLoggedIn
	}
	if !token.Valid {
		return nil, ErrNotLoggedIn
	}

	ctx = context.WithValue(ctx, vktJwtKey, token)
	return ctx, nil
}

func AuthJWT(ctx context.Context) (*jwt.Token, error) {
	v := ctx.Value(vktJwtKey)
	if v == nil {
		return nil, ErrNotLoggedIn
	}
	if t, ok := v.(*jwt.Token); ok {
		return t, nil
	} else {
		return nil, ErrNotLoggedIn
	}
}

func AuthUser(ctx context.Context) (int64, string, error) {
	t, err := AuthJWT(ctx)
	if err != nil {
		return 0, "", err
	}
	if err := t.Claims.Valid(); err != nil {
		return 0, "", err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); !ok {
		return 0, "", ErrNotLoggedIn
	} else if username, ok := claims["username"]; !ok {
		return 0, "", ErrNotLoggedIn
	} else if userstr, ok := username.(string); !ok {
		return 0, "", ErrNotLoggedIn
	} else if userid, ok := claims["userid"]; !ok {
		return 0, "", ErrNotLoggedIn
	} else if userid64, ok := userid.(int64); ok {
		// this is what should hapen...
		return userid64, userstr, nil
	} else if useridf64, ok := userid.(float64); ok {
		// but this is what happens?
		return int64(useridf64), userstr, nil
	} else {
		return 0, "", ErrNotLoggedIn
	}
}

func SetCookie(ctx context.Context, userid int64, username string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(jwtExpiry).Unix(),
	})
	token.Header["kid"] = jwtKeyId

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	cookie := generateCookie(tokenString)
	header := metadata.Pairs("set-cookie", cookie)

	return grpc.SendHeader(ctx, header)
}

func DeleteCookie(ctx context.Context) error {
	cookie := generateCookie("")
	header := metadata.Pairs("set-cookie", cookie)

	return grpc.SendHeader(ctx, header)
}

func generateCookie(value string) string {
	var cookie []string
	cookie = append(cookie, fmt.Sprintf("%s=%s", cookieName, value))
	cookie = append(cookie, "HttpOnly")
	if cookieSecure {
		cookie = append(cookie, "Secure")
	}
	if cookiePath != "" {
		cookie = append(cookie, "Path="+cookiePath)
	}
	if cookieSameSite != "" {
		cookie = append(cookie, "SameSite="+cookieSameSite)
	}
	if cookieMaxAge > 0 {
		cookie = append(cookie, fmt.Sprintf("Max-Age=%d", cookieMaxAge))
	}
	return strings.Join(cookie, "; ")
}
