// Copyright 2011 Andrew Scherkus
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/bobertlo/go-id3/id3"
)

func dumpFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open(%s): %s\n", path, err)
		return
	}
	defer f.Close()

	tags, err := id3.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "id3.Read(%s): %s\n", path, err)
		return
	}

	fmt.Println(path)
	fmt.Printf("Title\t%s\n", tags["title"])
	fmt.Printf("Artist\t%s\n", tags["artist"])
	fmt.Printf("Album\t%s\n", tags["album"])
	fmt.Printf("Year\t%s\n", tags["year"])
	fmt.Printf("Track\t%s\n", tags["track"])
	fmt.Printf("Disc\t%s\n", tags["disc"])
	fmt.Printf("Genre\t%s\n", tags["genre"])
	fmt.Printf("Length\t%s\n", tags["length"])
	fmt.Println()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [FILE]...\n", os.Args[0])
		return
	}

	for _, path := range os.Args[1:] {
		dumpFile(path)
	}
}
