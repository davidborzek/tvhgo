package tvheadend

type (
	GridResponse[T any] struct {
		Entries []T   `json:"entries"`
		Total   int64 `json:"total"`
	}

	ListResponse[T any] struct {
		Entries []T `json:"entries"`
	}

	ChannelGridEntry struct {
		UUID            string        `json:"uuid"`
		Enabled         bool          `json:"enabled"`
		Autoname        bool          `json:"autoname"`
		Name            string        `json:"name"`
		Number          int           `json:"number"`
		Icon            string        `json:"icon"`
		IconPublicURL   string        `json:"icon_public_url"`
		EpgAuto         bool          `json:"epgauto"`
		EpgLimit        int           `json:"epglimit"`
		EpgGrab         []interface{} `json:"epggrab"`
		DvrPreTime      int           `json:"dvr_pre_time"`
		DvrPstTime      int           `json:"dvr_pst_time"`
		EpgRunning      int           `json:"epg_running"`
		RemoteTimeshift bool          `json:"remote_timeshift"`
		Services        []string      `json:"services"`
		Tags            []string      `json:"tags"`
		Bouquet         string        `json:"bouquet"`
	}

	ChannelGrid GridResponse[ChannelGridEntry]

	DvrGridEntry struct {
		UUID        string `json:"uuid"`
		Enabled     bool   `json:"enabled"`
		Create      int64  `json:"create"`
		Watched     int    `json:"watched"`
		Start       int64  `json:"start"`
		StartExtra  int    `json:"start_extra"`
		StartReal   int64  `json:"start_real"`
		Stop        int64  `json:"stop"`
		StopExtra   int    `json:"stop_extra"`
		StopReal    int64  `json:"stop_real"`
		Duration    int64  `json:"duration"`
		Channel     string `json:"channel"`
		ChannelIcon string `json:"channel_icon"`
		Channelname string `json:"channelname"`
		Image       string `json:"image"`
		FanartImage string `json:"fanart_image"`
		// Title title of the event for each language
		Title     map[string]string `json:"title"`
		DispTitle string            `json:"disp_title"`
		// Subtitle subtitle of the event for each language
		Subtitle     map[string]string `json:"subtitle"`
		DispSubtitle string            `json:"disp_subtitle"`
		DispSummary  string            `json:"disp_summary"`
		// Desc description of the event for each language
		Description     map[string]string `json:"description"`
		DispDescription string            `json:"disp_description"`
		DispExtratext   string            `json:"disp_extratext"`
		Pri             int               `json:"pri"`
		Retention       int               `json:"retention"`
		Removal         int64             `json:"removal"`
		Playposition    int               `json:"playposition"`
		Playcount       int               `json:"playcount"`
		ConfigName      string            `json:"config_name"`
		Creator         string            `json:"creator"`
		Filename        string            `json:"filename"`
		Errorcode       int               `json:"errorcode"`
		Errors          int               `json:"errors"`
		DataErrors      int               `json:"data_errors"`
		DvbEid          int               `json:"dvb_eid"`
		Noresched       bool              `json:"noresched"`
		Norerecord      bool              `json:"norerecord"`
		Fileremoved     int               `json:"fileremoved"`
		Autorec         string            `json:"autorec"`
		AutorecCaption  string            `json:"autorec_caption"`
		Timerec         string            `json:"timerec"`
		TimerecCaption  string            `json:"timerec_caption"`
		Parent          string            `json:"parent"`
		Child           string            `json:"child"`
		ContentType     int               `json:"content_type"`
		CopyrightYear   int               `json:"copyright_year"`
		Broadcast       int64             `json:"broadcast"`
		EpisodeDisp     string            `json:"episode_disp"`
		URL             string            `json:"url"`
		Filesize        int               `json:"filesize"`
		Status          string            `json:"status"`
		SchedStatus     string            `json:"sched_status"`
		Duplicate       int               `json:"duplicate"`
		FirstAired      int               `json:"first_aired"`
		Category        []interface{}     `json:"category"`
		Credits         interface{}       `json:"credits"`
		Keyword         []interface{}     `json:"keyword"`
		Genre           []int             `json:"genre"`
	}

	DvrGrid GridResponse[DvrGridEntry]

	DvrCreateRecordingOpts struct {
		DispTitle     string `json:"disp_title"`
		DispExtratext string `json:"disp_extratext,omitempty"`
		Channel       string `json:"channel"`
		Start         int64  `json:"start"`
		Stop          int64  `json:"stop"`
		Comment       string `json:"comment,omitempty"`
		StartExtra    int    `json:"start_extra,omitempty"`
		StopExtra     int    `json:"stop_extra,omitempty"`
		Pri           int    `json:"pri,omitempty"`
		ConfigName    string `json:"config_name,omitempty"`
	}

	DvrUpdateRecordingOpts struct {
		Enabled       bool   `json:"enabled"`
		DispTitle     string `json:"disp_title"`
		DispExtratext string `json:"disp_extratext"`
		Channel       string `json:"channel"`
		Start         int64  `json:"start"`
		Stop          int64  `json:"stop"`
		Comment       string `json:"comment"`
		EpisodeDisp   string `json:"episode_disp"`
		StartExtra    int    `json:"start_extra"`
		StopExtra     int    `json:"stop_extra"`
		Pri           int    `json:"pri"`
		ConfigName    string `json:"config_name"`
		Owner         string `json:"owner"`
		Creator       string `json:"creator"`
		Removal       int    `json:"removal"`
		Retention     int    `json:"retention"`
		UUID          string `json:"uuid"`
	}

	DvrRecordingCreated struct {
		UUID string `json:"uuid"`
	}

	DvrConfig struct {
		UUID                                    string `json:"uuid"`
		Enabled                                 bool   `json:"enabled"`
		Name                                    string `json:"name"`
		Profile                                 string `json:"profile"`
		Pri                                     int    `json:"pri"`
		RetentionDays                           int    `json:"retention-days"`
		RemovalDays                             int64  `json:"removal-days"`
		RemoveAfterPlayback                     int    `json:"remove-after-playback"`
		PreExtraTime                            int    `json:"pre-extra-time"`
		PostExtraTime                           int    `json:"post-extra-time"`
		Clone                                   bool   `json:"clone"`
		RerecordErrors                          int    `json:"rerecord-errors"`
		ComplexScheduling                       bool   `json:"complex-scheduling"`
		FetchArtwork                            bool   `json:"fetch-artwork"`
		FetchArtworkKnownBroadcastsAllowUnknown bool   `json:"fetch-artwork-known-broadcasts-allow-unknown"`
		Storage                                 string `json:"storage"`
		StorageMfree                            int    `json:"storage-mfree"`
		StorageMused                            int    `json:"storage-mused"`
		DirectoryPermissions                    string `json:"directory-permissions"`
		FilePermissions                         string `json:"file-permissions"`
		Charset                                 string `json:"charset"`
		Pathname                                string `json:"pathname"`
		Cache                                   int    `json:"cache"`
		DayDir                                  bool   `json:"day-dir"`
		ChannelDir                              bool   `json:"channel-dir"`
		TitleDir                                bool   `json:"title-dir"`
		FormatTvmoviesSubdir                    string `json:"format-tvmovies-subdir"`
		FormatTvshowsSubdir                     string `json:"format-tvshows-subdir"`
		ChannelInTitle                          bool   `json:"channel-in-title"`
		DateInTitle                             bool   `json:"date-in-title"`
		TimeInTitle                             bool   `json:"time-in-title"`
		EpisodeInTitle                          bool   `json:"episode-in-title"`
		SubtitleInTitle                         bool   `json:"subtitle-in-title"`
		OmitTitle                               bool   `json:"omit-title"`
		CleanTitle                              bool   `json:"clean-title"`
		WhitespaceInTitle                       bool   `json:"whitespace-in-title"`
		WindowsCompatibleFilenames              bool   `json:"windows-compatible-filenames"`
		TagFiles                                bool   `json:"tag-files"`
		EpgUpdateWindow                         int    `json:"epg-update-window"`
		EpgRunning                              bool   `json:"epg-running"`
		AutorecMaxcount                         int    `json:"autorec-maxcount"`
		AutorecMaxsched                         int    `json:"autorec-maxsched"`
		SkipCommercials                         bool   `json:"skip-commercials"`
		WarmTime                                int    `json:"warm-time"`
	}

	DvrConfigGrid GridResponse[DvrConfig]

	EpgEventGridEntry struct {
		EventID       int64  `json:"eventId"`
		ChannelName   string `json:"channelName"`
		ChannelUUID   string `json:"channelUuid"`
		ChannelNumber string `json:"channelNumber"`
		ChannelIcon   string `json:"channelIcon"`
		Start         int64  `json:"start"`
		Stop          int64  `json:"stop"`
		Title         string `json:"title"`
		Subtitle      string `json:"subtitle,omitempty"`
		Description   string `json:"description,omitempty"`
		Widescreen    int    `json:"widescreen,omitempty"`
		NextEventID   int    `json:"nextEventId"`
		Genre         []int  `json:"genre,omitempty"`
		Subtitled     int    `json:"subtitled,omitempty"`
		AudioDesc     int    `json:"audiodesc,omitempty"`
		HD            int    `json:"hd,omitempty"`
		DvrUUID       string `json:"dvrUuid,omitempty"`
		DvrState      string `json:"dvrState,omitempty"`
	}

	EpgEventGrid struct {
		Entries []EpgEventGridEntry `json:"entries"`
		Total   int64               `json:"totalCount"`
	}

	EpgContentType struct {
		Key int    `json:"key"`
		Val string `json:"val"`
	}

	EpgContentTypeResponse ListResponse[EpgContentType]

	InodeParams struct {
		ID          string      `json:"id"`
		Type        string      `json:"type"`
		Caption     string      `json:"caption"`
		Description string      `json:"description"`
		Value       interface{} `json:"value"`
	}

	Idnode struct {
		UUID    string        `json:"uuid"`
		ID      string        `json:"id"`
		Text    string        `json:"text"`
		Caption string        `json:"caption"`
		Class   string        `json:"class"`
		Event   string        `json:"event"`
		Params  []InodeParams `json:"params"`
	}

	IdnodeLoadResponse ListResponse[Idnode]

	FilterQuery struct {
		Field      string      `json:"field"`
		Type       string      `json:"type"`
		Value      interface{} `json:"value"`
		Comparison string      `json:"comparison,omitempty"`
		Intsplit   string      `json:"intsplit,omitempty"`
	}
)
