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
const NODE_TYPE_CAMERAS = "cameras"
const NODE_TYPE_CORDON = "cordon"
const NODE_TYPE_CORDONS = "cordons"
const NODE_TYPE_ENTITY = "entity"
const NODE_TYPE_VERSIONINFO = "versioninfo"
const NODE_TYPE_VIEWSETTINGS = "viewsettings"
const NODE_TYPE_VISGROUPS = "visgroups"
const NODE_TYPE_WORLD = "world"

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
func (reader *Reader) Read() (vmf Vmf, err error) {
	bufReader := bufio.NewReader(reader.file)

	rootNode := Node{
		key: reader.file.Name(),
	}

	readScope(bufReader, &rootNode)

	for _,n := range rootNode.value {
		node := n.(Node)
		switch node.key {
		case NODE_TYPE_CAMERAS:
			vmf.Cameras = node
			break
		case NODE_TYPE_CORDON:
			vmf.Cordon = node
			break
		case NODE_TYPE_CORDONS:
			vmf.Cordons = node
			break
		case NODE_TYPE_ENTITY:
			vmf.Entities.value = append(vmf.Entities.value, node)
			break
		case NODE_TYPE_VERSIONINFO:
			vmf.VersionInfo = node
			break
		case NODE_TYPE_VIEWSETTINGS:
			vmf.ViewSettings = node
			break
		case NODE_TYPE_WORLD:
			vmf.World = node
			break
		case NODE_TYPE_VISGROUPS:
			vmf.VisGroup.value = append(vmf.VisGroup.value, node)
			break
		default:
			break
		}
	}

	return vmf,err
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

					if strings.Contains(line, CHAR_ESCAPE) == true {
						p[3] += strings.Split(line, CHAR_ESCAPE)[0]
						break
					}
					p[3] += line
				}
			}

			node := Node{
				key: p[1],
			}
			node.value = append(node.value, p[3])
			scope.value = append(scope.value, node)
			continue
		} else {
			// Name for new scope.
			newScope := Node{}
			newScope.key = strings.TrimSpace(strings.Trim(line, CHAR_DISCARD_CUTSET))
			scope.value = append(scope.value, *readScope(reader, &newScope))
		}
	}

	return scope
}