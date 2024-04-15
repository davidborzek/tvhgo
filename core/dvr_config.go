package core

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrDVRConfigNotFound = errors.New("dvr config not found")
)

type DVRConfigPriority string

const (
	DVRConfigPriorityImportant   DVRConfigPriority = "important"
	DVRConfigPriorityHigh        DVRConfigPriority = "high"
	DVRConfigPriorityNormal      DVRConfigPriority = "normal"
	DVRConfigPriorityLow         DVRConfigPriority = "low"
	DVRConfigPriorityUnimportant DVRConfigPriority = "unimportant"
	DVRConfigPriorityDefault     DVRConfigPriority = "default"
	DVRConfigPriorityUnknown     DVRConfigPriority = "unknown"
)

type (
	// DVRConfig defines the configuration for the DVR.
	DVRConfig struct {
		// ID is the unique identifier for the DVR configuration.
		ID string `json:"id"`
		// Enabled indicates if the DVR is enabled.
		Enabled bool `json:"enabled"`
		// Name is the name of the DVR configuration.
		Name string `json:"name"`
		// Original indicates if the DVR configuration is the original.
		Original bool `json:"original"`
		// StreamProfileID is the unique identifier of the stream profile to use.
		StreamProfileID string `json:"streamProfileId"`
		// Priority is the priority of the DVR configuration.
		Priority DVRConfigPriority `json:"priority"`
		// DeleteAfterPlaybackTime defines the amount of time in seconds to keep the recording after playback.
		// If set to 0, the recording will be kept indefinitely.
		DeleteAfterPlaybackTime int64 `json:"deleteAfterPlayback"`
		// RetentionDays is the amount of days to keep the recording info.
		RetentionDays int64 `json:"retentionDays"`
		// RemovalDays is the amount of days to keep the recording file.
		RemovalDays int64 `json:"removalDays"`
		// StartPadding optional padding in minutes to record
		// before the recording starts.
		StartPadding int `json:"startPadding"`
		// EndPadding optional padding in minutes to record
		// after the recording ends.
		EndPadding int `json:"endPadding"`
		// Clone indicates whether the scheduled entry should be cloned and re-recorded
		// if the recording fails.
		Clone bool `json:"clone"`
		// RerecordErrors defines the amount of errors that can occur before the recording is re-scheduled.
		RerecordErrors int `json:"rerecordErrors"`
		// TunerWarmUpTime is the time in seconds to wait for the tuner to get ready for recording.
		TunerWarmUpTime int `json:"tunerWarmUpTime"`
		// Storage is the storage settings for the DVR.
		Storage DVRConfigStorageSettings `json:"storage"`
		// Subdirectories is the subdirectory settings for the DVR.
		Subdirectories DVRConfigSubdirectorySettings `json:"subdirectories"`
		// File is the file settings for the DVR.
		File DVRConfigFileSettings `json:"file"`
		// EPG is the EPG settings for the DVR.
		EPG DVRConfigEPGSettings `json:"epg"`
		// Artwork is the artwork settings for the DVR.
		Artwork DVRConfigArtworkSettings `json:"artwork"`
		// Hooks is the hooks settings for the DVR.
		Hooks DVRConfigHooks `json:"hooks"`
	}

	// DVRConfigStorageSettings defines the storage settings for the DVR.
	DVRConfigStorageSettings struct {
		// Path is the path where the recording should be stored.
		Path string `json:"path"`
		// MaintainFreeSpace is the amount of free space to maintain on the storage path.
		MaintainFreeSpace int `json:"maintainFreeSpace"`
		// MaintainUsedSpace is the amount of used space to maintain on the storage path.
		MaintainUsedSpace int `json:"maintainUsedSpace"`
		// DirectoryPermissions is the permissions to create new directories with.
		DirectoryPermissions string `json:"directoryPermissions"`
		// FilePermissions is the permissions to create new files with.
		FilePermissions string `json:"filePermissions"`
		// Charset is the character set to use for the filenames.
		Charset string `json:"charset"`
		// PathnameFormat is the format of the pathname.
		// See Tvheadend Help for more information.
		PathnameFormat string `json:"pathnameFormat"`
		// CacheScheme is the cache scheme to use/used to store recordings.
		CacheScheme int `json:"cacheScheme"`
	}

	// DVRConfigSubdirectorySettings defines the subdirectory settings for the DVR.
	DVRConfigSubdirectorySettings struct {
		// DaySubdir indicates if a new directory should be created for each day.
		// It will only be created when something is recorded. The format is ISO Standard (YYYY-MM-DD).
		DaySubdir bool `json:"daySubdir"`
		// ChannelSubdir indicates if a new directory should be created for each channel.
		// If this and `daySubdir` are both enabled, the channel directory will be created inside the day directory.
		ChannelSubdir bool `json:"channelSubdir"`
		// TitleSubdir indicates if a new directory should be created for each title.
		// If this, `daySubdir` and `channelSubdir` are all enabled, the title directory will be created inside the channel directory.
		TitleSubdir bool `json:"titleSubdir"`
		// TvMoviesSubdirFormat is the subdirectory format for tvmovies when using the $q specifier.
		// This can contain only alphanumeric characters (A-Za-z0-9).
		TvMoviesSubdirFormat string `json:"tvMoviesSubdirFormat"`
		// TvhshowsSubdirFormat is the subdirectory format for tvshows when using the $q specifier.
		// This can contain only alphanumeric characters (A-Za-z0-9).
		TvShowsSubdirFormat string `json:"tvShowsSubdirFormat"`
	}

	// DVRConfigFileSettings defines the file settings for the DVR.
	DVRConfigFileSettings struct {
		// IncludeChannel indicates if the channel should be included in the title.
		// This applies fot the filename and the title in tag stored in the file.
		IncludeChannel bool `json:"channelInTitle"`
		// IncludeDate indicates if the date should be included in the title.
		// This applies fot the filename and the title in tag stored in the file.
		IncludeDate bool `json:"dateInTitle"`
		// IncludeTime indicates if the time should be included in the title.
		// This applies fot the filename and the title in tag stored in the file.
		IncludeTime bool `json:"timeInTitle"`
		// IncludeEpisode indicates if the episode should be included in the title when available.
		IncludeEpisode bool `json:"episodeInTitle"`
		// IncludeSubtitle indicates if the subtitle should be included in the title when available.
		IncludeSubtitle bool `json:"subtitleInTitle"`
		// OmitTitle indicates if the title should be omitted.
		OmitTitle bool `json:"omitTitle"`
		// CleanTitle indicates if the title should be cleaned.
		CleanTitle bool `json:"cleanTitle"`
		// AllowWhitespace indicates if whitespace should be included in the title.
		AllowWhitespace bool `json:"whitespaceInTitle"`
		// WindowsCompatibleFilename indicates if the filename should be Windows compatible.
		WindowsCompatibleFilename bool `json:"windowsCompatibleFilename"`
		// TagFiles indicates if the files should be tagged with metadata.
		TagFiles bool `json:"tagFiles"`
	}

	// DVRConfigEPGSettings defines the EPG settings for the DVR.
	DVRConfigEPGSettings struct {
		// DuplicateHandling defines the duplicate recording handling.
		DuplicateHandling int `json:"duplicateHandling"`
		// EpgUpdateWindow is the maximum allowed difference between the event start time
		// when the epg event is changed.
		EpgUpdateWindow int64 `json:"epgUpdateWindow"`
		// EpgRunning indicates if EITp/f should be used to decide event start and stop.
		// Also known as "Accurate Recording".
		EpgRunning bool `json:"epgRunning"`
		// SkipCommercials indicates if commercials should be skipped.
		SkipCommercials bool `json:"skipCommercials"`
		// Autorec is the autorec settings for the DVR.
		Autorec DVRConfigAutorecSettings `json:"autorec"`
	}

	// DVRConfigAutorecSettings defines the autorec settings for the DVR.
	DVRConfigAutorecSettings struct {
		// MaxCount is the maximum number of autorec that can be matched.
		MaxCount int `json:"maxCount"`
		// MaxSchedules is the maximum number of autorec that can be scheduled.
		MaxSchedules int `json:"maxSchedules"`
	}

	// DVRConfigArtworkSettings defines the artwork settings for the DVR.
	DVRConfigArtworkSettings struct {
		// Fetch indicates if artwork should be fetched from installed providers.
		// `tmdb` and `tvdb` must be correctly configured in the Tvheadend configuration.
		Fetch bool `json:"fetch"`
		// AllowUnidentifiableBroadcasts indicates if artwork should be fetched for broadcasts that cannot be identified.
		AllowUnidentifiableBroadcasts bool `json:"allowUnidentifiableBroadcasts"`
		// CommandLineOptions are additional command line options to use when fetching artwork.
		CommandLineOptions string `json:"commandLineOptions"`
	}

	DVRConfigHooks struct {
		// Start is the command to run when the DVR starts recording.
		Start string `json:"start"`
		// Stop is the command to run when the DVR stops recording.
		Stop string `json:"stop"`
		// Remove is the command to run when the DVR removes a recording.
		Remove string `json:"remove"`
	}
)

type (
	DVRConfigService interface {
		GetAll(ctx context.Context) ([]DVRConfig, error)
		Delete(ctx context.Context, id string) error
	}
)

func NewDVRConfigPriority(priority int) DVRConfigPriority {
	switch priority {
	case 0:
		return DVRConfigPriorityImportant
	case 1:
		return DVRConfigPriorityHigh
	case 2:
		return DVRConfigPriorityNormal
	case 3:
		return DVRConfigPriorityLow
	case 4:
		return DVRConfigPriorityUnimportant
	case 6:
		return DVRConfigPriorityDefault
	}

	return DVRConfigPriorityUnknown
}

func NewDVRConfig(cfg tvheadend.DVRConfig) DVRConfig {
	return DVRConfig{
		ID:                      cfg.UUID,
		Enabled:                 cfg.Enabled,
		Name:                    cfg.Name,
		Original:                cfg.Name == "",
		StreamProfileID:         cfg.Profile,
		Priority:                NewDVRConfigPriority(cfg.Pri),
		DeleteAfterPlaybackTime: cfg.RemoveAfterPlayback,
		RetentionDays:           cfg.RetentionDays,
		RemovalDays:             cfg.RemovalDays,
		StartPadding:            cfg.PreExtraTime,
		EndPadding:              cfg.PostExtraTime,
		Clone:                   cfg.Clone,
		RerecordErrors:          cfg.RerecordErrors,
		TunerWarmUpTime:         cfg.WarmTime,
		Storage: DVRConfigStorageSettings{
			Path:                 cfg.Storage,
			MaintainFreeSpace:    cfg.StorageMfree,
			MaintainUsedSpace:    cfg.StorageMused,
			DirectoryPermissions: cfg.DirectoryPermissions,
			FilePermissions:      cfg.FilePermissions,
			Charset:              cfg.Charset,
			PathnameFormat:       cfg.Pathname,
			CacheScheme:          cfg.Cache,
		},
		Subdirectories: DVRConfigSubdirectorySettings{
			DaySubdir:            cfg.DayDir,
			ChannelSubdir:        cfg.ChannelDir,
			TitleSubdir:          cfg.TitleDir,
			TvMoviesSubdirFormat: cfg.FormatTvmoviesSubdir,
			TvShowsSubdirFormat:  cfg.FormatTvshowsSubdir,
		},
		File: DVRConfigFileSettings{
			IncludeChannel:            cfg.ChannelInTitle,
			IncludeDate:               cfg.DateInTitle,
			IncludeTime:               cfg.TimeInTitle,
			IncludeEpisode:            cfg.EpisodeInTitle,
			IncludeSubtitle:           cfg.SubtitleInTitle,
			OmitTitle:                 cfg.OmitTitle,
			CleanTitle:                cfg.CleanTitle,
			AllowWhitespace:           cfg.WhitespaceInTitle,
			WindowsCompatibleFilename: cfg.WindowsCompatibleFilenames,
			TagFiles:                  cfg.TagFiles,
		},
		EPG: DVRConfigEPGSettings{
			DuplicateHandling: cfg.Record,
			EpgUpdateWindow:   cfg.EpgUpdateWindow,
			EpgRunning:        cfg.EpgRunning,
			SkipCommercials:   cfg.SkipCommercials,
			Autorec: DVRConfigAutorecSettings{
				MaxCount:     cfg.AutorecMaxcount,
				MaxSchedules: cfg.AutorecMaxsched,
			},
		},
		Artwork: DVRConfigArtworkSettings{
			Fetch:                         cfg.FetchArtwork,
			AllowUnidentifiableBroadcasts: cfg.FetchArtworkKnownBroadcastsAllowUnknown,
			CommandLineOptions:            cfg.FetchArtworkOptions,
		},
		Hooks: DVRConfigHooks{
			Start:  cfg.Preproc,
			Stop:   cfg.Postproc,
			Remove: cfg.Postremove,
		},
	}
}
