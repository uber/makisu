//  Copyright (c) 2018 Uber Technologies, Inc.
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

package pathutils

var (
	// DefaultBlacklist is the list of all paths that should be ignored when
	// copying directories around and computing filesystem diffs.
	DefaultBlacklist = append([]string{
		DefaultInternalDir,
		// Found through experiments
		"/dev",
		"/.dockerinit",
		"/srv",
		"/mnt",
	}, dockerBlacklist...)

	dockerBlacklist = []string{
		// Docker made these locations read-only.
		// See https://github.com/moby/moby/blob/a5f9783c930834b8e6035fb0ad9c22fd4bbfc355/daemon/initlayer/setup_unix.go
		"/.dockerenv",
		"/dev/console",
		"/dev/pts",
		"/dev/shm",
		"/etc/hosts",
		"/etc/hostname",
		"/etc/mtab",
		"/etc/resolv.conf",
		"/proc",
		"/sys",
	}
)

// DefaultStorageDir is the default directory makisu uses for persisted
// data like cached image layers.
// This directory should be mounted for better performance, if makisu
// runs in a container.
const DefaultStorageDir = "/makisu-storage"

// DefaultInternalDir is used for Makisu binary and certs.
const DefaultInternalDir = "/makisu-internal"

// DefaultBinaryPath is Makisu binary path.
// It should be excluded from all file operations.
const DefaultBinaryPath = "/makisu-internal/makisu"

// DefaultCACertsPath containsa list of common CA certs.
// These certs are generated by ca-certficates debian package and appended to system certs.
const DefaultCACertsPath = "/makisu-internal/certs/cacerts.pem"

// CacheKeyValueFileName is the name of local cache key value file.
const CacheKeyValueFileName = "cache_key_value.json"
