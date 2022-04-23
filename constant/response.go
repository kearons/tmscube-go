package constant

type ResponseCode int

const (
	SUCCESS  ResponseCode = 200
	NOTFOUND ResponseCode = 404
	ERROR    ResponseCode = 500
)

var CodeText = map[ResponseCode]string{
	SUCCESS:  "ok",
	NOTFOUND: "page not found",
	ERROR:    "server internal error",
}

func GetCodeText(code ResponseCode) string {
	if value, ok := CodeText[code]; ok {
		return value
	}
	return "unknown code"
}
