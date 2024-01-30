import { EpgEvent } from '@/clients/api/api.types';
import { useFetchEvent } from '@/hooks/epg';
import { useManageRecordingByEvent } from '@/hooks/recording';
import { cleanup, render } from '@testing-library/react';
import { PropsWithChildren } from 'react';
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import { afterEach, expect, test, vi } from 'vitest';
import EventView from './EventView';
import userEvent from '@testing-library/user-event';

vi.mock('@/hooks/epg');
vi.mock('@/hooks/recording');

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

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

test('should render without related events', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: buildEvent(),
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should render with recording', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: buildEvent(RECORDING_ID),
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should render with pending button', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: buildEvent(),
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: true,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should render with related events', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: buildEvent(),
    relatedEvents,
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should render error', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: 'some error',
    event: undefined,
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should render nothing', () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: undefined,
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();
  expect(fetchMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should create recording', async () => {
  const fetchMock = vi.fn();
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: fetchMock,
    error: null,
    event: buildEvent(),
    relatedEvents: [],
  });

  const createRecordingMock = vi.fn();
  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: createRecordingMock,
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });

  await userEvent.click(document.getByText('record'));
  expect(createRecordingMock).toHaveBeenNthCalledWith(1, EVENT_ID);
});

test('should navigate to recording', async () => {
  vi.mocked(useFetchEvent).mockReturnValue({
    fetch: async () => {},
    error: null,
    event: buildEvent(RECORDING_ID),
    relatedEvents: [],
  });

  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    createRecording: async () => {},
  } as any);

  const document = render(<EventView />, { wrapper: TestRouter });
  await userEvent.click(document.getByText('modify_recording'));

  expect(document.getByText('expect_recording_here')).toBeInTheDocument();
});

const TestRouter = ({ children }: PropsWithChildren) => {
  return (
    <MemoryRouter initialEntries={[`/guide/events/${EVENT_ID}`]}>
      <Routes>
        <Route path="/guide/events/:id" element={children} />
        <Route
          path={`/recordings/${RECORDING_ID}`}
          element={<span>expect_recording_here</span>}
        />
      </Routes>
    </MemoryRouter>
  );
};
