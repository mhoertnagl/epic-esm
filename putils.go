package main

func Forward(obj interface{}) (interface{}, error) {
	return obj, nil
}

func NewString(s []byte) (string, error) {
	return string(s), nil
}
