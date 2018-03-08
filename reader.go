package vmf

import (
	"os"
	"strings"
	"bufio"
	"io"
)

const CHAR_ENTER_SCOPE = "{"
const CHAR_EXIT_SCOPE = "}"
const CHAR_ESCAPE = "\""
const CHAR_DISCARD_CUTSET = "{} \r\n"

type Reader struct {
	file *os.File
}

// Return a new Vmf Reader
func NewReader(file *os.File) Reader {
	reader := Reader{}
	reader.file = file
	return reader
}

// Read buffer file into our defined structures
// Returns a fully mapped Vmf structure
func (reader *Reader) Read() Vmf {
	bufReader := bufio.NewReader(reader.file)

	rootNode := Node{
		Key: reader.file.Name(),
	}

	readScope(bufReader, &rootNode)

	return Vmf{}
}

// A KeyValue object, that may hold multiple Values
type Node struct {
	Key string
	Values []interface{}
}
// A KeyValue object, can only hold a single Value
type Property struct {
	Key string
	Value interface{}
}

// Read a single scope
// Constructs a KeyValue node tree for a single scope
// Recursively parses all child scopes too
// Param: scope is the current scope to write to
func readScope(reader *bufio.Reader, scope *Node) *Node {
	for {
		line,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// New scope
		if strings.Contains(line, CHAR_ENTER_SCOPE) {
			// Scope is opened when the key is read
			continue
		} else
		// Exit scope
		if strings.Contains(line, CHAR_EXIT_SCOPE) {
			break
		} else
		// Is a property
		if strings.Contains(line, CHAR_ESCAPE) == true {
			p := strings.Split(line, CHAR_ESCAPE)
			if strings.Count(line, CHAR_ESCAPE) == 3 {
				// Multi-line property value, because value quotes aren't closed.
				// Read lines until we encounter a closing quote.
				for {
					line,_ = reader.ReadString('\n')
					p[3] += line

					if strings.Contains(line, CHAR_ESCAPE) == true {
						break
					}
				}
			}

			scope.Values = append(scope.Values,
				Property{
					Key: p[1],
					Value: p[3],
				})
			continue
		} else {
			// Name for new scope.
			newScope := Node{}
			newScope.Key = strings.Trim(line, CHAR_DISCARD_CUTSET)
			scope.Values = append(scope.Values, readScope(reader, &newScope))
		}
	}

	return scope
}