package employee

import "testing"

type fakeStmtforMaleCount struct {
	Stmt
}

func (fakeStmtforMaleCount) Exec(stmt string,args ...string) (Result, error) {
	return Result{Count: 5}, nil
}

func TestEmployeeMaleCount(t *testing.T) {
	f := fakeStmtforMaleCount{}
	c, _ := MaleCount(f)
	if c != 5 {
		t.Errorf("want: %d, actual: %d", 5, c)
		return
	}
}