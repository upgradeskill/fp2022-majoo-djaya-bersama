package usecase

import (
	"errors"
	"log"
	"mini-pos/dto"
	"mini-pos/repository"
	"mini-pos/security"
	"mini-pos/util"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserUseCase interface {
	Login(ctx echo.Context) (dto.LoginResponse, []dto.ValidationMessage, error)
	Register(ctx echo.Context) (dto.User, []dto.ValidationMessage, error)
	UserInsert(ctx echo.Context, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error)
	UserUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error)
	UserDelete(id uint, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error)
}

type userUseCase struct {
	userRepository      repository.UserRepository
	authorizeRepository repository.AuthorizeRepository
	jwtService          security.JWTService
}

func InitUserUseCase(userRepository repository.UserRepository, authorizeRepository repository.AuthorizeRepository, jwtService security.JWTService) UserUseCase {
	return &userUseCase{
		userRepository:      userRepository,
		authorizeRepository: authorizeRepository,
		jwtService:          jwtService,
	}
}

func (uc *userUseCase) Login(ctx echo.Context) (dto.LoginResponse, []dto.ValidationMessage, error) {
	user := dto.User{}
	var invalidParameter []dto.ValidationMessage
	err := ctx.Bind(&user)
	response := dto.LoginResponse{}

	if err != nil {
		return dto.LoginResponse{}, nil, errors.New("failed to parse request")
	}
	//validation
	if user.Username == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "username", Message: "username is required"})
	}

	if user.Password == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "password", Message: "password is required"})
	}

	if len(invalidParameter) > 0 {
		return response, invalidParameter, nil
	}

	user, err = uc.userRepository.Login(user.Username, user.Password)
	if err != nil {
		return response, nil, err
	}

	if user.Username == "" {
		return response, nil, errors.New("username dan password tidak sesuai")
	}

	response.Username = user.Username
	response.Name = user.Name
	response.IsRole = user.IsRole
	response.IsActive = user.IsActive
	response.Token, err = uc.jwtService.GenerateToken(user.Id, user.Username, user.IsRole)

	if err != nil {
		return dto.LoginResponse{}, nil, errors.New("failed to generate token")
	}

	// save logged user into session
	session, _ := util.SessionStore.Get(ctx.Request(), util.SESSION_ID)
	session.Values["user_id"] = user.Id
	session.Values["is_role"] = user.IsRole
	session.Values["username"] = user.Username
	session.Save(ctx.Request(), ctx.Response())

	return response, nil, nil

}

func (uc *userUseCase) Register(ctx echo.Context) (dto.User, []dto.ValidationMessage, error) {

	payload := dto.User{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage
	if err != nil {
		return payload, nil, errors.New("failed to bind parametes")
	}

	//validation
	if payload.Username == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "username", Message: "username is required"})
	}

	if payload.Password == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "password", Message: "password is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.PhoneNumber == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "phone_number", Message: "password is required"})
	}

	payload.IsActive = dto.IsActive{IsActive: 1}
	payload.IsRole = 1

	if len(invalidParameter) > 0 {
		return dto.User{}, invalidParameter, nil
	}

	if payload.Password != "" {
		payload.Password, err = util.HashPassword(payload.Password)
		if err != nil {
			return dto.User{}, nil, err
		}
	}

	response, err := uc.userRepository.Insert(payload)
	response.Password = ""
	if err != nil {
		return dto.User{}, nil, err
	}

	return response, nil, nil
}

func (uc *userUseCase) UserInsert(ctx echo.Context, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error) {
	payload := dto.User{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}
	//validation
	if payload.Username == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "username", Message: "username is required"})
	}

	if payload.Password == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "password", Message: "password is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.PhoneNumber == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "phone_number", Message: "password is required"})
	}

	if payload.OutletId <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "outlet_id", Message: "Outlet id required"})
	}

	if claims.Role == 1 {
		err := uc.authorizeRepository.OwnerAuthorize(claims.Id, payload.OutletId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return payload, nil, errors.New("you don't have right")
			}
			return payload, nil, errors.New("failed to authorize")
		}

	} else {
		err := uc.authorizeRepository.StaffAuthorize(claims.Id, payload.OutletId)
		if err != nil {
			return payload, nil, errors.New("failed to authorize")
		}

	}

	if len(invalidParameter) > 0 {
		return dto.User{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}
	payload.CreatedBy = claims.Id

	if payload.Password != "" {
		payload.Password, err = util.HashPassword(payload.Password)
		if err != nil {
			return dto.User{}, nil, err
		}
	}

	response, err := uc.userRepository.Insert(payload)
	response.Password = ""
	if err != nil {
		return dto.User{}, nil, err
	}
	return response, nil, nil
}

func (uc *userUseCase) UserUpdate(ctx echo.Context, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error) {
	payload := dto.User{}
	err := ctx.Bind(&payload)
	var invalidParameter []dto.ValidationMessage

	if err != nil {
		log.Println(err)
		return payload, nil, errors.New("failed to bind parametes")
	}
	//validation
	if payload.Id <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	if payload.Username == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "username", Message: "username is required"})
	}

	if payload.Name == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "name", Message: "name is required"})
	}

	if payload.PhoneNumber == "" {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "phone_number", Message: "password is required"})
	}

	if payload.OutletId <= 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "outlet_id", Message: "Outlet id required"})
	}

	if claims.Role == 1 {
		err := uc.authorizeRepository.OwnerAuthorize(claims.Id, payload.OutletId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return payload, nil, errors.New("you don't have right")
			}
			return payload, nil, errors.New("failed to authorize")
		}

	} else {
		err := uc.authorizeRepository.StaffAuthorize(claims.Id, payload.OutletId)
		if err != nil {
			return payload, nil, errors.New("failed to authorize")
		}

	}

	if len(invalidParameter) > 0 {
		return dto.User{}, invalidParameter, nil
	}

	payload.IsActive = dto.IsActive{IsActive: 1}
	payload.CreatedBy = claims.Id

	if payload.Password != "" {
		payload.Password, err = util.HashPassword(payload.Password)
		if err != nil {
			return dto.User{}, nil, err
		}
	}

	response, err := uc.userRepository.Update(payload)
	response.Password = ""
	if err != nil {
		return dto.User{}, nil, err
	}
	return response, nil, nil
}
func (uc *userUseCase) UserDelete(id uint, claims *dto.UserClaims) (dto.User, []dto.ValidationMessage, error) {
	var invalidParameter []dto.ValidationMessage
	user := dto.User{}
	//validation
	if id == 0 {
		invalidParameter = append(invalidParameter, dto.ValidationMessage{Parameter: "id", Message: "id is required"})
	}

	// if claims.Role == 1 {
	// 	err := uc.authorizeRepository.OwnerAuthorize(claims.Id, id)
	// 	if err != nil {
	// 		if err == gorm.ErrRecordNotFound {
	// 			return user, nil, errors.New("you don't have right")
	// 		}
	// 		return user, nil, errors.New("failed to authorize")
	// 	}
	// } else {
	// 	err := uc.authorizeRepository.StaffAuthorize(claims.Id, id)
	// 	if err != nil {
	// 		if err == gorm.ErrRecordNotFound {
	// 			return user, nil, errors.New("you don't have right")
	// 		}
	// 		return user, nil, errors.New("failed to authorize")
	// 	}
	// }

	if len(invalidParameter) > 0 {
		return user, invalidParameter, nil
	}
	response, err := uc.userRepository.DeleteByID(id)
	if err != nil {
		return user, nil, err
	}
	return response, nil, nil
}
