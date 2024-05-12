export type ListResponse<T> = {
  entries: T[];
  total: number;
  offset: number;
};

export type UserResponse = {
  id: number;
  username: string;
  email: string;
  displayName: string;
  createdAt: number;
  updatedAt: number;
  twoFactor: boolean;
  isAdmin: boolean;
};

export type AuthInfo = {
  userId: number;
  sessionId: number | null;
  forwardAuth: boolean;
};

export type ErrorResponse = {
  message: string;
};

export type EpgEvent = {
  id: number;
  audioDesc: boolean;
  channelId: string;
  channelName: string;
  channelNumber: number;
  piconId: number;
  description: string;
  endsAt: number;
  hd: boolean;
  nextEventId: number;
  startsAt: number;
  subtitle: string;
  subtitled: boolean;
  title: string;
  widescreen: boolean;
  dvrUuid?: string;
  dvrState?: string;
};

export type EpgChannel = {
  channelId: string;
  channelName: string;
  channelNumber: number;
  piconId: number;
  events: EpgEvent[];
};

export type Channel = {
  id: string;
  enabled: boolean;
  name: string;
  number: number;
  piconId: number;
};

export type Recording = {
  channelId: string;
  eventId: number;
  piconId: number;
  channelName: string;
  createdAt: number;
  duration: number;
  enabled: boolean;
  filename: string;
  id: string;
  langTitle: {
    ger: string;
  };
  title: string;
  subtitle: string;
  description: string;
  extraText: string;
  originalStartsAt: number;
  originalEndsAt: number;
  startsAt: number;
  endsAt: number;
  startPadding: number;
  endPadding: number;
  status: string;
};

export type UpdateRecording = {
  title?: string;
  extraText?: string;
  startsAt?: number;
  endsAt?: number;
  comment?: string;
  startPadding?: number;
  endPadding?: number;
  priority?: number;
  enabled?: boolean;
  episode?: string;
};

export type UpdateUser = {
  displayName?: string;
  email?: string;
  username?: string;
};

export type CreateUser = {
  displayName: string;
  email: string;
  username: string;
  password: string;
  isAdmin: boolean;
};

export type UpdateUserPassword = {
  currentPassword?: string;
  password?: string;
};

export type Session = {
  id: number;
  userId: number;
  clientIp: string;
  userAgent: string;
  createdAt: number;
  lastUsedAt: number;
};

export type TwoFactorAuthSettings = {
  enabled: boolean;
};

export type TwoFactorAuthSetupResult = {
  url: string;
};

export type Token = {
  id: number;
  name: string;
  createdAt: number;
  updatedAt: number;
};

export type CreateTokenResponse = {
  token: string;
};

export interface DVRConfig {
  id: string;
  enabled: boolean;
  name: string;
  original: boolean;
  streamProfileId: string;
  priority: string;
  deleteAfterPlayback: number;
  recordingInfoRetention: DVRRetentionPolicy;
  recordingFileRetention: DVRRetentionPolicy;
  startPadding: number;
  endPadding: number;
  clone: boolean;
  rerecordErrors: number;
  tunerWarmUpTime: number;
  storage: DVRConfigStorage;
  subdirectories: DVRConfigSubdirectories;
  file: DVRConfigFile;
  epg: DvrConfigEPG;
  artwork: DVRConfigArtwork;
  hooks: DVRConfigHooks;
}

export type DVRConfigStorage = {
  path: string;
  maintainFreeSpace: number;
  maintainUsedSpace: number;
  directoryPermissions: string;
  filePermissions: string;
  charset: string;
  pathnameFormat: string;
  cacheScheme: string;
};

export type DVRConfigSubdirectories = {
  daySubdir: boolean;
  channelSubdir: boolean;
  titleSubdir: boolean;
  tvMoviesSubdirFormat: string;
  tvShowsSubdirFormat: string;
};

export type DVRConfigFile = {
  includeChannel: boolean;
  includeDate: boolean;
  includeTime: boolean;
  includeEpisode: boolean;
  includeSubtitle: boolean;
  omitTitle: boolean;
  cleanTitle: boolean;
  allowWhitespace: boolean;
  windowsCompatibleFilename: boolean;
  tagFiles: boolean;
};

export type DvrConfigEPG = {
  duplicateHandling: string;
  epgUpdateWindow: number;
  epgRunning: boolean;
  skipCommercials: boolean;
  autorec: DVRConfigAutorec;
};

export type DVRConfigAutorec = {
  maxCount: number;
  maxSchedules: number;
};

export type DVRConfigArtwork = {
  fetch: boolean;
  allowUnidentifiableBroadcasts: boolean;
  commandLineOptions: string;
};

export type DVRConfigHooks = {
  start: string;
  stop: string;
  remove: string;
};

export type DVRRetentionType =
  | 'days'
  | 'forever'
  | 'on_file_removal'
  | 'maintained_space';

export type DVRRetentionPolicy = {
  days: number;
  type: DVRRetentionType;
};

export type CreateRecordingOpts = {
  title: string;
  extraText?: string;
  startsAt: number;
  endsAt: number;
  channelId: string;
  startPadding?: number;
  endPadding?: number;
  comment?: string;
  configId?: string;
};
