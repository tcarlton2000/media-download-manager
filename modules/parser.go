package modules

import (
	"errors"
	"fmt"
	"media-download-manager/db"
	"media-download-manager/types"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type YoutubeDlParser struct {
	db       *db.Database
	download *types.Download
}

func (ydp YoutubeDlParser) Write(p []byte) (int, error) {
	s := string(p)

	err := ydp.parseTitle(s)
	if err != nil {
		return 0, err
	}

	err = ydp.parseStats(s)
	if err != nil {
		return 0, err
	}

	err = ydp.db.UpdateDownload(*ydp.download)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (ydp YoutubeDlParser) parseTitle(s string) error {
	nr, err := regexp.Compile("Destination: (.*)")
	if err != nil {
		return err
	}

	if nr.MatchString(s) {
		m := nr.FindStringSubmatch(s)
		name := m[1]
		ydp.download.Title = name
		ydp.download.Status = types.IN_PROGRESS
	}

	return nil
}

func (ydp YoutubeDlParser) parseStats(s string) error {
	pr, err := regexp.Compile(`((\d+\.)?\d+)%\s+of\s+((?:~\s*)?\d+\.\d+\w+)\s+at\s+(\d+\.\d+.+)\s+ETA\s+((\d+:)?\d+:\d+)`)
	if err != nil {
		return err
	}

	if pr.MatchString(s) {
		m := pr.FindStringSubmatch(s)
		tmp, err := strconv.ParseFloat(m[1], 32)
		if err != nil {
			return err
		}

		ydp.download.Progress = float32(tmp)
		ydp.download.TimeRemaining, err = parseTimeDuration(m[5])

		if err != nil {
			return err
		}
	}

	return nil
}

func parseTimeDuration(s string) (time.Duration, error) {
	ss := strings.Split(s, ":")
	if len(ss) != 2 {
		return 0, errors.New("invalid time format")
	}

	cs := fmt.Sprintf("%sm%ss", ss[0], ss[1])
	return time.ParseDuration(cs)
}
