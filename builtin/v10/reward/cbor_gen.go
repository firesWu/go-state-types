// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package reward

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufState = []byte{139}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufState); err != nil {
		return err
	}

	// t.CumsumBaseline (big.Int) (struct)
	if err := t.CumsumBaseline.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.CumsumRealized (big.Int) (struct)
	if err := t.CumsumRealized.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.EffectiveNetworkTime (abi.ChainEpoch) (int64)
	if t.EffectiveNetworkTime >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.EffectiveNetworkTime)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.EffectiveNetworkTime-1)); err != nil {
			return err
		}
	}

	// t.EffectiveBaselinePower (big.Int) (struct)
	if err := t.EffectiveBaselinePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochReward (big.Int) (struct)
	if err := t.ThisEpochReward.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochRewardSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.ThisEpochRewardSmoothed.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochBaselinePower (big.Int) (struct)
	if err := t.ThisEpochBaselinePower.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if t.Epoch >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}

	// t.TotalStoragePowerReward (big.Int) (struct)
	if err := t.TotalStoragePowerReward.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.SimpleTotal (big.Int) (struct)
	if err := t.SimpleTotal.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.BaselineTotal (big.Int) (struct)
	if err := t.BaselineTotal.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) (err error) {
	*t = State{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 11 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.CumsumBaseline (big.Int) (struct)

	{

		if err := t.CumsumBaseline.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CumsumBaseline: %w", err)
		}

	}
	// t.CumsumRealized (big.Int) (struct)

	{

		if err := t.CumsumRealized.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.CumsumRealized: %w", err)
		}

	}
	// t.EffectiveNetworkTime (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EffectiveNetworkTime = abi.ChainEpoch(extraI)
	}
	// t.EffectiveBaselinePower (big.Int) (struct)

	{

		if err := t.EffectiveBaselinePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.EffectiveBaselinePower: %w", err)
		}

	}
	// t.ThisEpochReward (big.Int) (struct)

	{

		if err := t.ThisEpochReward.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochReward: %w", err)
		}

	}
	// t.ThisEpochRewardSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.ThisEpochRewardSmoothed.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochRewardSmoothed: %w", err)
		}

	}
	// t.ThisEpochBaselinePower (big.Int) (struct)

	{

		if err := t.ThisEpochBaselinePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochBaselinePower: %w", err)
		}

	}
	// t.Epoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Epoch = abi.ChainEpoch(extraI)
	}
	// t.TotalStoragePowerReward (big.Int) (struct)

	{

		if err := t.TotalStoragePowerReward.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalStoragePowerReward: %w", err)
		}

	}
	// t.SimpleTotal (big.Int) (struct)

	{

		if err := t.SimpleTotal.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.SimpleTotal: %w", err)
		}

	}
	// t.BaselineTotal (big.Int) (struct)

	{

		if err := t.BaselineTotal.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.BaselineTotal: %w", err)
		}

	}
	return nil
}

var lengthBufThisEpochRewardReturn = []byte{130}

func (t *ThisEpochRewardReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufThisEpochRewardReturn); err != nil {
		return err
	}

	// t.ThisEpochRewardSmoothed (smoothing.FilterEstimate) (struct)
	if err := t.ThisEpochRewardSmoothed.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ThisEpochBaselinePower (big.Int) (struct)
	if err := t.ThisEpochBaselinePower.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *ThisEpochRewardReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ThisEpochRewardReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ThisEpochRewardSmoothed (smoothing.FilterEstimate) (struct)

	{

		if err := t.ThisEpochRewardSmoothed.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochRewardSmoothed: %w", err)
		}

	}
	// t.ThisEpochBaselinePower (big.Int) (struct)

	{

		if err := t.ThisEpochBaselinePower.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ThisEpochBaselinePower: %w", err)
		}

	}
	return nil
}

var lengthBufAwardBlockRewardParams = []byte{132}

func (t *AwardBlockRewardParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufAwardBlockRewardParams); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Penalty (big.Int) (struct)
	if err := t.Penalty.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.GasReward (big.Int) (struct)
	if err := t.GasReward.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.WinCount (int64) (int64)
	if t.WinCount >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.WinCount)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.WinCount-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *AwardBlockRewardParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AwardBlockRewardParams{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Miner: %w", err)
		}

	}
	// t.Penalty (big.Int) (struct)

	{

		if err := t.Penalty.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Penalty: %w", err)
		}

	}
	// t.GasReward (big.Int) (struct)

	{

		if err := t.GasReward.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.GasReward: %w", err)
		}

	}
	// t.WinCount (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.WinCount = int64(extraI)
	}
	return nil
}
