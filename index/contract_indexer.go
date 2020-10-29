package index

import (
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

var contractIndexPrefix = []byte("contractIndex")

type ContractIndex struct {
	db ethdb.Database
}

func NewContractIndex(db ethdb.Database) *ContractIndex {
	return &ContractIndex{
		db: db,
	}
}

type ContractParties struct {
	CreatorAddress      common.Address
	ParticipantAddreses []string
}

type contractPartiesRLP struct {
	CreatorAddress      common.Address
	ParticipantAddreses []string
}

func (ci ContractIndex) WriteIndex(contractAddress common.Address, contractParties *ContractParties) error {
	data, err := rlp.EncodeToBytes(contractParties)
	if err != nil {
		return err
	}
	if err = ci.db.Put(append(contractIndexPrefix, contractAddress.Bytes()...), data); err != nil {
		log.Error("Error writing contract index", "Contract Address", contractAddress)
		return err
	}
	return nil
}

func (ci ContractIndex) ReadIndex(contractAddress common.Address) (ContractParties, error) {
	var ca ContractParties
	contractPartiesBytes, err := ci.db.Get(append(contractIndexPrefix, contractAddress.Bytes()...))
	if err != nil {
		log.Error("Error retrieving Contract Addresses from index", "Contract Address", contractAddress)
		return ca, err
	}
	if len(contractPartiesBytes) == 0 {
		log.Error("Empty response returned", "Contract Address", contractAddress)
		return ca, errors.New("empty response querying contract index")
	}
	rlp.DecodeBytes(contractPartiesBytes, &ca)
	return ca, nil
}

func (cp *ContractParties) DecodeRLP(s *rlp.Stream) error {
	var partiesRLP contractPartiesRLP
	if err := s.Decode(&partiesRLP); err != nil {
		return err
	}
	cp.CreatorAddress, cp.ParticipantAddreses = partiesRLP.CreatorAddress, partiesRLP.ParticipantAddreses
	return nil
}

func (cp *ContractParties) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, contractPartiesRLP{
		CreatorAddress:      cp.CreatorAddress,
		ParticipantAddreses: cp.ParticipantAddreses,
	})
}