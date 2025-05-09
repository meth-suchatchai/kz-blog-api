package clientservices

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	clientmodels "github.com/meth-suchatchai/kz-blog-api/app/client/models"
	constant "github.com/meth-suchatchai/kz-blog-api/const"
	"github.com/meth-suchatchai/kz-blog-api/lib/errors"
	"github.com/meth-suchatchai/kz-blog-api/lib/taximail"
)

func (svc *defaultService) VerifyOTP(data clientmodels.VerifyOTPData) *fiber.Error {
	key := fmt.Sprintf("%v:%v", constant.NewUserKey, data.OTPReferenceNumber)

	verifyByte, err := svc.rdc.Get(context.TODO(), key).Bytes()
	if err != nil {
		log.Error("redis get failed: ", err)
		return errors.NewDefaultFiberError(err)
	}

	var verifyData clientmodels.RegisterOTPUser
	err = json.Unmarshal(verifyByte, &verifyData)
	if err != nil {
		log.Error("unmarshall error: ", err)
		return errors.NewDefaultFiberError(err)
	}

	_, vErr := svc.mail.VerifyOTP(taximail.VerifyOTPRequest{
		MessageId:    verifyData.MessageId,
		OTPCode:      data.OTPCode,
		MobileNumber: verifyData.MobileNumber,
	})
	if vErr != nil {
		log.Error("verifyOTP error: ", vErr)
		return vErr
	}

	vErr = svc.userRepository.VerifyUser(verifyData.Id)
	if vErr != nil {
		log.Error("verifyUser error: ", vErr)
		return vErr
	}

	return nil
}
