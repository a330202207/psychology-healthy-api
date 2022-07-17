// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package menu

package menu

import (
	"context"
)

type sMenu struct {
}

func init() {
	// service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

func (s *sMenu) Edit(ctx context.Context) (err error) {

	return
}
