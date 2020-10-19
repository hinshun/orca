package contentd

import (
	cid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
)

func Libp2pCidFromPubKey(data []byte) (cid.Cid, error) {
	hash, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return cid.Undef, err
	}

	return cid.NewCidV1(cid.Libp2pKey, hash), nil
}
