package define

const (
	Success = iota //成功
	QueryErr
	ReadRequestBodyErr
	PostTxErr
	UnmarshalErr
	PreExecErr

	InsertDBErr
	QueryDBErr
)
