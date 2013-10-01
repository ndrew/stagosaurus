package stagosaurus

type Deployer interface {
	Deploy([]Post) ([]Post, error)
}
