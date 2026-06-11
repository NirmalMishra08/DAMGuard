package utils

import (
	"errors"
	"fmt"
	"net/http"
)

func ErrRequiredInputMissing(field string) error {
	return errors.New("required input missing: " + field)
}

func ErrUnchangeable(field string) error {
	return errors.New("field unchangeable after creation: " + field)
}

func ErrHeaderMissing(field string) error {
	return errors.New("header missing: " + field)
}

func ErrAccessoryMappedToUser(phone string) error {
	return fmt.Errorf("accessory is already mapped to user with phone: %s", phone)
}

var (
	ErrExpiredToken                          = errors.New("expired token")
	ErrInvalidToken                          = errors.New("invalid token")
	ErrInvalidOtp                            = errors.New("invalid otp")
	ErrInternal                              = errors.New("internal error")
	ErrTokenMissing                          = errors.New("token missing")
	ErrContextMissing                        = errors.New("context missing")
	ErrUnauthorized                          = errors.New("unauthorized")
	ErrInvalidRole                           = errors.New("invalid role")
	ErrEnvironment                           = errors.New("environment error")
	ErrDatabase                              = errors.New("database error")
	ErrOdoLimit                              = errors.New("odo add limit reached")
	ErrEmptyResult                           = errors.New("empty result")
	ErrTooManyRequests                       = errors.New("too many requests")
	ErrEntryExists                           = errors.New("entry already exists")
	ErrInvalidUser                           = errors.New("invalid user")
	ErrSessionExpired                        = errors.New("session expired (something fishy may be happening)")
	ErrSmsUnableToSend                       = errors.New("unable to send sms")
	ErrExpiredOtp                            = errors.New("expired otp")
	ErrVinAlreadyPaired                      = errors.New("vin already paired")
	ErrVinNotPaired                          = errors.New("vin not paired to your profile")
	ErrInvalidPin                            = errors.New("invalid pin")
	ErrFileNotFound                          = errors.New("file not found")
	ErrEmailAlreadyVerified                  = errors.New("email already verified")
	ErrEmailNotVerified                      = errors.New("email not verified")
	ErrUrlParamsMissing                      = errors.New("url params missing")
	ErrModelDoesNotExist                     = errors.New("model does not exist")
	ErrInvalidFileFormat                     = errors.New("invalid file format")
	ErrSamePin                               = errors.New("same pin")
	ErrTooManyEntries                        = errors.New("too many entries")
	ErrInvalidObjectId                       = errors.New("invalid object id")
	ErrLocationNotFound                      = errors.New("location not found")
	ErrCouldNotSaveNotification              = errors.New("could not save notification")
	ErrCouldNotReestablishDatabaseConnection = errors.New("could not reestablish database connection")
	ErrEventCreationLimitExceeded            = errors.New("event creation limit exceeded")
	ErrAlreadyRegistered                     = errors.New("already registered")
	ErrNotValidRequest                       = errors.New("not valid request")
	ErrMultipartFileMissing                  = errors.New("multipart file missing")
	ErrNotInLocation                         = errors.New("Not in location")
	ErrInvalidUuid                           = errors.New("invalid uuid")
	ErrGrpcConnFailed                        = errors.New("grpc connection failed")
	ErrGrpcQueryError                        = errors.New("grpc could not execute query")
	ErrNoGpsMapped                           = errors.New("no gps mapped to this profile")
	ErrGrpcCouldNotExecuteQuery              = errors.New("grpc could not execute query")
	ErrUserAlreadyExsists                    = errors.New("user already exsists")
	ErrUserDoesNotExist                      = errors.New("user does not exist")
	ErrWrongPassword                         = errors.New("wrong password")
	ErrNotEnoughCoins                        = errors.New("not enough coins for voucher")
	ErrInputInErrData                        = errors.New("Errdata should contain either 0 or 1")
	ErrFramenumber                           = errors.New("invalid frame number input")
	ErrFrameExists                           = errors.New("frame number already exists")
	ErrInvalidQueryParams                    = errors.New("invalid query parameters")
	ErrTokenGenError                         = errors.New("error in generating token")
	ErrForbidden                             = errors.New("You can only add upto 5 favourite locations")
	ErrAccessory                             = errors.New("Accessory type does not exist")
	ErrAccessoryMapped                       = errors.New("Accessory mapped with other user")
	ErrTripIdMissing                         = errors.New("No tripId passed")
	ErrMultipleLogin                         = errors.New("Multiple Login")
	ErrMinLatLng                             = errors.New("at least two positions are required to create a map")
	ErrMapUrl                                = errors.New("error fetching the image")
	ErrWebsocket                             = errors.New("error connecting to websocket")
	AccessoryAlreadyRequested                = errors.New("accessory already requested")
	ErrControllerIdAccessoryMapped           = errors.New("controllerId already exist with other framenumber")
	ErrBatteryIdAccessoryMapped              = errors.New("batteryId already exist with other framenumber")
	ErrMotorIdAccessoryMapped                = errors.New("motorId already exist with other framenumber")
)

var CustomErrorType = map[error]int{
	ErrExpiredToken:                          http.StatusUnauthorized,
	ErrInvalidToken:                          http.StatusUnauthorized,
	ErrInvalidOtp:                            419, // defining 419 for when a users token doesn't exist in cache
	ErrInternal:                              http.StatusInternalServerError,
	ErrTokenMissing:                          http.StatusUnauthorized,
	ErrContextMissing:                        http.StatusInternalServerError,
	ErrUnauthorized:                          http.StatusForbidden,
	ErrInvalidRole:                           http.StatusBadRequest,
	ErrEnvironment:                           http.StatusInternalServerError,
	ErrDatabase:                              http.StatusInternalServerError,
	ErrOdoLimit:                              http.StatusNotAcceptable,
	ErrEmptyResult:                           http.StatusBadRequest,
	ErrTooManyRequests:                       http.StatusTooManyRequests,
	ErrEntryExists:                           http.StatusBadRequest,
	ErrInvalidUser:                           http.StatusBadRequest,
	ErrSessionExpired:                        420, // defining 420 for when a users token doesn't exist in cache (someone else has created a new session)
	ErrSmsUnableToSend:                       430, // defining 430 for when an SMS is unable to send
	ErrExpiredOtp:                            429,
	ErrVinAlreadyPaired:                      427, // defining 427 for when a vin is already paired
	ErrVinNotPaired:                          http.StatusBadRequest,
	ErrFileNotFound:                          http.StatusInternalServerError,
	ErrEmailAlreadyVerified:                  http.StatusBadRequest,
	ErrInvalidPin:                            http.StatusBadRequest,
	ErrEmailNotVerified:                      http.StatusBadRequest,
	ErrModelDoesNotExist:                     http.StatusBadRequest,
	ErrInvalidFileFormat:                     http.StatusBadRequest,
	ErrUrlParamsMissing:                      http.StatusBadRequest,
	ErrSamePin:                               http.StatusBadRequest,
	ErrTooManyEntries:                        http.StatusBadRequest,
	ErrInvalidObjectId:                       http.StatusBadRequest,
	ErrLocationNotFound:                      http.StatusBadRequest,
	ErrCouldNotSaveNotification:              http.StatusInternalServerError,
	ErrCouldNotReestablishDatabaseConnection: http.StatusInternalServerError,
	ErrInputInErrData:                        http.StatusNotAcceptable,
	ErrFramenumber:                           422,
	ErrFrameExists:                           427,
	ErrInvalidQueryParams:                    http.StatusBadRequest,
	ErrTokenGenError:                         http.StatusInternalServerError,
	ErrForbidden:                             http.StatusForbidden,
	ErrAccessory:                             http.StatusBadRequest,
	ErrAccessoryMapped:                       409,
	ErrTripIdMissing:                         http.StatusBadRequest,
	ErrMultipleLogin:                         http.StatusConflict,
	ErrMinLatLng:                             http.StatusBadRequest,
	ErrMapUrl:                                403,
	ErrWebsocket:                             http.StatusInternalServerError,
	ErrUserDoesNotExist:                      http.StatusForbidden,
	AccessoryAlreadyRequested:                409,
	ErrControllerIdAccessoryMapped:           400,
	ErrBatteryIdAccessoryMapped:              400,
	ErrMotorIdAccessoryMapped:                400,
}
