package define

import "github.com/pkg/errors"

const (
	Success = iota //成功
	QueryErr
	ReadRequestBodyErr
	PostTxErr
	UnmarshalErr
	PreExecErr
	PreInvokeWasmErr

	InsertDBErr
	QueryDBErr

	QueryContractErr

	LoadFileErr
	IpfsAddErr
	IpfsCatErr

	ReaderErr
	ConvertErr
	RightErr
	RequestErr
)

var (
	ErrRight   = errors.New("no permission")
	ErrRequest = errors.New("request error")
)
