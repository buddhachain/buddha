package define

import "github.com/pkg/errors"

const (
	Success = iota //成功
	QueryErr
	ReadRequestBodyErr
	PostTxErr
	UnmarshalErr
	MarshalErr
	MarshalContractArgsErr
	PreExecErr
	PreInvokeWasmErr

	InsertDBErr
	QueryDBErr
	DeleteDBErr

	QueryContractErr

	LoadFileErr
	IpfsAddErr
	IpfsCatErr

	ReaderErr
	ConvertErr
	RightErr
	RequestErr
	ContractTxErr

	UnkownContractMethod
)

var (
	ErrRight        = errors.New("no permission")
	ErrRequest      = errors.New("request error")
	ErrContractTx   = errors.New("invalid contract tx")
	ErrContractArgs = errors.New("unmarshal contract args failed")
)
