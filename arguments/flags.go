package arguments

import "flag"

var (
	fakeNameFlag = flag.String(
		"fake-name",
		"",
		"The name of the fake struct",
	)

	outputPathFlag = flag.String(
		"o",
		"",
		"The file or directory to which the generated fake will be written",
	)
)
