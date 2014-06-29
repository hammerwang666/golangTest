package config

func GetConfig(k string) interface{} {
	configLock.RLock()
	defer configLock.RUnlock()
	v := configMap[k]
	return v
}
