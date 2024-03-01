import axios from 'axios';
import {
  Channel,
  CreateTokenResponse,
  EpgChannel,
  EpgEvent,
  ErrorResponse,
  ListResponse,
  Recording,
  Session,
  Token,
  TwoFactorAuthSettings,
  TwoFactorAuthSetupResult,
  UpdateRecording,
  UpdateUser,
  UpdateUserPassword,
  UserResponse,
} from './api.types';
import qs from 'qs';

type PaginationQuery = {
  limit?: number;
  offset?: number;
};

type SortQuery = {
  sort_key?: string;
  sort_dir?: string;
};

type PaginationSortQuery = PaginationQuery & SortQuery;

export type GetEpgEventsQuery = PaginationSortQuery & {
  title?: string;
  lang?: string;
  channel?: string;
  contentType?: string;
  fullText?: boolean;
  nowPlaying?: boolean;
  durationMin?: number;
  durationMax?: number;
  startsAt?: number;
  endsAt?: number;
};

export type GetEpgChannelEventsQuery = PaginationSortQuery & {
  startsAt?: number;
  endsAt?: number;
};

export type RecordingStatus = 'upcoming' | 'finished' | 'failed' | 'removed';

export type GetRecordingsQuery = PaginationSortQuery & {
  status?: RecordingStatus;
};

export type GetChannelsQuery = PaginationSortQuery & {
  name?: string;
};

export class ApiError extends Error {
  constructor(
    public readonly code: number,
    message: string
  ) {
    super(message);
  }
}

const client = axios.create({
  baseURL: '/api',
  paramsSerializer: (params) => qs.stringify(params, { arrayFormat: 'repeat' }),
});

client.interceptors.response.use(
  (res) => res,
  (err) => {
    if (axios.isAxiosError(err) && err.response) {
      const data = err.response.data as ErrorResponse;
      return Promise.reject(new ApiError(err.response.status, data.message));
    }
    return Promise.reject(err);
  }
);

export async function login(username: string, password: string, code?: string) {
  await client.post('/login', {
    username,
    password,
    totp: code,
  });
}

export async function logout(): Promise<void> {
  await client.post('/logout');
}

export async function getUser(): Promise<UserResponse> {
  const response = await client.get<UserResponse>('/user');
  return response.data;
}

export async function getEpgEvents(
  q?: GetEpgEventsQuery
): Promise<ListResponse<EpgEvent>> {
  const response = await client.get<ListResponse<EpgEvent>>('/epg/events', {
    params: q,
  });
  return response.data;
}

export async function getRelatedEpgEvents(
  id: number,
  q?: PaginationSortQuery
): Promise<ListResponse<EpgEvent>> {
  const response = await client.get<ListResponse<EpgEvent>>(
    `/epg/events/${id}/related`,
    {
      params: q,
    }
  );
  return response.data;
}

export async function getEpgEvent(id: number): Promise<EpgEvent> {
  const response = await client.get<EpgEvent>(`/epg/events/${id}`);
  return response.data;
}

export async function getEpg(
  q?: GetEpgChannelEventsQuery
): Promise<Array<EpgChannel>> {
  const response = await client.get<Array<EpgChannel>>('/epg', {
    params: q,
  });
  return response.data;
}

export async function getChannel(id: string): Promise<Channel> {
  const response = await client.get<Channel>(`/channels/${id}`);
  return response.data;
}

export async function getChannels(
  q?: GetChannelsQuery
): Promise<Array<Channel>> {
  const response = await client.get<Array<Channel>>(`/channels`, {
    params: q,
  });
  return response.data;
}

export async function fetch<T>(path: string): Promise<T> {
  const response = await client.get<T>(path);
  return response.data;
}

export async function recordByEvent(
  eventId: number,
  configId?: string
): Promise<void> {
  await client.post('/recordings/event', {
    eventId,
    configId,
  });
}

export async function stopRecording(id: string): Promise<void> {
  await client.put(`/recordings/${id}/stop`);
}

export async function cancelRecording(id: string): Promise<void> {
  await client.put(`/recordings/${id}/cancel`);
}

export async function removeRecording(id: string): Promise<void> {
  await client.delete(`/recordings/${id}`);
}

export async function stopRecordings(ids: string[]): Promise<void> {
  await client.put(`/recordings/stop`, null, {
    params: { ids },
  });
}

export async function cancelRecordings(ids: string[]): Promise<void> {
  await client.put(`/recordings/cancel`, null, {
    params: { ids },
  });
}

export async function removeRecordings(ids: string[]): Promise<void> {
  await client.delete(`/recordings`, {
    params: { ids },
  });
}

export async function getRecordings(
  q?: GetRecordingsQuery
): Promise<ListResponse<Recording>> {
  const response = await client.get<ListResponse<Recording>>(`/recordings`, {
    params: q,
  });
  return response.data;
}

export async function getRecording(id: string): Promise<Recording> {
  const response = await client.get<Recording>(`/recordings/${id}`);
  return response.data;
}

export async function updateRecording(
  id: string,
  opts: UpdateRecording
): Promise<void> {
  await client.patch(`/recordings/${id}`, opts);
}

export async function updateUser(opts: UpdateUser): Promise<UserResponse> {
  const response = await client.patch<UserResponse>(`/user`, opts);
  return response.data;
}

export async function updateUserPassword(
  opts: UpdateUserPassword
): Promise<void> {
  await client.patch(`/user/password`, opts);
}

export function getRecordingUrl(id: string): string {
  return `/api/recordings/${id}/stream`;
}

export async function getSessions(): Promise<Array<Session>> {
  const response = await client.get<Array<Session>>(`/sessions`);
  return response.data;
}

export async function deleteSession(id: number): Promise<void> {
  return await client.delete(`/sessions/${id}`);
}

export async function getTwoFactorAuthSettings(): Promise<TwoFactorAuthSettings> {
  const response = await client.get<TwoFactorAuthSettings>(`/two-factor-auth`);
  return response.data;
}

export async function deactivateTwoFactorAuth(code: string): Promise<void> {
  await client.put(`/two-factor-auth/deactivate`, {
    code,
  });
}

export async function setupTwoFactorAuth(
  password: string
): Promise<TwoFactorAuthSetupResult> {
  const response = await client.put<TwoFactorAuthSetupResult>(
    `/two-factor-auth/setup`,
    {
      password,
    }
  );
  return response.data;
}

export async function activateTwoFactorAuth(
  password: string,
  code: string
): Promise<void> {
  await client.put(`/two-factor-auth/activate`, {
    password,
    code,
  });
}

export async function getTokens(): Promise<Array<Token>> {
  const response = await client.get<Array<Token>>(`/tokens`);
  return response.data;
}

export async function deleteToken(id: number): Promise<void> {
  await client.delete(`/tokens/${id}`);
}

export async function createToken(name: string): Promise<CreateTokenResponse> {
  const response = await client.post<CreateTokenResponse>(`/tokens`, { name });
  return response.data;
}
