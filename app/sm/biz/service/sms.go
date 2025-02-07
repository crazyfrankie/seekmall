package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/crazyfrankie/seekmall/app/sm/biz/repository"
	"github.com/crazyfrankie/seekmall/app/sm/biz/service/sms"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/crazyfrankie/seekmall/config"
	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
)

var (
	SecretKey = "mY1eF3oI6mR7iR2yU0pW0rH1kA2qU1tQ"
)

type SmsServer struct {
	repo *repository.SmsRepo
	sms  sms.Service
	sm.UnimplementedSmsServiceServer
}

func NewSmsServer(repo *repository.SmsRepo, sms sms.Service) *SmsServer {
	return &SmsServer{repo: repo, sms: sms}
}

func (s *SmsServer) RegisterServer(server *grpc.Server) {
	sm.RegisterSmsServiceServer(server, s)
}

func (s *SmsServer) SendSms(ctx context.Context, req *sm.SendSmsRequest) (*sm.SendSmsResponse, error) {
	code := generateCode()
	hash := generateHMAC(code, SecretKey)

	err := s.repo.Store(ctx, req.GetBiz(), req.GetPhone(), hash)
	if err != nil {
		return nil, err
	}

	// Send
	err = s.sms.Send(ctx, config.GetConf().SMS.TemplateID, []string{hash}, req.GetPhone())

	return &sm.SendSmsResponse{}, err
}

func (s *SmsServer) VerifySms(ctx context.Context, req *sm.VerifySmsRequest) (*sm.VerifySmsResponse, error) {
	encode := generateHMAC(req.GetCode(), SecretKey)

	err := s.repo.Verify(ctx, req.GetBiz(), req.GetPhone(), encode)
	if err != nil {
		return nil, err
	}

	return &sm.VerifySmsResponse{}, err
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
