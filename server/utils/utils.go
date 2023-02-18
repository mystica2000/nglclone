package utils

import "reflect"

func IsEmpty(obj interface{}) bool {

	v := reflect.ValueOf(obj)

	switch v.Kind() {
		case reflect.Interface, reflect.Ptr:
			return v.IsNil()
		case reflect.String, reflect.Slice:
		  return v.Len() == 0
	}

	return false;
}


func IsProtectedName(route string) bool {
	set := make(map[string]bool)
	set["home"] = true
	set["name"] = true
	set["login"] = true
	set["error"] = true
	set["redirect"] = true

	if _, ok := set[route]; ok {
			return true;
	} else {
			return false;
	}
}