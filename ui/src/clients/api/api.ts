import axios from 'axios';
import {
  Channel,
  EpgChannel,
  EpgEvent,
  ErrorResponse,
  ListResponse,
  Recording,
  UpdateRecording,
  UserResponse,
} from './api.types';

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

export class ApiError extends Error {
  constructor(public readonly code: number, message: string) {
    super(message);
  }
}

const client = axios.create({
  baseURL: '/api',
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

export async function login(username: string, password: string) {
  await client.post('/login', {
    username,
    password,
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

export async function getEpgChannelEvents(
  q?: GetEpgChannelEventsQuery
): Promise<ListResponse<EpgChannel>> {
  const response = await client.get<ListResponse<EpgChannel>>(
    '/epg/channel/events',
    {
      params: q,
    }
  );
  return response.data;
}

export async function getChannel(id: string): Promise<Channel> {
  const response = await client.get<Channel>(`/channels/${id}`);
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

export async function getRecordings(
  q?: GetRecordingsQuery
): Promise<Recording[]> {
  const response = await client.get<Recording[]>(`/recordings`, {
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
