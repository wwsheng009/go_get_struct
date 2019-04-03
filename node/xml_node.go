package node

import (
	"encoding/xml"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
)

type NodeType int

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type NodeXml struct {
	Namespace string
	Type      NodeType
	TypeValue *TypeValue
	Attr      []xml.Attr
	Childern  []*NodeXml
}

type TypeValue struct {
	Value string
	Type  string
}

func Parse(r io.Reader) (*NodeXml, error) {
	decoder := xml.NewDecoder(r)
	ns := &nodeStack{}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			ns.nodeStart(t)
		case xml.EndElement:
			ns.nodeEnd()
		case xml.CharData:
			ns.addValue(t)
		}
	}
	if ns == nil || ns.nodes == nil && len(ns.nodes) == 0 {
		return nil, errors.New("Не удалось обработать документ. Возможно структура файла не является корректным Xml документом")
	}
	return ns.nodes[0], nil
}

type nodeStack struct {
	sync.Mutex
	nodes []*NodeXml
}

func (s *nodeStack) nodeStart(element xml.StartElement) {
	if s == nil {
		return
	}
	//Для корневого элемента
	if len(s.nodes) == 0 {
		n := newNode(element)
		s.nodes = append(s.nodes, n)
		return
	}
	//Для остальных элементов
	index := len(s.nodes) - 1
	n := newNode(element)
	s.nodes = append(s.nodes, n)
	s.nodes[index].Childern = append(s.nodes[index].Childern, n)
}

func (s *nodeStack) nodeEnd() {
	if s != nil && s.nodes != nil && len(s.nodes) > 1 {
		s.nodes = pop(s.nodes)
	}
}

func (s *nodeStack) addValue(data xml.CharData) {
	value := string(data)
	if s == nil || s.nodes == nil || len(s.nodes) <= 0 ||
		strings.TrimSpace(value) == "" {
		return
	}
	typ := defineType(value)
	index := len(s.nodes) - 1
	s.nodes[index].TypeValue.Value = value
	s.nodes[index].TypeValue.Type = typ
}

func newNode(element xml.StartElement) *NodeXml {
	return &NodeXml{
		Namespace: element.Name.Local,
		Type:      ElementNode,
		TypeValue: &TypeValue{"", ""},
		Attr:      element.Attr,
		Childern:  nil,
	}
}

func pop(input []*NodeXml) []*NodeXml {
	if input == nil || len(input) == 0 {
		return input
	}
	return input[:len(input)-1]
}

func defineType(charData string) string {
	i, err := strconv.ParseFloat(charData, 10)
	if err != nil {
		return "string"
	}
	if i == math.Trunc(i) {
		return "int64"
	}
	return "float64"
}
