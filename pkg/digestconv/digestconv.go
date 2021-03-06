package digestconv

import (
	"encoding/hex"

	cid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
	digest "github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
)

func DigestToCid(dgst digest.Digest) (cid.Cid, error) {
	// in exceptionally rare cases we may have an empty digest
	// this check here prevents a panic from happening in those cases
	if dgst.String() == "" {
		return cid.Cid{}, errors.New("bad digest")
	}
	data, err := hex.DecodeString(dgst.Hex())
	if err != nil {
		return cid.Cid{}, errors.Wrap(err, "failed to decode digest hex")
	}

	encoded, err := multihash.Encode(data[:32], multihash.SHA2_256)
	if err != nil {
		return cid.Cid{}, errors.Wrap(err, "failed to encode digest as SHA256 multihash")
	}

	return cid.NewCidV1(cid.DagProtobuf, multihash.Multihash(encoded)), nil
}

func CidToDigest(c cid.Cid) (digest.Digest, error) {
	decoded, err := multihash.Decode(c.Hash())
	if err != nil {
		return "", err
	}

	return digest.NewDigestFromBytes(digest.Canonical, decoded.Digest), nil
}
