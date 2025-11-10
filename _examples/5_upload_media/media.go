package main

import (
	"context"

	"github.com/V0lkanTas/gotwi"
	"github.com/V0lkanTas/gotwi/media/upload"
	"github.com/V0lkanTas/gotwi/media/upload/types"
	"github.com/V0lkanTas/gotwi/tweet/managetweet"
	mtTypes "github.com/V0lkanTas/gotwi/tweet/managetweet/types"
)

func Initialize(c *gotwi.Client, p *types.InitializeInput) (*types.InitializeOutput, error) {
	res, err := upload.Initialize(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Append(c *gotwi.Client, p *types.AppendInput) (*types.AppendOutput, error) {
	res, err := upload.Append(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Finalize(c *gotwi.Client, p *types.FinalizeInput) (*types.FinalizeOutput, error) {
	res, err := upload.Finalize(context.Background(), c, p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func PostWithMedia(c *gotwi.Client, text string, mediaID string) (string, error) {
	p := &mtTypes.CreateInput{
		Text: gotwi.String(text),
		Media: &mtTypes.CreateInputMedia{
			MediaIDs: []string{mediaID},
		},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}
