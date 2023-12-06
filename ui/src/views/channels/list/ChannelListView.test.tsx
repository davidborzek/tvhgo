import { EpgEvent } from '@/clients/api/api.types';
import { useFetchEvents } from '@/hooks/epg';
import { cleanup, render } from '@testing-library/react';
import { PropsWithChildren } from 'react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { afterEach, expect, test, vi } from 'vitest';
import ChannelListView from './ChannelListView';
import userEvent from '@testing-library/user-event';

vi.mock('@/hooks/epg');

const buildChannel = (id: string, name: string): EpgEvent => {
  return {
    audioDesc: true,
    channelId: id,
    channelName: name,
    channelNumber: 1,
    description: `${name} - Description`,
    endsAt: 0,
    hd: true,
    id: 1,
    nextEventId: 2,
    piconId: 1,
    startsAt: 0,
    subtitle: `${name} - Subtitle`,
    subtitled: true,
    title: `${name} - Title`,
    widescreen: true,
  };
};

const buildChannels = (count: number): EpgEvent[] => {
  return [...new Array(count)].map((_, i) =>
    buildChannel(`id_${i}`, `Channel_${i}`)
  );
};

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

test('render empty state', () => {
  vi.mocked(useFetchEvents).mockReturnValue({
    error: null,
    events: [],
    total: 0,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });
  expect(document.asFragment()).toMatchSnapshot();

  expectUseFetchEvents(1, 0);
});

test('render nothing', () => {
  vi.mocked(useFetchEvents).mockReturnValue({
    error: null,
    events: null,
    total: 0,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });
  expect(document.asFragment()).toMatchSnapshot();

  expectUseFetchEvents(1, 0);
});

test('render error', () => {
  vi.mocked(useFetchEvents).mockReturnValue({
    error: 'some error',
    events: null,
    total: 0,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });
  expect(document.asFragment()).toMatchSnapshot();

  expectUseFetchEvents(1, 0);
});

test('render with one page', () => {
  const channels = buildChannels(5);
  vi.mocked(useFetchEvents).mockReturnValue({
    error: null,
    events: channels,
    total: channels.length,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();

  expectUseFetchEvents(1, 0);
});

test('render with multiple pages', () => {
  const channels = buildChannels(5);
  vi.mocked(useFetchEvents).mockReturnValue({
    error: null,
    events: channels,
    total: 100,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });

  expect(document.asFragment()).toMatchSnapshot();

  expectUseFetchEvents(1, 0);
});

test('pagination', async () => {
  const channels = buildChannels(5);
  vi.mocked(useFetchEvents).mockReturnValue({
    error: null,
    events: channels,
    total: 200,
  });

  const document = render(<ChannelListView />, { wrapper: TestRouter });
  expectUseFetchEvents(1, 0);

  await userEvent.click(document.getByTestId('next_page'));
  expectUseFetchEvents(2, 50);

  await userEvent.click(document.getByTestId('previous_page'));
  expectUseFetchEvents(3, 0);

  await userEvent.click(document.getByTestId('last_page'));
  expectUseFetchEvents(4, 200);

  await userEvent.click(document.getByTestId('first_page'));
  expectUseFetchEvents(5, 0);
});

const expectUseFetchEvents = (n: number, offset: number) =>
  expect(useFetchEvents).toHaveBeenNthCalledWith(n, {
    nowPlaying: true,
    limit: 50,
    offset,
    sort_key: 'channelNumber',
    sort_dir: 'asc',
  });

const TestRouter = ({ children }: PropsWithChildren) => {
  return (
    <MemoryRouter initialEntries={[`/channels`]}>
      <Routes>
        <Route path="/channels" element={children} />
      </Routes>
    </MemoryRouter>
  );
};
