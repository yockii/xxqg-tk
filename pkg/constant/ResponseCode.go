package constant

const (
	ResponseCodeErrorInternal = -iota - 10000
	ResponseCodeErrorParamParse
	ResponseCodeErrorParamRequired
	ResponseCodeErrorDataExists
	ResponseCodeErrorUnsupported
	ResponseCodeErrorResourceNotFound
	ResponseCodeErrorDataNotMatch
)
