package models

type Currencies struct {
	Name   string
	Amount int32
}

type Player struct {
	CreationTimestampMS uint64
	Username            string
	Team                int32
	Avatar              struct {
		Skin     int32
		Hair     int32
		Shirt    int32
		Pants    int32
		Hat      int32
		Shoes    int32
		Eyes     int32
		Backpack int32
	}
	MaxPokemonStorage int32
	MaxItemStorage    int32
	Currency          []Currencies
	ContactSettings   struct {
		SendMarketingEmails   bool
		SendPushNotifications bool
	}
	DailyBonus struct {
		NextCollectedTimestampMS            int64
		NextDefenderBonusCollectTimestampMS int64
	}
	EquippedBadge struct {
		BadgeType                         int32
		Level                             int32
		NextEquipChangeAllowedTimestampMS int64
	}
}
