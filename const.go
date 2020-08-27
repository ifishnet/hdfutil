// Copyright (c) 2013-2014 The ifishnet developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package hdfutil

const (
	// SatoshiPerBitcent is the number of satoshi in one bitcoin cent.
	SatoshiPerBitcent = 1e6

	// SatoshiPerBitcoin is the number of satoshi in one bitcoin (1 HDF).
	SatoshiPerBitcoin = 1e8

	// MaxSatoshi is the maximum transaction amount allowed in satoshi.
	MaxSatoshi = 21e6 * SatoshiPerBitcoin
)
