package checkers

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
    return &GenesisState{
        Params: DefaultParams(),
    }
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
    // Params 객체 유효성 검증
    if err := gs.Params.Validate(); err != nil {
        return err
    }
    // unique 맵을 생성 및 중복된 인덱스 체크
    unique := make(map[string]bool)
    for _, indexedStoredGame := range gs.IndexedStoredGameList {
        if length := len([]byte(indexedStoredGame.Index)); MaxIndexLength < length || length < 1 {
            // 인덱스의 길이가 정의된 최대 길이를 초과하거나 0인 경우 에러 반환
            return ErrIndexTooLong
        }
        if _, ok := unique[indexedStoredGame.Index]; ok {
            // 이미 사용된 인덱스가 다시 사용되면 에러 반환
            return ErrDuplicateAddress
        }
        if err := indexedStoredGame.StoredGame.Validate(); err != nil {
            // 저장된 각 유효성을 검증 후 에러가 있을 경우 반환
            return err
        }
        // 유효한 인덱스를 unique 맵에 기록
        unique[indexedStoredGame.Index] = true
    }
    // 모든 검증이 완료될 경우 nil 반환
    return nil
}
