package module

import (
  autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
  checkersv1 "github.com/leemingi/checkers/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
  return &autocliv1.ModuleOptions{
    Query: &autocliv1.ServiceCommandDescriptor{
      // checkers 게임 쿼리 서비스의 이름 설정
      Service: checkersv1.Query_ServiceDesc.ServiceName,
      // RPC 명령 옵션을 정의
      RpcCommandOptions: []*autocliv1.RpcCommandOptions{
        {
          RpcMethod: "GetGame",                                    // GetGame 메소드를 CLI에 등록
          Use:       "get-game index",                             // 명령 사용 방법
          Short:     "Get the current value of the game at index", // 짧은 설명
          PositionalArgs: []*autocliv1.PositionalArgDescriptor{
            {ProtoField: "index"},
          }, // 위치 기반 인자를 설명하는 배열
        },
      },
    },
    Tx: &autocliv1.ServiceCommandDescriptor{ // 트랜잭션 관련 CLI 명령을 설정
      Service: checkersv1.Msg_ServiceDesc.ServiceName, // 서비스 이름을 체커스 게임 서비스로 지정
      RpcCommandOptions: []*autocliv1.RpcCommandOptions{
        {
          RpcMethod: "CreateGame",                                                             // RPC 메서드 이름
          Use:       "create index black red",                                                 // 사용 방법 설명
          Short:     "Creates a new checkers game at the index for the black and red players", // 쩗은 설명
          PositionalArgs: []*autocliv1.PositionalArgDescriptor{ // 인자 설명
            {ProtoField: "index"}, // 인덱스 필드
            {ProtoField: "black"}, // 흑색 플레이어 필드
            {ProtoField: "red"},   // 적색 플레이어 필드ㅊㅊ
          },
        },
      },
    },
  }
}
