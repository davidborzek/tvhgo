package core

import (
	"context"
	"errors"

	"github.com/davidborzek/tvhgo/tvheadend"
)

var (
	ErrDVRConfigNotFound = errors.New("dvr config not found")
)

const (
	TvheadendRetentionForever = 2147483647
	TvheadendRetentionOther   = 2147483646
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

type DVRConfigRetentionType string

const (
	DVRConfigRetentionTypeForever         DVRConfigRetentionType = "forever"
	DVRConfigRetentionTypeDays            DVRConfigRetentionType = "days"
	DVRConfigRetentionTypeMaintainedSpace DVRConfigRetentionType = "maintained_space"
	DVRConfigRetentionTypeOnFileRemoval   DVRConfigRetentionType = "on_file_removal"
)

type DVRConfigRetentionTarget string

const (
	DVRConfigRetentionTargetInfo DVRConfigRetentionTarget = "info"
	DVRConfigRetentionTargetFile DVRConfigRetentionTarget = "file"
)

type DVRConfigCacheScheme string

const (
	DVRConfigCacheSchemeUnknown          DVRConfigCacheScheme = "unknown"
	DVRConfigCacheSchemeSystem           DVRConfigCacheScheme = "system"
	DVRConfigCacheSchemeDoNotKeep        DVRConfigCacheScheme = "do_not_keep"
	DVRConfigCacheSchemeSync             DVRConfigCacheScheme = "sync"
	DVRConfigCacheSchemeSyncAndDoNotKeep DVRConfigCacheScheme = "sync_and_do_not_keep"
)

type DVRConfigDuplicateHandling string

const (
	DVRConfigDuplicateHandlingRecordAll                       DVRConfigDuplicateHandling = "record_all"
	DVRConfigDuplicateHandlingRecordAllEpgUnique              DVRConfigDuplicateHandling = "all_epg_unique"
	DVRConfigDuplicateHandlingRecordAllDifferentEpisode       DVRConfigDuplicateHandling = "all_different_episode"
	DVRConfigDuplicateHandlingRecordAllDifferentSubtitle      DVRConfigDuplicateHandling = "all_different_subtitle"
	DVRConfigDuplicateHandlingRecordAllDifferentDescription   DVRConfigDuplicateHandling = "all_different_description"
	DVRConfigDuplicateHandlingRecordAllOncePerMonth           DVRConfigDuplicateHandling = "all_once_per_month"
	DVRConfigDuplicateHandlingRecordAllOncePerWeek            DVRConfigDuplicateHandling = "all_once_per_week"
	DVRConfigDuplicateHandlingRecordAllOncePerDay             DVRConfigDuplicateHandling = "all_once_per_day"
	DVRConfigDuplicateHandlingRecordLocalDifferentEpisode     DVRConfigDuplicateHandling = "local_different_episode"
	DVRConfigDuplicateHandlingRecordLocalDifferentTitle       DVRConfigDuplicateHandling = "local_different_title"
	DVRConfigDuplicateHandlingRecordLocalDifferentSubtitle    DVRConfigDuplicateHandling = "local_different_subtitle"
	DVRConfigDuplicateHandlingRecordLocalDifferentDescription DVRConfigDuplicateHandling = "local_different_description"
	DVRConfigDuplicateHandlingRecordLocalOncePerMonth         DVRConfigDuplicateHandling = "local_once_per_month"
	DVRConfigDuplicateHandlingRecordLocalOncePerWeek          DVRConfigDuplicateHandling = "local_once_per_week"
	DVRConfigDuplicateHandlingRecordLocalOncePerDays          DVRConfigDuplicateHandling = "local_once_per_day"
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
		// RecordingInfoRetention is the retention policy for the recording info.
		RecordingInfoRetention DVRConfigRetentionPolicy `json:"recordingInfoRetention"`
		// RecordingFileRetention is the retention policy for the recording file.
		RecordingFileRetention DVRConfigRetentionPolicy `json:"recordingFileRetention"`
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
		CacheScheme DVRConfigCacheScheme `json:"cacheScheme"`
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
		IncludeChannel bool `json:"includeChannel"`
		// IncludeDate indicates if the date should be included in the title.
		// This applies fot the filename and the title in tag stored in the file.
		IncludeDate bool `json:"includeDate"`
		// IncludeTime indicates if the time should be included in the title.
		// This applies fot the filename and the title in tag stored in the file.
		IncludeTime bool `json:"includeTime"`
		// IncludeEpisode indicates if the episode should be included in the title when available.
		IncludeEpisode bool `json:"includeEpisode"`
		// IncludeSubtitle indicates if the subtitle should be included in the title when available.
		IncludeSubtitle bool `json:"includeSubtitle"`
		// OmitTitle indicates if the title should be omitted.
		OmitTitle bool `json:"omitTitle"`
		// CleanTitle indicates if the title should be cleaned.
		CleanTitle bool `json:"cleanTitle"`
		// AllowWhitespace indicates if whitespace should be included in the title.
		AllowWhitespace bool `json:"allowWhitespace"`
		// WindowsCompatibleFilename indicates if the filename should be Windows compatible.
		WindowsCompatibleFilename bool `json:"windowsCompatibleFilename"`
		// TagFiles indicates if the files should be tagged with metadata.
		TagFiles bool `json:"tagFiles"`
	}

	// DVRConfigEPGSettings defines the EPG settings for the DVR.
	DVRConfigEPGSettings struct {
		// DuplicateHandling defines the duplicate recording handling.
		DuplicateHandling DVRConfigDuplicateHandling `json:"duplicateHandling"`
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

	DVRConfigRetentionPolicy struct {
		// Type is the type of retention policy.
		Type DVRConfigRetentionType `json:"type"`
		// Days is the amount of days to keep the recording info.
		Days int64 `json:"days"`
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

func NewDVRConfigRetentionPolicy(retention int64, t DVRConfigRetentionTarget) DVRConfigRetentionPolicy {
	policyType := DVRConfigRetentionTypeDays
	days := int64(0)

	switch retention {
	case TvheadendRetentionForever:
		policyType = DVRConfigRetentionTypeForever
	case TvheadendRetentionOther:
		if t == DVRConfigRetentionTargetFile {
			policyType = DVRConfigRetentionTypeMaintainedSpace
		} else {
			policyType = DVRConfigRetentionTypeOnFileRemoval
		}
	default:
		days = retention
	}

	return DVRConfigRetentionPolicy{
		Type: policyType,
		Days: days,
	}
}

func NewDVRConfigCacheScheme(cacheScheme int) DVRConfigCacheScheme {
	switch cacheScheme {
	case 0:
		return DVRConfigCacheSchemeUnknown
	case 1:
		return DVRConfigCacheSchemeSystem
	case 2:
		return DVRConfigCacheSchemeDoNotKeep
	case 3:
		return DVRConfigCacheSchemeSync
	case 4:
		return DVRConfigCacheSchemeSyncAndDoNotKeep
	}

	return DVRConfigCacheSchemeUnknown
}

func NewDVRConfigDuplicateHandling(duplicateHandling int) DVRConfigDuplicateHandling {
	switch duplicateHandling {
	case 0:
		return DVRConfigDuplicateHandlingRecordAll
	case 1:
		return DVRConfigDuplicateHandlingRecordAllDifferentEpisode
	case 2:
		return DVRConfigDuplicateHandlingRecordAllDifferentSubtitle
	case 3:
		return DVRConfigDuplicateHandlingRecordAllDifferentDescription
	case 4:
		return DVRConfigDuplicateHandlingRecordAllOncePerWeek
	case 5:
		return DVRConfigDuplicateHandlingRecordAllOncePerDay
	case 6:
		return DVRConfigDuplicateHandlingRecordLocalDifferentEpisode
	case 7:
		return DVRConfigDuplicateHandlingRecordLocalDifferentTitle
	case 8:
		return DVRConfigDuplicateHandlingRecordLocalDifferentSubtitle
	case 9:
		return DVRConfigDuplicateHandlingRecordLocalDifferentDescription
	case 10:
		return DVRConfigDuplicateHandlingRecordLocalOncePerWeek
	case 11:
		return DVRConfigDuplicateHandlingRecordLocalOncePerDays
	case 12:
		return DVRConfigDuplicateHandlingRecordAllOncePerMonth
	case 13:
		return DVRConfigDuplicateHandlingRecordLocalOncePerMonth
	case 14:
		return DVRConfigDuplicateHandlingRecordAllEpgUnique
	}

	return DVRConfigDuplicateHandlingRecordAll
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
		RecordingInfoRetention:  NewDVRConfigRetentionPolicy(cfg.RetentionDays, DVRConfigRetentionTargetInfo),
		RecordingFileRetention:  NewDVRConfigRetentionPolicy(cfg.RemovalDays, DVRConfigRetentionTargetFile),
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
			CacheScheme:          NewDVRConfigCacheScheme(cfg.Cache),
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
			DuplicateHandling: NewDVRConfigDuplicateHandling(cfg.Record),
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
