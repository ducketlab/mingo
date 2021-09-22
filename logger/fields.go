package logger

type Field struct {
	Key   string
	Value interface{}
}

func NewAny(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

func NewFieldsFromKV(kvs map[string]interface{}) []Field {
	f := make([]Field, 0, len(kvs)/2+1)
	for k, v := range kvs {
		f = append(f, Field{Key: k, Value: v})
	}
	return f
}
