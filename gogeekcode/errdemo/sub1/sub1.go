package sub1

import (
	"errors"
	"github.com/Kate-liu/GoBeginner/gogeekcode/errdemo/sub1/sub2"
)

func Diff(foo int, bar int) error {
	if foo < 0 {
		return errors.New("diff error")
	}
	if err := sub2.Diff(foo, bar); err != nil {
		return err
	}
	return nil
}
