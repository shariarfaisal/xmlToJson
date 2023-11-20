# Package Documentation: xmlToJson

The `xmlToJson` package provides functionality for converting XML data to a structured JSON representation. It includes functions for reading XML data from a file, converting XML to a hierarchical structure, and generating JSON output.

## Types

### XmlNode

```go
type XmlNode struct {
    Name       string
    ParentName string
    Attributes map[string]string
    Children   []*XmlNode
    Content    string
}
```

- `Name`: The name of the XML node.
- `ParentName`: The name of the parent XML node.
- `Attributes`: A map of attribute names to attribute values.
- `Children`: A slice of child XML nodes.
- `Content`: The content of the XML node.

## Functions

### ReadFile

```go
func ReadFile(filePath string) ([]byte, error)
```

- Reads XML data from the specified file path.
- Returns the XML data as a byte slice.

### XmlToNode

```go
func XmlToNode(xmlData []byte) (*XmlNode, error)
```

- Converts XML data to a hierarchical `XmlNode` structure.
- Returns the root `XmlNode` of the hierarchy.

### (node \*XmlNode) ToMap

```go
func (node *XmlNode) ToMap() interface{}
```

- Converts an `XmlNode` to a map representation.
- Returns a map representing the structure of the XML node.

### (node \*XmlNode) ToJson

```go
func (node *XmlNode) ToJson() ([]byte, error)
```

- Converts an `XmlNode` to a JSON representation.
- Returns a JSON-encoded byte slice representing the structure of the XML node.

### XmlToJson

```go
func XmlToJson(xmlData []byte) ([]byte, error)
```

- Converts XML data to a JSON representation.
- Returns a JSON-encoded byte slice.

## Usage Example

```go
package main

import (
	"fmt"
	"github.com/your-username/xmlToJson"
)

func main() {
	xmlData, err := xmlToJson.ReadFile("path/to/your/file.xml")
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}

	node, err := xmlToJson.XmlToNode(xmlData)
	if err != nil {
		fmt.Println("Error converting XML to node:", err)
		return
	}

	jsonData, err := node.ToJson()
	if err != nil {
		fmt.Println("Error converting node to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
```

## Notes

- This package assumes well-formed XML data.
- The XML to JSON conversion preserves the hierarchical structure of the XML data.
- The package uses the `encoding/xml` and `encoding/json` standard packages for XML and JSON processing, respectively.

Feel free to adapt and use the package in your projects!
