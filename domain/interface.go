package domain

type IService interface {
	Registration() error
	Verification() error
	Login() error
}
