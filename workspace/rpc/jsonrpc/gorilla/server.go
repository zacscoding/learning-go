package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"log"
	"net/http"
	"regexp"
)

// ================= SMS service
type SmsService struct {
}

type SendSmsArgs struct {
	Number  string `json:"number"`
	Content string `Json:"content"`
}

func (s *SmsService) Send(r *http.Request, args *SendSmsArgs, result *string) error {
	fmt.Println("SmsService::Send() is called. number", args.Number, ", content", args.Content)
	err := s.checkNumber(args.Number)
	if err != nil {
		return err
	}

	*result = "Success to send sms"
	return nil
}

func (s *SmsService) IsSended(r *http.Request, number *string, result *string) error {
	fmt.Println("SmsService::IsSended() is called. number :", *number)
	err := s.checkNumber(*number)
	if err != nil {
		return err
	}

	*result = "true"
	return nil
}

func (s *SmsService) checkNumber(number string) error {
	re := regexp.MustCompile(`^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`)

	if !re.MatchString(number) {
		return errors.New("invalid phone number")
	}

	return nil
}

// ================= Email service
type EmailService struct {
}

type SendEmailArgs struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

func (e *EmailService) Send(r *http.Request, args *SendEmailArgs, result *string) error {
	fmt.Println("EmailService::Send() is called. number", args.Email, ", content", args.Content)
	*result = "Success to send email"
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	sms := new(SmsService)
	email := new(EmailService)

	_ = s.RegisterService(sms, "sms")
	_ = s.RegisterService(email, "email")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	log.Fatal(http.ListenAndServe(":3000", r))
}
