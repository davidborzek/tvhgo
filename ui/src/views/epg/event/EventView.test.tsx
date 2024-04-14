import { afterEach, beforeEach, expect, test, vi } from 'vitest';
import { cleanup, render } from '@testing-library/react';
import { useLoaderData, useNavigate, useRevalidator } from 'react-router-dom';

import { EpgEvent } from '@/clients/api/api.types';
import { Component as EventView } from './EventView';
import { useManageRecordingByEvent } from '@/hooks/recording';
import userEvent from '@testing-library/user-event';

vi.mock('@/hooks/epg');
vi.mock('@/hooks/recording');
vi.mock('react-router-dom');

const EVENT_ID = 1;
const RECORDING_ID = 'someRecordingID';

const buildEvent = (dvrUuid?: string): EpgEvent => {
  return {
    audioDesc: true,
    channelId: 'someChannelId',
    channelName: 'Some Channel',
    channelNumber: 1,
    description: 'Some Description',
    endsAt: 0,
    hd: true,
    id: EVENT_ID,
    nextEventId: 2,
    piconId: 1,
    startsAt: 0,
    subtitle: 'Some Subtitle',
    subtitled: true,
    title: 'Some Title',
    widescreen: true,
    dvrUuid,
  };
};

const relatedEvents: EpgEvent[] = [
  {
    audioDesc: true,
    channelId: 'someChannelId',
    channelName: 'Some Channel',
    channelNumber: 1,
    description: 'Some Description 2',
    endsAt: 0,
    hd: true,
    id: 2,
    nextEventId: 3,
    piconId: 1,
    startsAt: 0,
    subtitle: 'Some Subtitle 2',
    subtitled: true,
    title: 'Some Title 2',
    widescreen: true,
  },
  {
    audioDesc: true,
    channelId: 'someChannelId2',
    channelName: 'Some Channel2',
    channelNumber: 1,
    description: 'Some Description 3',
    endsAt: 0,
    hd: true,
    id: 3,
    nextEventId: 4,
    piconId: 1,
    startsAt: 0,
    subtitle: 'Some Subtitle 3',
    subtitled: true,
    title: 'Some Title 3',
    widescreen: true,
  },
];

const navigateMock = vi.fn();
const revalidateMock = vi.fn();

beforeEach(() => {
  vi.mocked(useNavigate).mockReturnValue(navigateMock);
  vi.mocked(useRevalidator).mockReturnValue({
    revalidate: revalidateMock,
    state: 'idle',
  });
});

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

test('should render without related events', () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(), []]);

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('should render with recording', () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(RECORDING_ID), []]);

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('should render with pending button', () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(), []]);

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: true,
    createRecording: async () => {},
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('should render with related events', () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(), relatedEvents]);

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('should create recording', async () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(), []]);

  const createRecordingMock = vi.fn();
  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: createRecordingMock,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);

  await userEvent.click(document.getByText('record'));
  expect(createRecordingMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should navigate to recording', async () => {
  vi.mocked(useLoaderData).mockReturnValue([buildEvent(RECORDING_ID), []]);

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } as any);

  const document = render(<EventView />);
  await userEvent.click(document.getByText('modify_recording'));

  expect(navigateMock).toHaveBeenCalledWith(`/recordings/${RECORDING_ID}`);
});
