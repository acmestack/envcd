package util

import (
	"github.com/acmestack/envcd/internal/pkg/context"
	"github.com/acmestack/envcd/pkg/entity/data"
)

// BuildContext build plugin context
//  @param params params
//  @return *context.Context context
//  @return error error
func BuildContext(params interface{}) (*context.Context, error) {
	// TODO build context
	return nil, nil
}

// ParseContext parse context to envcd data
//  @param ctx context
//  @return *data.EnvcdData data
//  @return error error
func ParseContext(ctx *context.Context) (*data.EnvcdData, error) {
	// TODO parse context to envcd data
	return nil, nil
}
