# Developing
To re-generate the soap code, run `go generate`.

For unknown reasons, there are several types in the generated code which are redundant/recursive.  You will need to delete these.

Otherwise you may then `go build` as usual.
