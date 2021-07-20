// Copyright 2019 Andy Pan & Dietoad. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// reference from https://github.com/panjf2000/ants/blob/master/internal/spinlock.go

package mutex

import (
    "runtime"
    "sync"
    "sync/atomic"
)

type spinLock uint32

const maxBackoff = 64

func (sl *spinLock) Lock() {
    backoff := 1
    for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
        // Leverage the exponential backoff algorithm, see https://en.wikipedia.org/wiki/Exponential_backoff.
        for i := 0; i < backoff; i++ {
            runtime.Gosched()
        }
        if backoff < maxBackoff {
            backoff <<= 1
        }
    }
}

func (sl *spinLock) Unlock() {
    atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() sync.Locker {
    return new(spinLock)
}
