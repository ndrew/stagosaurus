package stagosaurus

// Deployer — is an interface for deploying posts
//
type Deployer interface {
	Deploy(Config, []Post) ([]Post, error)
}
