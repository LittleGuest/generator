// golang和mysql类型转换

package generator

var conv map[string]interface{}

func init() {
	if conv == nil {
		conv = make(map[string]interface{}, 0)
	}
	conv["integer"] = [...]string{"int", "int32", "uint", "uint32", "rune"}
	conv["bigint"] = [...]string{"int64", "uint64"}
	conv["bool"] = "bool"
	conv["varchar"] = "string"
	conv["longtext"] = "string"
	conv["date"] = "time.Time"
	conv["datetime"] = "time.Time"
	conv["tinyint"] = [...]string{"byte", "int8", "uint8"}
	conv["smallint"] = [...]string{"int16", "uint16"}
	conv["double"] = [...]string{"float32", "float64"}
	conv["numeric"] = "float64"
}
