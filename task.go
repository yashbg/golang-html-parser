package task

import (
	"bytes"
	"errors"
	"io/ioutil"

	"golang.org/x/net/html"
)

var ErrNotImplemented = errors.New("function not implemented")

// ReadHTMLFromFile should read the file from the current directory, if it exists.
// The file data should be returned as a string.
func ReadHTMLFromFile(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	return string(content), err
}

// CreateBuffer should transfer the contents of a string to a buffer.
func CreateBuffer(data string) bytes.Buffer {
	return *bytes.NewBufferString(data)
}

// CreateTree should create the tree representation of HTML represented by the buffer.
func CreateTree(buf bytes.Buffer) (*html.Node, error) {
	return html.Parse(bytes.NewReader(buf.Bytes()))
}

// CountDivTags should return the count of <div> tags in the document tree.
func CountDivTags(node *html.Node) int {
	tagsCount := make(map[string]int)
	dfs(node, tagsCount)
	return tagsCount["div"]
}

// dfs is a utility function which will help you count the number of unique tags.
func dfs(node *html.Node, tagsCount map[string]int) {
	tagsCount[node.DataAtom.String()]++
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, tagsCount)
	}
}

// ExtractAllUniqueTagsInSortedOrder should return the unique tags in the document.
// These tags should also be sorted alphabetically.
func ExtractAllUniqueTagsInSortedOrder(node *html.Node) []string {
	return nil
}

// ExtractAllComments returns the list of all comments as they appear in the document.
// You also need to remove all the leading and trailing spaces in the comments.
// HINT: You might need to read about variadic functions.
func ExtractAllComments(node *html.Node) []string {
	return nil
}

// ExtractAllLinks returns all the links in the document, in order of appearance.
func ExtractAllLinks(node *html.Node) []string {
	return nil
}
