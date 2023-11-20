package xmlToJson

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"strings"
)

type XmlNode struct {
	Name       string
	ParentName string
	Attributes map[string]string
	Children   []*XmlNode
	Content    string
}

func ReadFile(filePath string) ([]byte, error) {
	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return xmlData, nil
}

func XmlToNode(xmlData []byte) (*XmlNode, error) {
	decoder := xml.NewDecoder(strings.NewReader(string(xmlData)))

	current := &XmlNode{
		Name:     "root",
		Children: make([]*XmlNode, 0),
	}

	stack := make([]*XmlNode, 0)

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			newNode := &XmlNode{
				Name: se.Name.Local,
			}

			if len(se.Attr) > 0 {
				newNode.Attributes = make(map[string]string)
				for _, attr := range se.Attr {
					newNode.Attributes[attr.Name.Local] = attr.Value
				}
			}

			newNode.ParentName = current.Name
			current.Children = append(current.Children, newNode)
			stack = append(stack, current)
			current = newNode
		case xml.EndElement:
			if len(stack) > 0 {
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			content := string([]byte(se))
			if strings.TrimSpace(content) != "" {
				current.Content = content
			}
		}
	}

	return current, nil
}

func (node *XmlNode) ToMap() interface{} {
	if len(node.Children) == 0 && node.Content == "" && len(node.Attributes) == 0 {
		return ""
	}

	nodeMap := make(map[string]interface{})

	// Add attributes, if any
	if len(node.Attributes) > 0 {
		for key, value := range node.Attributes {
			nodeMap["_"+key] = value
		}
	}

	// If there is content and no children, just return the content
	if node.Content != "" && len(node.Children) == 0 && len(node.Attributes) == 0 {
		return node.Content
	} else if node.Content != "" {
		nodeMap["__text"] = node.Content
	}

	// Recursively process children
	for _, child := range node.Children {
		nodeMap["_objectName"] = node.Name
		key, ok := nodeMap[child.Name]
		if ok {
			if _, ok := key.([]interface{}); ok {
				nodeMap[child.Name] = append(key.([]interface{}), child.ToMap())
			} else {
				nodeMap[child.Name] = []interface{}{key, child.ToMap()}
			}
		} else {
			nodeMap[child.Name] = child.ToMap()
		}
	}

	return nodeMap
}

func (node *XmlNode) ToJson() ([]byte, error) {
	if len(node.Children) == 0 && node.Content == "" && len(node.Attributes) == 0 {
		return json.Marshal("")
	}

	nodeMap := make(map[string]interface{})

	// Add attributes, if any
	if len(node.Attributes) > 0 {
		for key, value := range node.Attributes {
			nodeMap["_"+key] = value
		}
	}

	// If there is content and no children, just return the content
	if node.Content != "" && len(node.Children) == 0 && len(node.Attributes) == 0 {
		return json.Marshal(node.Content)
	} else if node.Content != "" {
		nodeMap["__text"] = node.Content
	}

	// Recursively process children
	for _, child := range node.Children {
		childJSON, err := child.ToJson()
		if err != nil {
			return nil, err
		}

		key, ok := nodeMap[child.Name]
		if ok {
			if _, ok := key.([]interface{}); ok {
				nodeMap[child.Name] = append(key.([]interface{}), json.RawMessage(childJSON))
			} else {
				nodeMap[child.Name] = []interface{}{key, json.RawMessage(childJSON)}
			}
		} else {
			nodeMap[child.Name] = json.RawMessage(childJSON)
		}
	}

	return json.Marshal(nodeMap)
}

func XmlToJson(xmlData []byte) ([]byte, error) {
	node, err := XmlToNode(xmlData)
	if err != nil {
		return nil, err
	}

	return node.ToJson()
}
