package stf

import (
	"context"
)

func GetExecutionContext(ctx context.Context) *ExecutionContext {
	executionCtx, ok := ctx.(*ExecutionContext)
	if !ok {
		return nil
	}
	return executionCtx
}
