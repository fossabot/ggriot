package models

// MatchData is a struct that will hold the information about the match data api call.
type MatchData struct {
	SeasonID              int   `json:"seasonId"`
	QueueID               int   `json:"queueId"`
	GameID                int64 `json:"gameId"`
	ParticipantIdentities []struct {
		Player struct {
			CurrentPlatformID string `json:"currentPlatformId"`
			SummonerName      string `json:"summonerName"`
			MatchHistoryURI   string `json:"matchHistoryUri"`
			PlatformID        string `json:"platformId"`
			CurrentAccountID  int    `json:"currentAccountId"`
			ProfileIcon       int    `json:"profileIcon"`
			SummonerID        int    `json:"summonerId"`
			AccountID         int    `json:"accountId"`
		} `json:"player"`
		ParticipantID int `json:"participantId"`
	} `json:"participantIdentities"`
	GameVersion string `json:"gameVersion"`
	PlatformID  string `json:"platformId"`
	GameMode    string `json:"gameMode"`
	MapID       int    `json:"mapId"`
	GameType    string `json:"gameType"`
	Teams       []struct {
		FirstDragon          bool          `json:"firstDragon"`
		Bans                 []interface{} `json:"bans"`
		FirstInhibitor       bool          `json:"firstInhibitor"`
		Win                  string        `json:"win"`
		FirstRiftHerald      bool          `json:"firstRiftHerald"`
		FirstBaron           bool          `json:"firstBaron"`
		BaronKills           int           `json:"baronKills"`
		RiftHeraldKills      int           `json:"riftHeraldKills"`
		FirstBlood           bool          `json:"firstBlood"`
		TeamID               int           `json:"teamId"`
		FirstTower           bool          `json:"firstTower"`
		VilemawKills         int           `json:"vilemawKills"`
		InhibitorKills       int           `json:"inhibitorKills"`
		TowerKills           int           `json:"towerKills"`
		DominionVictoryScore int           `json:"dominionVictoryScore"`
		DragonKills          int           `json:"dragonKills"`
	} `json:"teams"`
	Participants []struct {
		Stats struct {
			FirstBloodAssist               bool `json:"firstBloodAssist"`
			VisionScore                    int  `json:"visionScore"`
			MagicDamageDealtToChampions    int  `json:"magicDamageDealtToChampions"`
			LargestMultiKill               int  `json:"largestMultiKill"`
			TotalTimeCrowdControlDealt     int  `json:"totalTimeCrowdControlDealt"`
			LongestTimeSpentLiving         int  `json:"longestTimeSpentLiving"`
			Perk1Var1                      int  `json:"perk1Var1"`
			Perk1Var3                      int  `json:"perk1Var3"`
			Perk1Var2                      int  `json:"perk1Var2"`
			TripleKills                    int  `json:"tripleKills"`
			Perk5                          int  `json:"perk5"`
			Perk4                          int  `json:"perk4"`
			PlayerScore9                   int  `json:"playerScore9"`
			PlayerScore8                   int  `json:"playerScore8"`
			Kills                          int  `json:"kills"`
			PlayerScore1                   int  `json:"playerScore1"`
			PlayerScore0                   int  `json:"playerScore0"`
			PlayerScore3                   int  `json:"playerScore3"`
			PlayerScore2                   int  `json:"playerScore2"`
			PlayerScore5                   int  `json:"playerScore5"`
			PlayerScore4                   int  `json:"playerScore4"`
			PlayerScore7                   int  `json:"playerScore7"`
			PlayerScore6                   int  `json:"playerScore6"`
			Perk5Var1                      int  `json:"perk5Var1"`
			Perk5Var3                      int  `json:"perk5Var3"`
			Perk5Var2                      int  `json:"perk5Var2"`
			TotalScoreRank                 int  `json:"totalScoreRank"`
			NeutralMinionsKilled           int  `json:"neutralMinionsKilled"`
			DamageDealtToTurrets           int  `json:"damageDealtToTurrets"`
			PhysicalDamageDealtToChampions int  `json:"physicalDamageDealtToChampions"`
			DamageDealtToObjectives        int  `json:"damageDealtToObjectives"`
			Perk2Var2                      int  `json:"perk2Var2"`
			Perk2Var3                      int  `json:"perk2Var3"`
			TotalUnitsHealed               int  `json:"totalUnitsHealed"`
			Perk2Var1                      int  `json:"perk2Var1"`
			Perk4Var1                      int  `json:"perk4Var1"`
			TotalDamageTaken               int  `json:"totalDamageTaken"`
			Perk4Var3                      int  `json:"perk4Var3"`
			LargestCriticalStrike          int  `json:"largestCriticalStrike"`
			LargestKillingSpree            int  `json:"largestKillingSpree"`
			QuadraKills                    int  `json:"quadraKills"`
			MagicDamageDealt               int  `json:"magicDamageDealt"`
			Item2                          int  `json:"item2"`
			Item3                          int  `json:"item3"`
			Item0                          int  `json:"item0"`
			Item1                          int  `json:"item1"`
			Item6                          int  `json:"item6"`
			Item4                          int  `json:"item4"`
			Item5                          int  `json:"item5"`
			Perk1                          int  `json:"perk1"`
			Perk0                          int  `json:"perk0"`
			Perk3                          int  `json:"perk3"`
			Perk2                          int  `json:"perk2"`
			Perk3Var3                      int  `json:"perk3Var3"`
			Perk3Var2                      int  `json:"perk3Var2"`
			Perk3Var1                      int  `json:"perk3Var1"`
			DamageSelfMitigated            int  `json:"damageSelfMitigated"`
			MagicalDamageTaken             int  `json:"magicalDamageTaken"`
			Perk0Var2                      int  `json:"perk0Var2"`
			TrueDamageTaken                int  `json:"trueDamageTaken"`
			Assists                        int  `json:"assists"`
			Perk4Var2                      int  `json:"perk4Var2"`
			GoldSpent                      int  `json:"goldSpent"`
			TrueDamageDealt                int  `json:"trueDamageDealt"`
			ParticipantID                  int  `json:"participantId"`
			PhysicalDamageDealt            int  `json:"physicalDamageDealt"`
			SightWardsBoughtInGame         int  `json:"sightWardsBoughtInGame"`
			TotalDamageDealtToChampions    int  `json:"totalDamageDealtToChampions"`
			PhysicalDamageTaken            int  `json:"physicalDamageTaken"`
			TotalPlayerScore               int  `json:"totalPlayerScore"`
			Win                            bool `json:"win"`
			ObjectivePlayerScore           int  `json:"objectivePlayerScore"`
			TotalDamageDealt               int  `json:"totalDamageDealt"`
			Deaths                         int  `json:"deaths"`
			PerkPrimaryStyle               int  `json:"perkPrimaryStyle"`
			PerkSubStyle                   int  `json:"perkSubStyle"`
			TurretKills                    int  `json:"turretKills"`
			FirstBloodKill                 bool `json:"firstBloodKill"`
			TrueDamageDealtToChampions     int  `json:"trueDamageDealtToChampions"`
			GoldEarned                     int  `json:"goldEarned"`
			KillingSprees                  int  `json:"killingSprees"`
			UnrealKills                    int  `json:"unrealKills"`
			FirstTowerAssist               bool `json:"firstTowerAssist"`
			FirstTowerKill                 bool `json:"firstTowerKill"`
			ChampLevel                     int  `json:"champLevel"`
			DoubleKills                    int  `json:"doubleKills"`
			InhibitorKills                 int  `json:"inhibitorKills"`
			Perk0Var1                      int  `json:"perk0Var1"`
			CombatPlayerScore              int  `json:"combatPlayerScore"`
			Perk0Var3                      int  `json:"perk0Var3"`
			VisionWardsBoughtInGame        int  `json:"visionWardsBoughtInGame"`
			PentaKills                     int  `json:"pentaKills"`
			TotalHeal                      int  `json:"totalHeal"`
			TotalMinionsKilled             int  `json:"totalMinionsKilled"`
			TimeCCingOthers                int  `json:"timeCCingOthers"`
		} `json:"stats"`
		Spell1ID                  int    `json:"spell1Id"`
		ParticipantID             int    `json:"participantId"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
		Spell2ID                  int    `json:"spell2Id"`
		TeamID                    int    `json:"teamId"`
		Timeline                  struct {
			Lane               string `json:"lane"`
			ParticipantID      int    `json:"participantId"`
			CsDiffPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"csDiffPerMinDeltas"`
			GoldPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"goldPerMinDeltas"`
			XpDiffPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"xpDiffPerMinDeltas"`
			CreepsPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"creepsPerMinDeltas"`
			XpPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"xpPerMinDeltas"`
			Role                        string `json:"role"`
			DamageTakenDiffPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"damageTakenDiffPerMinDeltas"`
			DamageTakenPerMinDeltas struct {
				Zero10 float64 `json:"0-10"`
			} `json:"damageTakenPerMinDeltas"`
		} `json:"timeline"`
		ChampionID int `json:"championId"`
	} `json:"participants"`
	GameDuration int   `json:"gameDuration"`
	GameCreation int64 `json:"gameCreation"`
}

// MatchHistory is the returned struct.
type MatchHistory struct {
	Matches []struct {
		Lane       string `json:"lane"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		PlatformID string `json:"platformId"`
		Timestamp  int64  `json:"timestamp"`
		Queue      int    `json:"queue"`
		Role       string `json:"role"`
		Season     int    `json:"season"`
	} `json:"matches"`
	EndIndex   int `json:"endIndex"`
	StartIndex int `json:"startIndex"`
	TotalGames int `json:"totalGames"`
}

// MatchTimeline is the struct returned when requesting a match's timeline.
type MatchTimeline struct {
	Frames []struct {
		Timestamp         int `json:"timestamp"`
		ParticipantFrames struct {
			Num1 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"1"`
			Num2 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"2"`
			Num3 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"3"`
			Num4 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"4"`
			Num5 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"5"`
			Num6 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"6"`
			Num7 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"7"`
			Num8 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"8"`
			Num9 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"9"`
			Num10 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"10"`
		} `json:"participantFrames"`
		Events []interface{} `json:"events"`
	} `json:"frames"`
	FrameInterval int `json:"frameInterval"`
}
