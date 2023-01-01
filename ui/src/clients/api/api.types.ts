export type ListResponse<T> = {
  entries: T[];
  total: number;
  offset: number;
};

export type UserResponse = {
  id: string;
  username: string;
  email: string;
  displayNAme: string;
  createdAt: number;
  updatedAt: number;
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
