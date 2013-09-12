package stagosaurus

type Deployer interface {
	Deploy([]*Post) error
}
