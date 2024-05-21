package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/handarudwiki/golang-ewalet/dto"
	"github.com/handarudwiki/golang-ewalet/internal/auth"
	"github.com/handarudwiki/golang-ewalet/internal/util"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository   domain.UserRepository
	chacheRepository domain.CacheRepository
	emailService     domain.EmailService
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository, emailService domain.EmailService) domain.UserService {
	return &userService{
		userRepository:   userRepository,
		chacheRepository: cacheRepository,
		emailService:     emailService,
	}
}

func (s userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		fmt.Println("iki error: ", err.Error())
		return dto.AuthRes{}, domain.ErrAuthFailed
	}
	jwtService := auth.NewJwtService()
	token, err := jwtService.GenerateToken(user)
	if err != nil {
		return dto.AuthRes{}, err
	}

	return dto.AuthRes{
		Token: token,
	}, nil
}

func (s userService) ValidateToken(ctx context.Context, tokenString string) (dto.UserData, error) {
	// Memuat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		return dto.UserData{}, err
	}

	// Mendapatkan secret key dari variabel lingkungan
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	// Memverifikasi token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return dto.UserData{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		userData := dto.UserData{
			ID:       int64(claims["user_id"].(float64)),
			Name:     claims["name"].(string),
			Username: claims["username"].(string),
			Phone:    claims["phone"].(string),
		}
		return userData, nil
	}

	return dto.UserData{}, fmt.Errorf("invalid token")
}

func (u userService) Register(ctx context.Context, req dto.UserRegisterReq) (dto.UserRegisterRes, error) {
	exist, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil && err != domain.ErrUserNotFound {
		// Jika terjadi error yang bukan "user tidak ditemukan", kembalikan error
		return dto.UserRegisterRes{}, err
	}

	if exist == (domain.User{}) {
		// Jika user ditemukan, kembalikan error bahwa username sudah diambil
		fmt.Println("error bagian sini")
		return dto.UserRegisterRes{}, domain.ErrUsernameTakeb
	}

	// Jika username tidak ditemukan, lanjutkan dengan pembuatan user baru
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	user := domain.User{
		Name:     req.Name,
		Password: string(hashedPassword),
		Phone:    req.Phone,
		Email:    req.Email,
		Username: req.Username,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	otpCode := util.GenerateRandomNumber(4)
	referenceID := util.GenerateRandomString(12)

	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	_ = u.chacheRepository.Set("otp:"+referenceID, []byte(otpCode))
	_ = u.chacheRepository.Set("user-ref:"+referenceID, []byte(req.Username))

	err = u.emailService.Send(req.Email, "Kirim OTP", "Ini OTP anda : "+otpCode)

	if err != nil {
		fmt.Println("error sending emailc")
		return dto.UserRegisterRes{}, err
	}

	return dto.UserRegisterRes{ReferenceID: referenceID}, nil
}

func (u userService) ValidateOTP(ctx context.Context, req dto.ValidateOtpReq) error {
	val, err := u.chacheRepository.Get("otp:" + req.ReferenceID)
	fmt.Println(string(val))
	if err != nil {
		return domain.ErrOTPInvalid
	}

	if req.OTP != string(val) {
		fmt.Println("surata")
		return domain.ErrOTPInvalid
	}

	val, err = u.chacheRepository.Get("user-ref:" + req.ReferenceID)

	if err != nil {
		return err
	}

	user, err := u.userRepository.FindByUsername(ctx, string(val))

	if err != nil {
		return err
	}

	verifiedAt := time.Now()

	err = u.userRepository.Update(ctx, &user, verifiedAt)

	if err != nil {
		fmt.Println("error update")
		return err
	}

	return nil
}
