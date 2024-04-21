package keeper

import (
  "context"
  "errors"
  "fmt"

  "cosmossdk.io/collections"
  "github.com/leemingi/checkers"
  "github.com/leemingi/checkers/rules"
)

// Keeper 타입의 객체를 내장하여 checkers.MsgServer 인터페이스를 구현
type msgServer struct {
  k Keeper
}

// 컴파일 시점에 인터페이스 구현을 강제하기 위한 선언
var _ checkers.MsgServer = msgServer{}

// MsgServer 인터페이스의 구현체를 반환
func NewMsgServerImpl(keeper Keeper) checkers.MsgServer {
  return &msgServer{k: keeper}
}

// MsgCreateGame 메시지에 대한 핸들러를 정의
func (ms msgServer) CreateGame(ctx context.Context, msg *checkers.MsgCreateGame) (*checkers.MsgCreateGameResponse, error) {
  // 게임 인덱스의 길이 검증
  if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
    return nil, checkers.ErrIndexTooLong
  }
  // 이미 제김이 존재하는지 확인
  if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
    return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
  }
  // 새로운 게임 보드 생성
  newBoard := rules.New()
  storedGame := checkers.StoredGame{
    Board: newBoard.String(),                 // 게임 보드 상태
    Turn:  rules.PieceStrings[newBoard.Turn], // 차례
    Black: msg.Black,                         // 흑색 플레이어의 주소
    Red:   msg.Red,                           // 적색 플레이어의 주소
  }
  // 게임 상태에 대한 유효성 검사
  if err := storedGame.Validate(); err != nil {
    return nil, err
  }
  // 게임 상태 저장
  if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
    return nil, err
  }
  // 이상이 없을 경우 게임 생성 응답 반환
  return &checkers.MsgCreateGameResponse{}, nil
}
