package types

type ManagementEntry struct {
	ID     string
	Action string
}

type ManagementChan chan ManagementEntry
