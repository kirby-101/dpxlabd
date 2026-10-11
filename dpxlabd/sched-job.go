package dpxlabd

type SchedJobContext struct{}

type SchedJob interface {
	Run(ctx *SchedJobContext) error
	Abort() error
}
