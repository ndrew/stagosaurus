package stagosaurus

// 'Filestem' abstraction for retrieving/stroring Assets
//
type Storage interface {
	GetAsset(string) (Asset, error)
	StoreAsset(Asset) error
	//    Find(func(string, ))
	//func (this *Config) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {

}

/*
func (this *Config) FindByKey(predicate func(interface{}) bool) map[interface{}]interface{} {
    return this.Find(func(k interface{}, v interface{}) bool {
        return predicate(k)
    })
}

func (this *Config) FindByValue(predicate func(interface{}) bool) map[interface{}]interface{} {
    return this.Find(func(k interface{}, v interface{}) bool {
        return predicate(v)
    })
}

func (this *Config) Find(predicate func(interface{}, interface{}) bool) map[interface{}]interface{} {
    res := make(map[interface{}]interface{})
    for k, v := range this.cfg {
        if predicate(k, v) {
            res[k] = v
        }
    }
    return res
}

*/
