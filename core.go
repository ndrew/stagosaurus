package stagosaurus

import (
	"errors"
	"reflect"
)

// Asset
// the smallest named/storable piece of information
type Asset interface {
	GetName() string
	GetContents() (*[]byte, error)
}

// the simpliest in-memory asset
//
type binaryAsset struct {
	name string
	data *[]byte
}

// Binary Asset Constructor
//
func BinaryAsset(name string, data *[]byte) Asset {
	asset := new(binaryAsset)
	asset.name = name
	asset.data = data
	return asset
}

func (this *binaryAsset) GetName() string {
	return this.name
}

func (this *binaryAsset) GetContents() (*[]byte, error) {
	return this.data, nil
}

// lazy loaded Asset
//
type lazyAsset struct {
	name     string
	loadFunc func(string) (*[]byte, error)
}

// Lazy loaded Asset contructor
//
func LazyAsset(name string, loadFunc func(string) (*[]byte, error)) Asset {
	asset := new(lazyAsset)
	asset.name = name
	asset.loadFunc = loadFunc
	return asset
}

func (this *lazyAsset) GetName() string {
	return this.name
}

func (this *lazyAsset) GetContents() (*[]byte, error) {
	return this.loadFunc(this.name)
}

// Post - an Asset container with meta-data
//
type Post interface {
	GetMeta() *Config
	GetAssets() []Asset
	// duplicate Asset interface,
	GetName() string
	GetContents() (*[]byte, error)
}

type postImpl struct {
	name   string
	data   *[]byte
	meta   *Config
	assets []Asset
}

// in-mem Post Constructor
//
func NewPost(name string, data interface{}, meta *Config, assets []Asset) (Post, error) {
	post := new(postImpl)
	post.name = name

	val := toValue(data)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		if val.Type() == typeOfBytes {
			binary := val.Bytes()
			post.data = &binary
		} else {
			return nil, errors.New("can't create post. Type " + val.Type().String() + " can't be converted to []byte")
		}
	case reflect.String:
		binary := []byte(val.String())
		post.data = &binary
	default:
		return nil, errors.New("can't create post. Type " + val.Type().String() + " can't be converted to []byte")
	}

	post.meta = meta
	post.assets = assets
	return post, nil
}

func (this *postImpl) GetName() string {
	return this.name
}

func (this *postImpl) GetContents() (*[]byte, error) {
	return this.data, nil
}

func (this *postImpl) GetMeta() *Config {
	return this.meta
}

func (this *postImpl) GetAssets() []Asset {
	return this.assets
}

// Posts' retrieving abstraction
//
type PostSource interface {
	GetPosts() ([]Post, error)
}
