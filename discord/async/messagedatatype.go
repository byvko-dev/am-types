package async

type MessageDataType int

const (
	Empty MessageDataType = iota

	ReplayProccessRequest

	PlayerStatsRequest
	ClanStatsRequest

	ClanLeaderboardRequest
	PlayersLeaderboardRequest

	DebugRequest
)
