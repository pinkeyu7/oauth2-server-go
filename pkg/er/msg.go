package er

var msg = map[int]string{
	ErrorParamInvalid:          "Wrong parameter format or invalid",
	UnauthorizedError:          "Unauthorized",
	ForbiddenError:             "Forbidden error",
	ResourceNotFoundError:      "Resource not found",
	TokenExpiredError:          "Token is expired",
	AWSInitError:               "aws sdk init error",
	FirebaseIdTokenError:       "Firebase IdToken verify error",
	DataDuplicateError:         "Data duplicate error",
	LimitExceededError:         "Limit exceeded error",
	DecryptError:               "Decrypt error",
	OauthClientDataError:       "Oauth client data invalid error",
	OauthValidateError:         "Oauth validate error",
	UnknownError:               "Database unknown error",
	DBInsertError:              "Database insertion error",
	DBUpdateError:              "Database update error",
	DBDeleteError:              "Database delete error",
	DBDuplicateKeyError:        "Database data duplicate key error",
	RedisSerError:              "Redis server error",
	UploadFileErrUnknown:       "Unknown error",
	UploadFileErrNotExist:      "File not exist",
	UploadFileErrSizeOverLimit: "File size over limit",
	UploadFileErrFileName:      "File name over limit",
	UploadFileErrEmpty:         "File is empty",
	UploadFileErrTypeNotMatch:  "File type not match",
	UploadFileErrRowOverLimit:  "File row over limit",
	OauthUnknownError:          "Oauth unknown error",
}
