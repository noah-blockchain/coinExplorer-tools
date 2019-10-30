package models

type ExtenderEnvironment struct { // todo wsLink
	AppName                         string
	BaseCoin                        string
	CoinsUpdateTime                 int
	Debug                           bool
	DbName                          string
	DbUser                          string
	DbPassword                      string
	DbHost                          string
	DbPort                          int
	DbMinIdleConns                  int
	DbPoolSize                      int
	WsHost                          string
	WsPort                          int
	WsKey                           string
	NodeApi                         string
	ApiHost                         string
	ApiPort                         int
	TxChunkSize                     int
	AddrChunkSize                   int
	EventsChunkSize                 int
	StakeChunkSize                  int
	WrkSaveRewardsCount             int
	WrkSaveSlashesCount             int
	WrkSaveTxsCount                 int
	WrkSaveTxsOutputCount           int
	WrkSaveInvTxsCount              int
	WrkSaveAddressesCount           int
	WrkSaveValidatorTxsCount        int
	WrkUpdateBalanceCount           int
	WrkGetBalancesFromNodeCount     int
	WrkUpdateTxsIndexNumBlocks      int
	WrkUpdateTxsIndexTime           int
	RewardAggregateEveryBlocksCount int
	RewardAggregateTimeInterval     string
}
