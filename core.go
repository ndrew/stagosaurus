package stagosaurus

import (
	"errors"
	"reflect"
)

// Asset describes an asset
// the smallest named/storable piece of information
type Asset interface {
	GetName() string
	GetContents() (*[]byte, error)
}

// Post - an Asset container with meta-data
//
type Post interface {
	GetConfig() Config
	GetAssets() []Asset
	// duplicate Asset interface,
	GetName() string
	GetContents() (*[]byte, error)
}

// ConfigSource — config retrieving abstraction
//
type ConfigSource interface {
	GetConfig() (Config, error)
}

// PostSource — Posts' retrieving abstraction
//
type PostSource interface {
	GetPosts(Config) ([]Post, error)
}

/////////
// impl

type postImpl struct {
	name   string
	data   *[]byte
	meta   Config
	assets []Asset
}

// the simpliest in-memory asset
//
type binaryAsset struct {
	name string
	data *[]byte
}

// BinaryAsset — binary asset constructor
//
func BinaryAsset(name string, data *[]byte) Asset {
	asset := new(binaryAsset)
	asset.name = name
	asset.data = data
	return asset
}

// GetName — returns name of an asset
//
func (asset *binaryAsset) GetName() string {
	return asset.name
}

// GetContents — returns contents of an asset
//
func (asset *binaryAsset) GetContents() (*[]byte, error) {
	return asset.data, nil
}

// lazy loaded Asset
//
type lazyAsset struct {
	name     string
	loadFunc func(string) (*[]byte, error)
}

// LazyAsset — Lazy loaded Asset contructor
//
func LazyAsset(name string, loadFunc func(string) (*[]byte, error)) Asset {
	asset := new(lazyAsset)
	asset.name = name
	asset.loadFunc = loadFunc
	return asset
}

// GetName — returns name of an asset
//
func (asset *lazyAsset) GetName() string {
	return asset.name
}

// GetContents — returns contents of an asset
//
func (asset *lazyAsset) GetContents() (*[]byte, error) {
	return asset.loadFunc(asset.name)
}

// NewPost — in-mem Post Constructor
//
func NewPost(name string, data interface{}, meta Config, assets []Asset) (Post, error) {
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

// GetName — returns name of the post
//
func (post *postImpl) GetName() string {
	return post.name
}

// GetContents — returns contents of the post
//
func (post *postImpl) GetContents() (*[]byte, error) {
	return post.data, nil
}

// GetConfig — returns post metadata
//
func (post *postImpl) GetConfig() Config {
	return post.meta
}

// GetAssets – returns list of post's assets
//
func (post *postImpl) GetAssets() []Asset {
	return post.assets
}
