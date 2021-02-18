package transport

import (
	"github.com/fesyunoff/phone-book/pkg/service"
	. "github.com/swipe-io/swipe/v2"
)

func Swipe() {
	Build(
		Service(

			HTTPServer(),
			Interface((*service.Service)(nil), ""),

			JSONRPCEnable(),
			JSONRPCPath("/{method:.*}"),

			// OpenapiEnable(),
			// OpenapiOutput("./docs"),

			// MethodDefaultOptions(
			// Logging(true),
			// Instrumenting(true),
			// ),
		),
	)
}
