package cmd

type LogCollectorFlags struct {
	failedOnly bool
	freeSpace  int
	contentIds []string
	hostfile   string
	hostnames  []string
	startDate  string
	endDate    string
	noPrompt   bool
	workingDir string
	segmentDir string
	osOnly     bool
	standby    bool
}
