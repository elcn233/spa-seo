package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func getBody(url string) (string, error) {
	// // create opts
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.ExecPath(execPath),
		chromedp.DisableGPU,
	)

	// create context
	allocatorContext, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	var body string
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		if err = chromedp.Navigate(url).Do(ctx); err != nil {
			return err
		}
		if err = chromedp.WaitVisible(`html`).Do(ctx); err != nil {
			return err
		}
		if err = chromedp.OuterHTML("html", &body).Do(ctx); err != nil {
			return err
		}
		return nil
	})); err != nil {
		return "", err
	}

	return body, nil
}
