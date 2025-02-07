package constants

import "github.com/crazyfrankie/gem/gerrors"

var (
	VerifyTooMany = gerrors.NewBizError(10001, "too many verifications")
	SendTooMany   = gerrors.NewBizError(10002, "send too many")
)
