package client

import (
	"context"
	"testing"
)

const (
	mockClientId     = "mock_client_id"
	mockClientSecret = "mock_client_secret"
)

func TestV1beta_Check(t *testing.T) {
	type args struct {
		msgID string
	}
	tests := []struct {
		name    string
		cfg     Config
		args    args
		want    SmsCheckResp
		wantErr bool
	}{
		{
			name: "test Sms check on mock server",
			cfg: Config{
				ServerUrl:    Gateway,
				ServerMode:   ServerModeMock,
				ClientID:     mockClientId,
				ClientSecret: mockClientSecret,
			},
			args: args{
				msgID: "20200000-0000-0000-0000-380670000200",
			},
			want: SmsCheckResp{
				ReqId:  "eebacce4cc709aaa3f9c501c75492325",
				MsgId:  "20200000-0000-0000-0000-380670000200",
				Status: "delivered",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New(context.Background(), tt.cfg)
			got, err := client.Check(tt.args.msgID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Status != tt.want.Status {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
			if got.MsgId != tt.want.MsgId {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV1beta_Scoring(t *testing.T) {
	type args struct {
		phoneNumber string
		modelId     int
	}
	tests := []struct {
		name    string
		cfg     Config
		args    args
		want    ScoringResp
		wantErr bool
	}{
		{
			name: "test Scoring on mock server",
			cfg: Config{
				ServerUrl:    Gateway,
				ServerMode:   ServerModeMock,
				ClientID:     mockClientId,
				ClientSecret: mockClientSecret,
			},
			args: args{
				phoneNumber: "380670000200",
				modelId:     7,
			},
			want:    ScoringResp{Score: 500},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New(context.Background(), tt.cfg)
			got, err := client.Scoring(tt.args.phoneNumber, tt.args.modelId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scoring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Score != tt.want.Score {
				t.Errorf("Scoring() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV1beta_Send(t *testing.T) {
	type args struct {
		req SmsSendReq
	}
	tests := []struct {
		name    string
		cfg     Config
		args    args
		want    SmsSendResp
		wantErr bool
	}{
		{
			name: "test Send on mock server",
			cfg: Config{
				ServerUrl:    Gateway,
				ServerMode:   ServerModeMock,
				ClientID:     mockClientId,
				ClientSecret: mockClientSecret,
			},
			args: args{
				req: SmsSendReq{
					From: "messagedesk",
					To:   "380670000200",
					Text: "Hello World!",
				},
			},
			want: SmsSendResp{
				MsgId: "20200000-0000-0000-0000-380670000200",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := New(context.Background(), tt.cfg)
			got, err := client.Send(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.MsgId != tt.want.MsgId {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}
