package configs

type vscConfigChecker interface {
	check(url string, token string) error
}

type defaultVcsConfigChecker struct {
}

func (defaultVcsConfigChecker) check(url string, token string) error {
	return nil
}
