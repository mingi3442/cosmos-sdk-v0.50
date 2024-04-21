package keeper

import (
  "context"
  "errors"

  "cosmossdk.io/collections"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "github.com/leemingi/checkers"
)

// 컴파일 시점에 인터페이스 구현을 확인하기 위한 선언
var _ checkers.QueryServer = queryServer{}

// QueryServer 인터페이스의 구현체 반환
func NewQueryServerImpl(k Keeper) checkers.QueryServer {
  return queryServer{k}
}

// queryServer 구조체는 Keeper를 포함하여 게임의 상태를 관리
type queryServer struct {
  k Keeper
}

// Query/GetGame RPC 메서드에 대한 핸들러 정의
func (qs queryServer) GetGame(ctx context.Context, req *checkers.QueryGetGameRequest) (*checkers.QueryGetGameResponse, error) {
  // 요청된 인덱스에 해당하는 게임을 조회
  game, err := qs.k.StoredGames.Get(ctx, req.Index)
  if err == nil {
    // 게임이 성공적으로 조회되면 게임 정보를 포함하여 응답을 반환
    return &checkers.QueryGetGameResponse{Game: &game}, nil
  }
  // 게임이 없을 경우 에러 없이 nil 응답을 반환
  if errors.Is(err, collections.ErrNotFound) {
    return &checkers.QueryGetGameResponse{Game: nil}, nil
  }
  // 그 외의 에러는 내부 서버 에러로 처리하여 GRPC 상태와 함께 반환
  return nil, status.Error(codes.Internal, err.Error())
}
