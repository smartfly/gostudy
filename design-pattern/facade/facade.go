package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (imp *apiImpl) Test() string {
	aRet := imp.a.TestA()
	bRet := imp.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// AModuleAPI
type AModuleAPI interface {
	TestA() string
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// BModuleAPI
type BModuleAPI interface {
	TestB() string
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
