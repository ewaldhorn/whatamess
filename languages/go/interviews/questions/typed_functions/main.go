package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string
type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) string {
	newFilename := s.filenameTransformFunc(filename)
	return newFilename
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	return hex.EncodeToString(hash[:])
}

func prefixFilename(filename string) string {
	return fmt.Sprintf("prefix_%s", filename)
}

func genericPrefix(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {
	server := &Server{
		filenameTransformFunc: prefixFilename,
	}

	server.handleRequest("yolo.lol")
}
