package core_test

import (
	"fmt"
	"testing"

	"github.com/davidborzek/tvhgo/core"
	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
)

func TestNewDVRConfig(t *testing.T) {
	unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
		name:              "someName",
		priority:          1,
		cacheScheme:       1,
		duplicateHandling: 1,
		fileRetention:     30,
		infoRetention:     90,
	})
	mapped := core.NewDVRConfig(unmapped)

	assert.Equal(t, unmapped.UUID, mapped.ID)
	assert.Equal(t, unmapped.Enabled, mapped.Enabled)
	assert.Equal(t, unmapped.Name, mapped.Name)
	assert.False(t, mapped.Original)
	assert.Equal(t, unmapped.Profile, mapped.StreamProfileID)
	assert.Equal(t, core.DVRConfigPriorityHigh, mapped.Priority)
	assert.Equal(t, unmapped.RemoveAfterPlayback, mapped.DeleteAfterPlaybackTime)
	assert.Equal(t, core.DVRConfigRetentionTypeDays, mapped.RecordingInfoRetention.Type)
	assert.Equal(t, unmapped.RetentionDays, mapped.RecordingInfoRetention.Days)
	assert.Equal(t, core.DVRConfigRetentionTypeDays, mapped.RecordingFileRetention.Type)
	assert.Equal(t, unmapped.RemovalDays, mapped.RecordingFileRetention.Days)
	assert.Equal(t, unmapped.PreExtraTime, mapped.StartPadding)
	assert.Equal(t, unmapped.PostExtraTime, mapped.EndPadding)
	assert.Equal(t, unmapped.Clone, mapped.Clone)
	assert.Equal(t, unmapped.RerecordErrors, mapped.RerecordErrors)
	assert.Equal(t, unmapped.WarmTime, mapped.TunerWarmUpTime)
	assert.Equal(t, unmapped.Storage, mapped.Storage.Path)
	assert.Equal(t, unmapped.StorageMfree, mapped.Storage.MaintainFreeSpace)
	assert.Equal(t, unmapped.StorageMused, mapped.Storage.MaintainUsedSpace)
	assert.Equal(t, unmapped.DirectoryPermissions, mapped.Storage.DirectoryPermissions)
	assert.Equal(t, unmapped.FilePermissions, mapped.Storage.FilePermissions)
	assert.Equal(t, unmapped.Charset, mapped.Storage.Charset)
	assert.Equal(t, unmapped.Pathname, mapped.Storage.PathnameFormat)
	assert.Equal(t, core.DVRConfigCacheSchemeSystem, mapped.Storage.CacheScheme)
	assert.Equal(t, unmapped.DayDir, mapped.Subdirectories.DaySubdir)
	assert.Equal(t, unmapped.ChannelDir, mapped.Subdirectories.ChannelSubdir)
	assert.Equal(t, unmapped.TitleDir, mapped.Subdirectories.TitleSubdir)
	assert.Equal(t, unmapped.FormatTvmoviesSubdir, mapped.Subdirectories.TvMoviesSubdirFormat)
	assert.Equal(t, unmapped.FormatTvshowsSubdir, mapped.Subdirectories.TvShowsSubdirFormat)
	assert.Equal(t, unmapped.ChannelInTitle, mapped.File.IncludeChannel)
	assert.Equal(t, unmapped.DateInTitle, mapped.File.IncludeDate)
	assert.Equal(t, unmapped.TimeInTitle, mapped.File.IncludeTime)
	assert.Equal(t, unmapped.EpisodeInTitle, mapped.File.IncludeEpisode)
	assert.Equal(t, unmapped.SubtitleInTitle, mapped.File.IncludeSubtitle)
	assert.Equal(t, unmapped.OmitTitle, mapped.File.OmitTitle)
	assert.Equal(t, unmapped.CleanTitle, mapped.File.CleanTitle)
	assert.Equal(t, unmapped.WhitespaceInTitle, mapped.File.AllowWhitespace)
	assert.Equal(t, unmapped.WindowsCompatibleFilenames, mapped.File.WindowsCompatibleFilename)
	assert.Equal(t, unmapped.TagFiles, mapped.File.TagFiles)
	assert.Equal(t, core.DVRConfigDuplicateHandlingRecordAllDifferentEpisode, mapped.EPG.DuplicateHandling)
	assert.Equal(t, unmapped.EpgUpdateWindow, mapped.EPG.EpgUpdateWindow)
	assert.Equal(t, unmapped.EpgRunning, mapped.EPG.EpgRunning)
	assert.Equal(t, unmapped.SkipCommercials, mapped.EPG.SkipCommercials)
	assert.Equal(t, unmapped.AutorecMaxcount, mapped.EPG.Autorec.MaxCount)
	assert.Equal(t, unmapped.AutorecMaxsched, mapped.EPG.Autorec.MaxSchedules)
	assert.Equal(t, unmapped.FetchArtwork, mapped.Artwork.Fetch)
	assert.Equal(t, unmapped.FetchArtworkKnownBroadcastsAllowUnknown, mapped.Artwork.AllowUnidentifiableBroadcasts)
	assert.Equal(t, unmapped.FetchArtworkOptions, mapped.Artwork.CommandLineOptions)
	assert.Equal(t, unmapped.Preproc, mapped.Hooks.Start)
	assert.Equal(t, unmapped.Postproc, mapped.Hooks.Stop)
	assert.Equal(t, unmapped.Postremove, mapped.Hooks.Remove)
}

func TestNewDVRConfigOriginalTrueOnEmptyName(t *testing.T) {
	unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{})
	mapped := core.NewDVRConfig(unmapped)

	assert.Equal(t, unmapped.Name, mapped.Name)
	assert.True(t, mapped.Original)
}

func TestNewDVRConfigMapsPriorityCorrectly(t *testing.T) {
	var tests = []struct {
		input    int
		expected core.DVRConfigPriority
	}{
		{0, core.DVRConfigPriorityImportant},
		{1, core.DVRConfigPriorityHigh},
		{2, core.DVRConfigPriorityNormal},
		{3, core.DVRConfigPriorityLow},
		{4, core.DVRConfigPriorityUnimportant},
		{6, core.DVRConfigPriorityDefault},
		{10, core.DVRConfigPriorityUnknown},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("priority=%d, expected=%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
				priority: tt.input,
			})
			mapped := core.NewDVRConfig(unmapped)

			assert.Equal(t, tt.expected, mapped.Priority)
		})
	}
}

func TestNewDVRConfigMapsCacheSchemeCorrectly(t *testing.T) {
	var tests = []struct {
		input    int
		expected core.DVRConfigCacheScheme
	}{
		{0, core.DVRConfigCacheSchemeUnknown},
		{1, core.DVRConfigCacheSchemeSystem},
		{2, core.DVRConfigCacheSchemeDoNotKeep},
		{3, core.DVRConfigCacheSchemeSync},
		{4, core.DVRConfigCacheSchemeSyncAndDoNotKeep},
		{5, core.DVRConfigCacheSchemeUnknown},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("priority=%d, expected=%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
				cacheScheme: tt.input,
			})
			mapped := core.NewDVRConfig(unmapped)

			assert.Equal(t, tt.expected, mapped.Storage.CacheScheme)
		})
	}
}

func TestNewDVRConfigMapsDuplicateHandlingCorrectly(t *testing.T) {
	var tests = []struct {
		input    int
		expected core.DVRConfigDuplicateHandling
	}{
		{0, core.DVRConfigDuplicateHandlingRecordAll},
		{1, core.DVRConfigDuplicateHandlingRecordAllDifferentEpisode},
		{2, core.DVRConfigDuplicateHandlingRecordAllDifferentSubtitle},
		{3, core.DVRConfigDuplicateHandlingRecordAllDifferentDescription},
		{4, core.DVRConfigDuplicateHandlingRecordAllOncePerWeek},
		{5, core.DVRConfigDuplicateHandlingRecordAllOncePerDay},
		{6, core.DVRConfigDuplicateHandlingRecordLocalDifferentEpisode},
		{7, core.DVRConfigDuplicateHandlingRecordLocalDifferentTitle},
		{8, core.DVRConfigDuplicateHandlingRecordLocalDifferentSubtitle},
		{9, core.DVRConfigDuplicateHandlingRecordLocalDifferentDescription},
		{10, core.DVRConfigDuplicateHandlingRecordLocalOncePerWeek},
		{11, core.DVRConfigDuplicateHandlingRecordLocalOncePerDays},
		{12, core.DVRConfigDuplicateHandlingRecordAllOncePerMonth},
		{13, core.DVRConfigDuplicateHandlingRecordLocalOncePerMonth},
		{14, core.DVRConfigDuplicateHandlingRecordAllEpgUnique},
		{15, core.DVRConfigDuplicateHandlingRecordAll},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("priority=%d, expected=%s", tt.input, tt.expected)
		t.Run(testname, func(t *testing.T) {
			unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
				duplicateHandling: tt.input,
			})
			mapped := core.NewDVRConfig(unmapped)

			assert.Equal(t, tt.expected, mapped.EPG.DuplicateHandling)
		})
	}
}

func TestNewDVRConfigMapsInfoRetentionCorrectly(t *testing.T) {
	var tests = []struct {
		input    int64
		expected core.DVRConfigRetentionPolicy
	}{
		{20, core.DVRConfigRetentionPolicy{
			Days: 20,
			Type: core.DVRConfigRetentionTypeDays},
		},
		{core.TvheadendRetentionForever, core.DVRConfigRetentionPolicy{
			Days: 0,
			Type: core.DVRConfigRetentionTypeForever},
		},
		{core.TvheadendRetentionOther, core.DVRConfigRetentionPolicy{
			Days: 0,
			Type: core.DVRConfigRetentionTypeOnFileRemoval},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("priority=%d, expected=%s,%d", tt.input, tt.expected.Type, tt.expected.Days)
		t.Run(testname, func(t *testing.T) {
			unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
				infoRetention: tt.input,
			})
			mapped := core.NewDVRConfig(unmapped)

			assert.Equal(t, tt.expected, mapped.RecordingInfoRetention)
		})
	}
}

func TestNewDVRConfigMapsFileRetentionCorrectly(t *testing.T) {
	var tests = []struct {
		input    int64
		expected core.DVRConfigRetentionPolicy
	}{
		{20, core.DVRConfigRetentionPolicy{
			Days: 20,
			Type: core.DVRConfigRetentionTypeDays},
		},
		{core.TvheadendRetentionForever, core.DVRConfigRetentionPolicy{
			Days: 0,
			Type: core.DVRConfigRetentionTypeForever},
		},
		{core.TvheadendRetentionOther, core.DVRConfigRetentionPolicy{
			Days: 0,
			Type: core.DVRConfigRetentionTypeMaintainedSpace},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("priority=%d, expected=%s,%d", tt.input, tt.expected.Type, tt.expected.Days)
		t.Run(testname, func(t *testing.T) {
			unmapped := newTvhTestDVRConfig(tvhTestDvrConfig{
				fileRetention: tt.input,
			})
			mapped := core.NewDVRConfig(unmapped)

			assert.Equal(t, tt.expected, mapped.RecordingFileRetention)
		})
	}
}

type tvhTestDvrConfig struct {
	name              string
	priority          int
	cacheScheme       int
	duplicateHandling int
	fileRetention     int64
	infoRetention     int64
}

func newTvhTestDVRConfig(opts tvhTestDvrConfig) tvheadend.DVRConfig {
	return tvheadend.DVRConfig{
		UUID:                                    "someUUID",
		Enabled:                                 true,
		Name:                                    opts.name,
		Profile:                                 "someProfile",
		Pri:                                     opts.priority,
		RetentionDays:                           opts.infoRetention,
		RemovalDays:                             opts.fileRetention,
		RemoveAfterPlayback:                     120,
		PreExtraTime:                            5,
		PostExtraTime:                           10,
		Clone:                                   true,
		RerecordErrors:                          8,
		FetchArtwork:                            true,
		FetchArtworkKnownBroadcastsAllowUnknown: true,
		FetchArtworkOptions:                     "--foo=bar",
		Storage:                                 "/storage",
		StorageMfree:                            1000,
		StorageMused:                            600,
		DirectoryPermissions:                    "0755",
		FilePermissions:                         "0644",
		Charset:                                 "UTF-8",
		Pathname:                                "somePathnameFormat",
		Cache:                                   opts.cacheScheme,
		DayDir:                                  true,
		ChannelDir:                              true,
		TitleDir:                                true,
		FormatTvmoviesSubdir:                    "someTvmoviesSubdir",
		FormatTvshowsSubdir:                     "someTvshowsSubdir",
		ChannelInTitle:                          true,
		DateInTitle:                             true,
		TimeInTitle:                             true,
		EpisodeInTitle:                          true,
		SubtitleInTitle:                         true,
		OmitTitle:                               true,
		CleanTitle:                              true,
		WhitespaceInTitle:                       true,
		WindowsCompatibleFilenames:              true,
		TagFiles:                                true,
		EpgUpdateWindow:                         40,
		EpgRunning:                              true,
		AutorecMaxcount:                         45,
		AutorecMaxsched:                         35,
		Record:                                  opts.duplicateHandling,
		SkipCommercials:                         true,
		Preproc:                                 "some start hook",
		Postproc:                                "some stop hook",
		Postremove:                              "some remove hook",
		WarmTime:                                33,
	}
}
