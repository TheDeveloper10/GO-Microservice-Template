package repository

import "msvc-template/internal/util"

type Example struct {
}

func (repo *Example) Func() util.StatusCode {
	return util.StatusSuccess
}
