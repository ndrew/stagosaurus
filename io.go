package stagosaurus

// 'Filestem' abstraction for retrieving/stroring Assets
//
type Storage interface {
	GetAsset(string) (Asset, error)
	StoreAsset(Asset) error
}
