package configs

type TestConnectionParams struct {
	Url   string
	Token string
}

type ExistsParams struct {
	Url   string
	Token string
	Path  string
}

type FakeClient struct {
	NextErr         error
	LastParamTest   TestConnectionParams
	LastParamExists ExistsParams
}

func (f *FakeClient) TestConnection(url string, token string) error {
	err := f.NextErr
	f.LastParamTest = TestConnectionParams{
		Url:   url,
		Token: token,
	}
	f.NextErr = nil
	return err
}

func (f *FakeClient) Exists(url string, token string, path string) error {
	err := f.NextErr
	f.LastParamExists = ExistsParams{
		Url:   url,
		Token: token,
		Path:  path,
	}
	f.NextErr = nil
	return err
}

type FakeConfigService struct {
	NextErr          error
	NextReadVcs      VcsConfig
	LastSaveVcsParam VcsConfig
	NextReadDeps     []DepConfig
	LastAddDepParam  DepConfig
}

func (f *FakeConfigService) SaveVcs(config VcsConfig) error {
	err := f.NextErr
	f.LastSaveVcsParam = config
	f.NextErr = nil
	return err
}

func (f *FakeConfigService) ReadVcs() (VcsConfig, error) {
	read := f.NextReadVcs
	f.NextReadVcs = VcsConfig{}
	f.NextErr = nil
	return read, f.NextErr
}

func (f *FakeConfigService) AddDep(config DepConfig) error {
	err := f.NextErr
	f.LastAddDepParam = config
	f.NextErr = nil
	return err
}

func (f *FakeConfigService) ReadDeps() ([]DepConfig, error) {
	read := f.NextReadDeps
	f.NextReadDeps = nil
	f.NextErr = nil
	return read, f.NextErr

}
