package checkers

import (
  fmt "fmt"

  "cosmossdk.io/errors"
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/leemingi/checkers/rules"
)

// StoredGame 구조체에서 흑색 플레이어의 주소를 AccAddress 형식으로 변환하여 반환
func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
  // Bech32 형식의 문자열 주소를 AccAddress로 변환
  black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
  // 변환 중 에러가 발생할 경우 해당 에러를 ErrInvalidBlock에 래핑하여 반환
  return black, errors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

// StoredGame 구조체에서 적생 플레이어의 주소를 AccAddress 형식으로 변환하여 반환
func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
  // Bech32 형식의 문자열 주소를 AccAddress로 변환
  red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
  // 변환 중 에러가 발생할 경우 해당 에러를 ErrInvalidBlock에 래핑하여 반환
  return red, errors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

// 저장된 게임의 게임 보드를 파싱하여 rules.Game 포인터 반환
func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
  //보드의 상태를 문자열에서 Game 구조체로 파싱
  board, errBoard := rules.Parse(storedGame.Board)
  if errBoard != nil {
    // 파싱 에러가 발생시 에러 반환
    return nil, errors.Wrapf(errBoard, ErrGameNotParseable.Error())
  }
  // 게임의 현재 턴을 설정
  board.Turn = rules.StringPieces[storedGame.Turn].Player
  if board.Turn.Color == "" {
    return nil, errors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
  }
  return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
  // 흑색 플레이어 주소 유효성 검증
  _, err = storedGame.GetBlackAddress()
  if err != nil {
    return err
  }
  // 적색 플레이어의 주소 유효성 검증
  _, err = storedGame.GetRedAddress()
  if err != nil {
    return err
  }
  // 게임 보드의 유효성 검증
  _, err = storedGame.ParseGame()
  return err
}
