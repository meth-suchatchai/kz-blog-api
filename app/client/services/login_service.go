package clientservices

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	clientmodels "github.com/kuroshibaz/app/client/models"
	usermodels "github.com/kuroshibaz/app/user/models"
	"github.com/kuroshibaz/lib/errors"
	kzjwt "github.com/kuroshibaz/lib/jwt"
)

func (svc *defaultService) Login(data clientmodels.LoginData) (*usermodels.User, *kzjwt.AccessToken, *fiber.Error) {
	mobileNumber := fmt.Sprintf("%v", data.MobileNumber)
	user, err := svc.userRepository.GetUserByMobileNumber(mobileNumber)
	log.Info("GetUserByMobileNumber: ", user, err)
	if err != nil {
		//check when mobile number not found
		return nil, nil, err
	}

	//Identify password user first
	if user.Password != svc.encryptedHash(data.Password) {
		return nil, nil, errors.NewDefaultFiberMessageError("password incorrect")
	}

	if !user.Active {
		return nil, nil, errors.NewDefaultFiberMessageError("account is not verified. please check your email")
	}

	if user.IsTFA {
		return nil, nil, fiber.NewError(errors.ErrCode2FA, "user has 2FA")
	}

	ac, err := svc.auth.JwtCreateToken(user)
	if err != nil {
		return nil, nil, err
	}

	return user, ac, nil
}
