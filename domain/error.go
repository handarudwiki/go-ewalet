package domain

import "errors"

var ErrAuthFailed = errors.New("error: Auth failed")
var ErrUsernameTakeb = errors.New("errror : username already taken")
var ErrOTPInvalid = errors.New("errror : otp is invalid")
var ErrUserNotFound = errors.New("errror : user not found")
var ErrAccountNotFound = errors.New("error : Account not found")
var ErrInsufficientBalance = errors.New("error : insufficient balance")
var ErrInquiryNotFound = errors.New("error : Inquiry not found")
var ErrPinInvalid = errors.New("error :invalid pin")
