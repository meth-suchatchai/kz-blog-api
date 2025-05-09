package clientservices

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
	"github.com/meth-suchatchai/kz-blog-api/lib/kzstring"
	"github.com/meth-suchatchai/kz-blog-api/lib/taximail"
	"time"
)

func (svc *defaultService) Register(data clientmodels.RegisterData) (*clientmodels.RegisterOTPUser, *fiber.Error) {
	user, err := svc.userRepository.CreateUser(&dbmodels.User{
		MobileNumber:      fmt.Sprintf("%d", data.MobileNumber),
		FullName:          data.Name,
		PasswordEncrypted: svc.encryptedHash(data.Password),
	})
	if err != nil {
		return nil, err
	}

	//Send Mobile OTP
	mobileNumber, err := kzstring.ReplaceMobileCountryCode(data.CountryCode, data.MobileNumber)
	if err != nil {
		return nil, err
	}

	response, err := svc.mail.SendOTP(taximail.OTPRequest{
		MobileName: mobileNumber,
	})
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v", constant.NewUserKey, response.Data.OtpRefNo)
	otpUSer := &clientmodels.RegisterOTPUser{
		Id:                 user.Id,
		Name:               user.Name,
		MobileNumber:       mobileNumber,
		MessageId:          response.Data.MessageId,
		OTPReferenceNumber: response.Data.OtpRefNo,
	}
	val, vErr := json.Marshal(otpUSer)
	if vErr != nil {
		return nil, errors.NewDefaultFiberError(vErr)
	}
	vErr = svc.rdc.Set(context.TODO(), key, val, time.Minute*10).Err()
	if vErr != nil {
		return nil, errors.NewDefaultFiberError(vErr)
	}

	return otpUSer, nil
}
