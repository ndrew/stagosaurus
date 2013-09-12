package blog

type Deployer interface {
	Deploy([]*Post) error
}
