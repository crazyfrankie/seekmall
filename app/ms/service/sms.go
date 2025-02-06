package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/crazyfrankie/seekmall/app/ms/repository"
	"github.com/crazyfrankie/seekmall/app/ms/service/sms"
	"github.com/crazyfrankie/seekmall/config"
)

var (
	SecretKey = "mY1eF3oI6mR7iR2yU0pW0rH1kA2qU1tQ"
)

type SmsSvc interface {
	SendCode(ctx context.Context, phone string) error
	VerifyCode(ctx context.Context, phone, code string) error
}

type SmsService struct {
	repo *repository.SmsRepo
	sms  sms.Service
}

func NewSmsService(repo *repository.SmsRepo, sms sms.Service) SmsSvc {
	return &SmsService{repo: repo, sms: sms}
}

func (s *SmsService) SendCode(ctx context.Context, phone string) error {
	code := generateCode()

	hash := generateHMAC(code, SecretKey)

	err := s.repo.Store(ctx, phone, hash)
	if err != nil {
		return err
	}

	// Send
	err = s.sms.Send(ctx, config.GetConf().SMS.TemplateID, []string{code}, phone)

	return err
}

func (s *SmsService) VerifyCode(ctx context.Context, phone, code string) error {
	encode := generateHMAC(code, SecretKey)

	err := s.repo.Verify(ctx, phone, encode)

	return err
}

func generateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var code strings.Builder
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		code.WriteString(strconv.Itoa(digit))
	}

	return code.String()
}

func generateHMAC(code, key string) string {
	h := hmac.New(sha256.New, []byte(key))

	h.Write([]byte(code))

	return hex.EncodeToString(h.Sum(nil))
}
