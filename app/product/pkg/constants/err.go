package constants

import "github.com/crazyfrankie/gem/gerrors"

var (
	Success        = gerrors.NewBizError(00000, "success")
	InternalServer = gerrors.NewBizError(00001, "internal server error")
)

var (
	ProductDraftNotFound = gerrors.NewBizError(20000, "product draft not found")
	ProductDraftExists   = gerrors.NewBizError(20001, "product exists")
)
