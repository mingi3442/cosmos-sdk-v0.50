package checkers

import (
    types "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces는 인터페이스 타입들을 레지스트리에 등록
func RegisterInterfaces(registry types.InterfaceRegistry) {
    // sd.Msg 인터페이스를 구현하는 모든 구현체를 등록하며 MsgCreateGame 메시지 타입을 등록
    registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgCreateGame{},
    )
    // msgservice 패키지를 사용하며 자동 생성된 gRPVC 서비스 설명자를 레지스트리에 등록
    msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
