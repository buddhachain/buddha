package define

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
)
