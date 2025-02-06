package handler

import (
	"bytes"
	"encoding/json"
	"github.com/crazyfrankie/seekmall/app/ms/service"
	service2 "github.com/crazyfrankie/seekmall/app/user/service"
	svcmocks "github.com/crazyfrankie/seekmall/biz/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Code int
	Msg  string
	Data any
}

func TestSendCode(t *testing.T) {
	testCases := []struct {
		name string

		reqBody  string
		mock     func(ctrl *gomock.Controller) (service2.UserSvc, service.SmsSvc)
		wantCode int
		wantBody Response
	}{
		{
			name: "success",
			mock: func(ctrl *gomock.Controller) (service2.UserSvc, service.SmsSvc) {
				user, sms := svcmocks.NewMockUserSvc(ctrl), svcmocks.NewMockSmsSvc(ctrl)
				sms.EXPECT().SendCode(gomock.Any(), "13117127070").Return(nil)
				return user, sms
			},
			reqBody: `{
						"phone": "13117127070"
					}`,
			wantCode: 200,
			wantBody: Response{
				Code: 00000,
				Msg:  "success",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			user, sms := tc.mock(ctrl)

			hdl := NewUserHandler(user, sms)

			server := gin.Default()
			hdl.RegisterRoute(server)

			newReq, err := http.NewRequest(http.MethodPost, "/api/user/send-code", bytes.NewReader([]byte(tc.reqBody)))
			assert.NoError(t, err)
			newReq.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			server.ServeHTTP(resp, newReq)

			var res Response
			err = json.Unmarshal(resp.Body.Bytes(), &res)
			assert.Equal(t, tc.wantCode, resp.Code)
			assert.Equal(t, tc.wantBody, res)
		})
	}
}
