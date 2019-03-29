// Copyright 2019 github.com/ucirello
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command poo creates a HTTP server than stream piles of poo (💩).
package main

import (
	"net/http"
	"time"
)

func poop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "💩", http.StatusInternalServerError)
		return
	}
	for {
		select {
		case <-r.Context().Done():
			return
		case <-t.C:
			if _, err := w.Write([]byte("💩")); err != nil {
				return
			}
			f.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/", poop)
	http.ListenAndServe(":8080", nil)
}
