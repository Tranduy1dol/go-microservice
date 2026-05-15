package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/domain"
	"github.com/Tranduy1dol/kotoba-press-core/internal/logger"
	"github.com/Tranduy1dol/kotoba-press-core/internal/port"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var log = logger.New(logger.ComponentAuth)

type GoogleOAuthService struct {
	config   *oauth2.Config
	jwtSvc   *JWTService
	userRepo port.UserRepository
}

func NewGoogleOAuthService(cfg config.OAuthConfig, jwtSvc *JWTService, userRepo port.UserRepository) *GoogleOAuthService {
	return &GoogleOAuthService{
		config: &oauth2.Config{
			ClientID:     cfg.GoogleClientID,
			ClientSecret: cfg.GoogleClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
		jwtSvc:   jwtSvc,
		userRepo: userRepo,
	}
}

func (s *GoogleOAuthService) GetAuthUrl() (string, error) {
	state, err := RandToken()
	if err != nil {
		return "", err
	}

	return s.config.AuthCodeURL(state), nil
}

func (s *GoogleOAuthService) HandleCallback(ctx context.Context, code string) (string, *domain.User, error) {
	token, err := s.config.Exchange(ctx, code)
	if err != nil {
		log.Error("failed to exchange oauth code", "error", err)
		return "", nil, err
	}

	client := s.config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Error("failed to fetch user info from google", "error", err)
		return "", nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var userInfo struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return "", nil, err
	}

	user, err := s.userRepo.Upsert(ctx, userInfo.ID, userInfo.Email, userInfo.Name, userInfo.Picture)
	if err != nil {
		return "", nil, err
	}

	jwtToken, err := s.jwtSvc.IssueToken(user.ID, user.Role)
	if err != nil {
		log.Error("failed to issue jwt token", "user_id", user.ID, "error", err)
		return "", nil, err
	}

	log.Info("user successfully authenticated via google", "user_id", user.ID, "email", user.Email, "role", user.Role)
	return jwtToken, user, nil
}

func RandToken() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
