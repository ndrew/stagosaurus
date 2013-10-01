package stagosaurus

type Deployer interface {
	Deploy(Config, []Post) ([]Post, error)
}
