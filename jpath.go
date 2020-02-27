package jpath

import (
	"encoding/json"
	"github.com/rajnikant12345/jpath.git/vcopy"
	"fmt"
)

// JsonPath structure representing a JsonPath
type JsonPath struct {
	Path map[string]interface{}
}

/*
CompileNewJsonPath this function takes, multiple path strings and returns a JsonPath object.
*/
func CompileNewJsonPath(paths []string) (*JsonPath, error) {
	path := ""
	for _, v := range paths {
		path += `"` + v + `":"",`
	}
	path = "{" + path[:len(path)-1] + "}"

	m := map[string]interface{}{}

	err := json.Unmarshal([]byte(path), &m)

	if err != nil {
		return nil, err
	}

	j := JsonPath{map[string]interface{}{}}

	err = viper.MergeConfigMap(m)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&j.Path)
	if err != nil {
		return nil, err
	}

	fmt.Println(path)

	return &j, nil

}

func (e *JsonPath) getJsonAtPath(data interface{}, tokens map[string]interface{}) (out interface{}, err error) {

	for k, v := range tokens {
		if ismap(data) {
			in := data.(map[string]interface{})
			val, ok := in[k]
			if !ok {
				continue
			}
			if ismap(v) && isslice(val) {

				idx, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "JSON token must be an array ")
				}

				for k, v := range idx {
					index := isJsonPathSlice(k)

					if index == -1 {
						return nil, PathError(ConfigError, "Invalid index in configuration")
					} else {
						sl := val.([]interface{})
						if index == 0 {
							return sl, err
						} else {

							if index > len(sl) {
								return nil, PathError(ConfigError, "Invalid index in configuration")
							}
							if isfloat(sl[index-1]) {
								return sl[index-1], nil
							} else if isstring(sl[index-1]) {
								return sl[index-1], nil
							} else if ismap(sl[index-1]) {
								tok, ok := v.(map[string]interface{})
								if !ok {
									return nil, PathError(ConfigError, "Invalid configuration key "+k)
								}
								return e.getJsonAtPath(sl[index-1], tok)

							} else if isslice(sl[index-1]) {
								tok, ok := v.(map[string]interface{})
								if !ok {
									return nil, PathError(ConfigError, "Invalid configuration key "+k)
								}
								return e.getJsonAtPath(sl[index-1], tok)
							}
						}
					}
				}
			} else if ismap(v) && ismap(val) {
				tok, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "Invalid configuration key "+k)
				}
				return e.getJsonAtPath(val, tok)
			} else if v == "" && isstring(val) {
				return in[k], nil
			} else if v == "" && isfloat(val) {
				return in[k], nil
			} else if v == "" && ismap(val) {
				return in[k], nil
			}
		} else if isslice(data) {
			index := isJsonPathSlice(k)
			if index == -1 {
				return nil, PathError(ConfigError, "Invalid index in configuration")
			}
			in := data.([]interface{})
			if index == 0 {
				return in, nil
			} else {
				if index > len(in) {
					return nil, PathError(ConfigError, "Invalid index in configuration")
				}
				tok, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "Invalid configuration key "+k)
				}
				return e.getJsonAtPath(in[index-1], tok)
			}

		}
	}
	return nil, nil
}

/*
GetJsonAtPathValue returns a value in path , path is specified
while creating a JsonPath object.
*/
func (e *JsonPath) GetJsonAtPathValue(js interface{}) (interface{}, error) {
	return e.getJsonAtPath(js, e.Path)
}

/*
MapJsonAtPathValue calls function f on each value which occur in path , path is specified
while creating a JsonPath object. You can return a value in f anf this function will set value in
that path.
*/
func (e *JsonPath) MapJsonAtPathValue(js interface{}, f func(interface{}) interface{}) (interface{}, error) {
	return e.mapJsonAtPath(js, e.Path, f)
}

func (e *JsonPath) mapJsonAtPath(data interface{}, tokens map[string]interface{}, mapfunc func(interface{}) interface{}) (out interface{}, err error) {
	for k, v := range tokens {
		if ismap(data) {
			in := data.(map[string]interface{})
			val, ok := in[k]
			if !ok {
				continue
			}
			if ismap(v) && isslice(val) {

				idx, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "JSON token must be an array ")
				}

				for k, v := range idx {
					index := isJsonPathSlice(k)

					if index == -1 {
						break
						//return nil, PathError(ConfigError, "Invalid index in configuration")
					} else {
						sl := val.([]interface{})
						if index == 0 {
							for index := 1; index <= len(sl); index++ {
								if isfloat(sl[index-1]) {
									if mapfunc != nil {
										sl[index-1] = mapfunc(sl[index-1])
									}
								} else if isstring(sl[index-1]) {
									if mapfunc != nil {
										sl[index-1] = mapfunc(sl[index-1])
									}
								} else if ismap(sl[index-1]) {
									tok, ok := v.(map[string]interface{})
									if !ok {
										return nil, PathError(ConfigError, "Invalid configuration key "+k)
									}
									if mapfunc != nil {
										sl[index-1], err = e.mapJsonAtPath(sl[index-1], tok, mapfunc)
										if err != nil {
											return
										}
									}
								} else if isslice(sl[index-1]) {
									tok, ok := v.(map[string]interface{})
									if !ok {
										return nil, PathError(ConfigError, "Invalid configuration key "+k)
									}
									if mapfunc != nil {
										sl[index-1], err = e.mapJsonAtPath(sl[index-1], tok, mapfunc)
										if err != nil {
											return
										}
									}
								}
							}
						} else {
							if index > len(sl) {
								return nil, PathError(ConfigError, "Invalid index in configuration")
							}
							if isfloat(sl[index-1]) {
								if mapfunc != nil {
									sl[index-1] = mapfunc(sl[index-1])
								}
							} else if isstring(sl[index-1]) {
								if mapfunc != nil {
									sl[index-1] = mapfunc(sl[index-1])
								}
							} else if ismap(sl[index-1]) {
								tok, ok := v.(map[string]interface{})
								if !ok {
									return nil, PathError(ConfigError, "Invalid configuration key "+k)
								}
								if mapfunc != nil {
									sl[index-1], err = e.mapJsonAtPath(sl[index-1], tok, mapfunc)
									if err != nil {
										return
									}
								}
							} else if isslice(sl[index-1]) {
								tok, ok := v.(map[string]interface{})
								if !ok {
									return nil, PathError(ConfigError, "Invalid configuration key "+k)
								}
								if mapfunc != nil {
									sl[index-1], err = e.mapJsonAtPath(sl[index-1], tok, mapfunc)
									if err != nil {
										return
									}
								}
							}
						}

					}
				}
			} else if ismap(v) && ismap(val) {
				tok, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "Invalid configuration key "+k)
				}
				if mapfunc != nil {
					in[k], err = e.mapJsonAtPath(val, tok, mapfunc)
					if err != nil {
						return
					}
				}
			} else if v == "" && isstring(val) {
				if mapfunc != nil {
					in[k] = mapfunc(in[k])
				}
			} else if v == "" && isfloat(val) {
				if mapfunc != nil {
					in[k] = mapfunc(in[k])
				}
			}
		} else if isslice(data) {
			index := isJsonPathSlice(k)
			if index == -1 {
				continue
				//return nil, PathError(ConfigError, "Invalid index in configuration")
			}
			in := data.([]interface{})
			if index == 0 {
				for index := 0; index < len(in); index++ {
					tok, ok := v.(map[string]interface{})
					if !ok {
						return nil, PathError(ConfigError, "Invalid configuration key "+k)
					}
					if mapfunc != nil {
						in[index], err = e.mapJsonAtPath(in[index], tok, mapfunc)
						if err != nil {
							return nil, err
						}
					}
				}
			} else {
				if index > len(in) {
					return nil, PathError(ConfigError, "Invalid index in configuration")
				}
				tok, ok := v.(map[string]interface{})
				if !ok {
					return nil, PathError(ConfigError, "Invalid configuration key "+k)
				}
				if mapfunc != nil {
					in[index-1], err = e.mapJsonAtPath(in[index-1], tok, mapfunc)
					if err != nil {
						return nil, err
					}
				}
			}

		}
	}
	return data, nil
}
