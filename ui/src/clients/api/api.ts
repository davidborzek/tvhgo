import axios from 'axios';
import {
  Channel,
  EpgChannel,
  EpgEvent,
  ErrorResponse,
  ListResponse,
  UserResponse,
} from './api.types';

type PaginationQuery = {
  limit?: number;
  offset?: number;
};

type SortQuery = {
  sort_key?: string;
  sort_direction?: string;
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
