// Application which greets you.
package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	url        = os.Getenv("URL")
	userID     = os.Getenv("USER_ID")
	titleMatch = os.Getenv("TITLE")
	textMatch  = os.Getenv("TEXT")
)

func run() error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf), chromedp.WithErrorf(log.Printf))
	defer cancel()
	var text string
	var title string
	log.Println("Starting")
	sel := `#TenantID`

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Title(&title),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if title != titleMatch {
				return errors.New("title did not match")
			}
			log.Println("Title matched")
			return nil
		}),
		chromedp.WaitReady(sel),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Done waiting for id")
			return nil
		}),
		chromedp.SendKeys(sel, userID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Done sending keys")
			return nil
		}),
		chromedp.Submit(sel),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Done Submitting")
			return nil
		}),
		chromedp.WaitVisible("[value=Yes]"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Done waiting for it to be visible")
			return nil
		}),
		chromedp.Text("form > p > strong", &text),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if strings.TrimSpace(text) != textMatch {
				return errors.New("text didn't match: " + text)
			}
			return nil
		}),
		chromedp.Submit("[value=Yes]"),
		chromedp.WaitReady("body"),
		chromedp.Text("p > strong", &text),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if strings.TrimSpace(text) != time.Now().Month().String() {
				return errors.New("month didn't match: " + text)
			}
			return nil
		}),
	)
	return err
}

func main() {
	if err := run(); err != nil {
		log.Println("err: ", err)
		os.Exit(1)
	}
	log.Println("Everything was successful. Exiting.")
}
