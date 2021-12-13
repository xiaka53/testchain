package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xiaka53/testchain/x/testchain/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count
	byteKey := []byte(types.PostKey)
	//使用k.storeKey(testchai)和PostCountKey获取stores
	stores := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, post.Id)
	appendValue := k.cdc.MustMarshal(&post)
	stores.Set(bz, appendValue)
	k.SetPostCount(ctx, count+1)
	return count
}

//Get PostCountKey
func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.PostCountKey)
	//使用k.storeKey(testchai)和PostCountKey获取stores
	stores := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := stores.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

//Set PostCountKey
func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.PostCountKey)
	//使用k.storeKey(testchai)和PostCountKey获取stores
	stores := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	stores.Set(byteKey, bz)
}
