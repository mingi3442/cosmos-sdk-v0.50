package keeper

import (
	"fmt"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/leemingi/checkers"
)

type Keeper struct {
	// 바이너리 Codec Instance로 데이터 직렬화 및 역직렬화에 사용
	cdc codec.BinaryCodec
	// 주소 Codec Instance로 주소 형식의 직렬화 및 역직렬화에 사용
	addressCodec address.Codec
	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string // MsgUpdateParams와 다른 권한을 가진 메시지를 실행할 수 있는 주소로 보통 x/gov모듈 계정
	
	// state management
	Schema collections.Schema // 데이터 스키마를 정의하는 인스턴스
	Params collections.Item[checkers.Params] // 체커 게임의 파라미터를 저장하는 항목
	StoredGames collections.Map[string, checkers.StoredGame] // 진행 중인 체커 게임을 저장하는 데 사용하는 맵
}

  

// NewKeeper creates a new Keeper instance[]()
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	// authority 주소 유효성 검사
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}
	// 스키마 빌더 생성 => 데이터 저장 구조를 설정하는 데 사용
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc: cdc,
		addressCodec: addressCodec,
		authority: authority,
		// Params와 StoredGames는 각각 파라미터와 게임 데이터를 저장하는 컬렉션 아이템으로 초기화
		Params: collections.NewItem(sb, checkers.ParamsKey, "params", codec.CollValue[checkers.Params](cdc)),
		StoredGames: collections.NewMap(sb,
		checkers.StoredGamesKey, "storedGames", collections.StringKey,
		codec.CollValue[checkers.StoredGame](cdc)),
	}
	// 스키마 빌드 후 에러가 발생할 경우 패닉
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	// 생성된 스키마를 Keeper 인스턴스에 할당
	k.Schema = schema
	return k // Keeper Instance Return
}
	
	// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}