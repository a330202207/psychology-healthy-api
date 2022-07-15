// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/14 16:02
// @Package menu

package menu

import (
	"github.com/a330202207/psychology-healthy-api/internal/service"
)

type sMenu struct {
}

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

func (s *sMenu) Edit() {

}
