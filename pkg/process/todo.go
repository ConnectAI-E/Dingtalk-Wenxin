package process

import (
	"github.com/ConnectAI-E/Dingtalk-Wenxin/public"
	"github.com/ConnectAI-E/go-wenxin/baidubce"
	ai_customv1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/ai_custom/v1"
	baidubcev1 "github.com/ConnectAI-E/go-wenxin/gen/go/baidubce/v1"

	"context"
)

func SingleQa(question string, user string) (string, error) {
	ctx := context.Background()
	var opts []baidubce.Option
	opts = append(opts, baidubce.WithTokenRequest(&baidubcev1.TokenRequest{
		GrantType:    "client_credentials",
		ClientId:     public.Config.BaiduClientID,
		ClientSecret: public.Config.BaiduClientSecret,
	}))
	client, _ := baidubce.New(opts...)

	req := &ai_customv1.ChatCompletionsRequest{
		User: user,
		Messages: []*ai_customv1.Message{
			{Role: "user", Content: question},
		},
	}
	res, _ := client.ChatCompletions(ctx, req)

	return res.Result, nil
}
